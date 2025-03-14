package api

import "opentalaria/protocol"

type HeartbeatAPI struct {
	Request Request
}

func (h HeartbeatAPI) Name() string {
	return "Heartbeat"
}

func (h HeartbeatAPI) GetRequest() Request {
	return h.Request
}

func (h HeartbeatAPI) GetHeaderVersion(requestVersion int16) int16 {
	return (&protocol.MetadataResponse{Version: requestVersion}).GetHeaderVersion()
}
