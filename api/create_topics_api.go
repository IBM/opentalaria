package api

import (
	"opentalaria/protocol"
	"opentalaria/utils"

	"github.com/google/uuid"
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
	if err != nil {
		return nil, err
	}

	response := GenerateCreateTopicsResponse(m.GetRequest().Header.RequestApiVersion, req)
	return protocol.Encode(response)
}

func GenerateCreateTopicsResponse(version int16, req protocol.CreateTopicsRequest) *protocol.CreateTopicsResponse {
	response := protocol.CreateTopicsResponse{}

	response.Version = version
	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	for _, topic := range req.Topics {
		response.Topics = append(response.Topics, protocol.CreatableTopicResult{
			Version:           req.Version,
			Name:              topic.Name,
			TopicID:           uuid.New(),
			ErrorCode:         int16(utils.ErrNoError),
			NumPartitions:     topic.NumPartitions,
			ReplicationFactor: topic.ReplicationFactor,
			Configs:           convertConfigs(topic.Configs),
		})
	}

	return &response
}

func convertConfigs(reqConfig []protocol.CreateableTopicConfig) []protocol.CreatableTopicConfigs {
	result := make([]protocol.CreatableTopicConfigs, len(reqConfig))

	for i, conf := range reqConfig {
		result[i] = protocol.CreatableTopicConfigs{
			Version: conf.Version,
			Name:    conf.Name,
			Value:   conf.Value,
		}
	}

	return result
}
