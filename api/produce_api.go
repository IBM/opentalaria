package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleProduceRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	produceRequest := protocol.ProduceRequest{}
	_, err := protocol.VersionedDecode(req.Message, &produceRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response, err := req.Config.Plugin.Produce(produceRequest)
	if err != nil {
		return nil, 0, err
	}

	response.Version = req.Header.RequestApiVersion

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	resp, err := protocol.Encode(&response)
	return resp, response.GetHeaderVersion(), err
}
