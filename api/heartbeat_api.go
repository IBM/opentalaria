package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleHeartbeatRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	reqApiVer := req.Header.RequestApiKey
	heartbeatRequest := protocol.MetadataRequest{}
	_, err := protocol.VersionedDecode(req.Message, &heartbeatRequest, reqApiVer)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.MetadataResponse{}
	response.Version = reqApiVer

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	resp, err := protocol.Encode(&response)
	return resp, reqApiVer, err
}
