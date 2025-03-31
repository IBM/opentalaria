package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleAPIVersionsRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	// handle response
	apiVersionRequest := protocol.ApiVersionsRequest{}
	_, err := protocol.VersionedDecode(req.Message, &apiVersionRequest, apiVersion)
	if err != nil {
		return nil, 0, err
	}

	registeredAPIs := make([]protocol.ApiVersion, 0)

	if len(opts) > 0 {
		apis := opts[0].([]config.RegisteredAPI)
		for _, api := range apis {
			registeredAPIs = append(registeredAPIs, protocol.ApiVersion{
				ApiKey:     api.ApiKey,
				MinVersion: api.MinVersion,
				MaxVersion: api.MaxVersion,
			})
		}
	}

	response := protocol.ApiVersionsResponse{
		Version:        req.Header.RequestApiVersion,
		ErrorCode:      0,
		ApiKeys:        registeredAPIs,
		ThrottleTimeMs: 0,
	}

	message, err := protocol.Encode(&response)
	responseHeaderVersion := (&protocol.ApiVersionsResponse{Version: req.Header.RequestApiVersion}).GetHeaderVersion()

	return message, responseHeaderVersion, err
}
