package api

import (
	"github.com/ibm/opentalaria/config"

	"github.com/ibm/opentalaria/protocol"
)

type DeleteTopicsAPI struct {
	Request Request
	Config  *config.Config
}

func (m DeleteTopicsAPI) Name() string {
	return "DeleteTopics"
}

func (m DeleteTopicsAPI) GetRequest() Request {
	return m.Request
}

func (m DeleteTopicsAPI) GetHeaderVersion(requestVersion int16) int16 {
	return (&protocol.DeleteTopicsResponse{Version: requestVersion}).GetHeaderVersion()
}

func (m DeleteTopicsAPI) GeneratePayload() ([]byte, error) {
	req := protocol.DeleteTopicsRequest{}
	_, err := protocol.VersionedDecode(m.GetRequest().Message, &req, m.GetRequest().Header.RequestApiVersion)

	resp := m.GenerateDeleteTopicsResponse(m.GetRequest().Header.RequestApiVersion, req, err)

	return protocol.Encode(resp)
}

func (m DeleteTopicsAPI) GenerateDeleteTopicsResponse(version int16, req protocol.DeleteTopicsRequest, err error) *protocol.DeleteTopicsResponse {
	response := protocol.DeleteTopicsResponse{}

	response.Version = version
	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	// v5< specific code. In v6+ we have to iterate over req.Topics
	for _, topic := range req.TopicNames {
		err := m.Config.Plugin.DeleteTopic(topic)

		response.Responses = append(response.Responses, protocol.DeletableTopicResult{
			Version:   req.Version,
			Name:      &topic,
			ErrorCode: int16(err),
		})
	}

	return &response
}
