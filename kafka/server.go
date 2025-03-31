package kafka

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net"
	"os"
	"runtime"
	"strconv"

	"github.com/ibm/opentalaria/config"

	"github.com/ibm/opentalaria/protocol"

	"golang.org/x/sync/semaphore"
)

type Server struct {
	host   string
	port   string
	config *config.Config

	// holds registered apis with the socket server where the key is the API Key as defined by the kafka protocol.
	apis           map[int16]config.ApiDefinition
	registeredAPIs []config.RegisteredAPI
}

type Client struct {
	conn   net.Conn
	config *config.Config

	// copied field from Server struct
	apis           map[int16]config.ApiDefinition
	registeredAPIs []config.RegisteredAPI
}

func NewServer(config *config.Config) *Server {
	var host, port string
	if len(config.Broker.Listeners) > 0 {
		listener := config.Broker.Listeners[0]
		host = listener.Host
		port = strconv.Itoa(int(listener.Port))
	}

	return &Server{
		host:   host,
		port:   port,
		config: config,
	}
}

func (server *Server) Run() {
	ctx := context.TODO()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		slog.Error("error creating tcp listener", "err", err)
		return
	}
	defer listener.Close()

	slog.Info(fmt.Sprintf("tcp server listening on %s:%s", server.host, server.port))

	cpu := os.Getenv("GOMAXPROCS")
	if cpu == "" {
		cpu = "0"
	}
	numberOfCpu, err := strconv.Atoi(cpu)
	if err != nil {
		slog.Error("error creating connection", "error", err)
		return
	}
	// Adding more CPU's only helps up to number of available Go routines
	// For example GOMAXPROCS(8) and semaphore.NewWeighted(8) means each Go routine will be executed on different CPU
	// However if we set GOMAXPROCS(4) and semaphore.NewWeighted(8) we will have only 4 CPU's to handle 8 Go routines
	runtime.GOMAXPROCS(numberOfCpu)
	slog.Debug("number of available CPU's ", "GOMAXPROCS", numberOfCpu)

	var conCapacity int64
	conPoolStr := server.config.Env.GetString("max.connections")
	if conPoolStr == "" {
		//If env variable max.connections was not set we use default val of MaxInt64
		conCapacity = math.MaxInt64
	} else {
		//If env variable is set, we need to convert it to int64
		c, err := strconv.ParseInt(conPoolStr, 10, 64)
		if err != nil {
			slog.Error("error setting max.connections", "error", err)
			return
		}
		conCapacity = c
	}

	slog.Debug("max.connections set to ", "max.connections", conCapacity)

	//semaphore package mimics a typical “worker pool” pattern,
	//but without the need to explicitly shut down idle workers when the work is done
	sem := semaphore.NewWeighted(int64(conCapacity))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("error accepting tcp connections", "err", err)
		}

		client := &Client{
			conn:           conn,
			config:         server.config,
			apis:           server.apis,
			registeredAPIs: server.registeredAPIs,
		}

		if err := sem.Acquire(ctx, 1); err != nil {
			slog.Error("Failed to acquire semaphore: %v", "err", err)
			break
		}
		go func() {
			defer sem.Release(1)
			client.handleRequest()
		}()
	}
	// Acquire all of the tokens to wait for any remaining workers to finish
	if err := sem.Acquire(ctx, int64(conCapacity)); err != nil {
		slog.Error("Failed to acquire semaphore: %v", "err", err)
	}
}

func (client *Client) handleRequest() {
	defer client.conn.Close()

Exit:
	// read from socket until there are no more bytes left.
	for {
		// first 4 bytes contain the message size
		sizeBytes := make([]byte, 4)
		_, err := io.ReadFull(client.conn, sizeBytes[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("tcp read error", "err", err)
			break
		}
		size := binary.BigEndian.Uint32(sizeBytes)

		// read the rest of the message into the buffer.
		messageBytes := make([]byte, size)

		if _, err := io.ReadFull(client.conn, messageBytes[:]); err != nil {
			slog.Error("error decoding message", "err", err)
			break
		}

		// save the message to a file to use for testing later.
		// encoded := hex.EncodeToString(messageBytes)
		// fmt.Println(encoded)

		// We parse the header twice, first time parse only API key and API version, from which we can
		// infer the correct header version and then parse that again in the API code to get the full header.
		header := &protocol.RequestHeader{}
		protocol.VersionedDecode(messageBytes, header, 1)

		slog.Debug(header.String())

		if definition, ok := client.apis[header.RequestApiKey]; ok {
			definition.ApiRequest.SetVersion(header.RequestApiVersion)
			req, err := makeRequest(messageBytes, client.config, client.conn, definition.ApiRequest.GetHeaderVersion())
			if err != nil {
				slog.Error("error creating request", "err", err)
				// This break exits the outer for loop and closes the socket connection.
				// If there is an error in the metadata exchange for example, we don't want to continue consuming the rest of the APIs.
				break Exit
			}

			var opts any

			// API 18 is a special case where we want to return all supported apis.
			// We pass all registered APIs and min/max version as optional parameters.
			if header.RequestApiKey == 18 {
				opts = client.registeredAPIs
			}

			resp, respHeaderVersion, err := definition.HandlerFunc(req, header.RequestApiVersion, opts)
			if err != nil {
				slog.Error("error executing handler func", "err", err)
				break Exit
			}

			// write response
			payload := make([]byte, 0)

			resHeader := protocol.ResponseHeader{
				Version:       respHeaderVersion,
				CorrelationID: req.Header.CorrelationID,
			}

			resHeaderBytes, err := protocol.Encode(&resHeader)
			if err != nil {
				slog.Error("error executing handler func", "err", err)
				break Exit
			}
			// TODO: calculate the payload size before merging the header with the message payload, to avoid the append operation
			payload = append(payload, resHeaderBytes...)

			payload = append(payload, resp...)

			// prepend payload size to the final byte array that will be sent back via the wire
			result := make([]byte, 0)
			result = binary.BigEndian.AppendUint32(result, uint32(len(payload)))
			result = append(result, payload...)

			slog.Debug(fmt.Sprintf("writing %d bytes", len(result)), "api", definition.ApiRequest.GetKey())

			_, err = client.conn.Write(result)
			if err != nil {
				slog.Error("error writing response", "err", err)
				break Exit
			}
		}
	}
}

func (s *Server) RegisterAPI(api protocol.API, minVersion, maxVersion int16, handlerFunc func(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error)) {
	if s.apis == nil {
		s.apis = make(map[int16]config.ApiDefinition)
	}

	apiKey := api.GetKey()

	if _, ok := s.apis[apiKey]; !ok {
		s.apis[apiKey] = config.ApiDefinition{
			ApiRequest:  api,
			HandlerFunc: handlerFunc,
		}

		s.registeredAPIs = append(s.registeredAPIs, config.RegisteredAPI{
			ApiKey:     apiKey,
			MinVersion: minVersion,
			MaxVersion: maxVersion,
		})
	}
}

func makeRequest(msg []byte, conf *config.Config, conn net.Conn, headerVersion int16) (config.Request, error) {
	// parse the full header, based on API key and version
	header := &protocol.RequestHeader{}
	headerSize, err := protocol.VersionedDecode(msg, header, headerVersion)
	if err != nil {
		return config.Request{}, err
	}

	return config.Request{
		Header:  *header,
		Config:  conf,
		Message: msg[headerSize:],
		Conn:    conn,
	}, nil
}
