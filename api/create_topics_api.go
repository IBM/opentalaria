package api

import (
	"opentalaria/protocol"
	"opentalaria/utils"
)

type CreateTopicsAPI struct {
	Request Request
}

func (m CreateTopicsAPI) Name() string {
	return "CreateTopics"
}

func (m CreateTopicsAPI) GetRequest() Request {
	return m.Request
}

func (m CreateTopicsAPI) GetHeaderVersion(requestVersion int16) int16 {
	return (&protocol.CreateTopicsResponse{Version: requestVersion}).GetHeaderVersion()
}

func (m CreateTopicsAPI) GeneratePayload() ([]byte, error) {
	req := protocol.CreateTopicsRequest{}
	_, err := protocol.VersionedDecode(m.GetRequest().Message, &req, m.GetRequest().Header.RequestApiVersion)

	resp := GenerateCreateTopicsResponse(m.GetRequest().Header.RequestApiVersion, req, err)

	return protocol.Encode(resp)
}

func GenerateCreateTopicsResponse(version int16, req protocol.CreateTopicsRequest, err error) *protocol.CreateTopicsResponse {
	response := protocol.CreateTopicsResponse{}

	response.Version = version
	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	errorCode := int16(utils.ErrNoError)
	if err != nil {
		errorCode = int16(utils.ErrInvalidRequest)
	}

	for _, topic := range req.Topics {
		response.Topics = append(response.Topics, protocol.CreatableTopicResult{
			Version:   req.Version,
			Name:      topic.Name,
			ErrorCode: errorCode,
		})
	}

	return &response
}
