package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleHeartbeatRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	apiVer := req.Header.RequestApiKey
	heartbeatRequest := protocol.MetadataRequest{}
	_, err := protocol.VersionedDecode(req.Message, &heartbeatRequest, apiVer)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.MetadataResponse{}
	response.Version = apiVer

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	resp, err := protocol.Encode(&response)
	return resp, response.GetHeaderVersion(), err
}
