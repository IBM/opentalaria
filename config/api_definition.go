package config

import (
	"net"

	"github.com/ibm/opentalaria/protocol"
)

type Request struct {
	Header  protocol.RequestHeader
	Message []byte
	Conn    net.Conn
	Config  *Config
}

type RegisteredAPI struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

type ApiDefinition struct {
	ApiRequest protocol.API
	// HandlerFunc returns the API Response payload as a byte array, the API Response header, which is calculated by the request header and an optional error
	HandlerFunc func(req Request, apiVersion int16, opts ...any) ([]byte, int16, error)
}
