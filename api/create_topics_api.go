package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleCreateTopics(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	createTopicsRequest := protocol.CreateTopicsRequest{}
	_, err := protocol.VersionedDecode(req.Message, &createTopicsRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.CreateTopicsResponse{
		Version: req.Header.RequestApiVersion,
	}

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	for _, topic := range createTopicsRequest.Topics {
		err := req.Config.Plugin.AddTopic(topic)

		response.Topics = append(response.Topics, protocol.CreatableTopicResult{
			Version:   createTopicsRequest.Version,
			Name:      topic.Name,
			ErrorCode: int16(err),
		})
	}

	resp, err := protocol.Encode(&response)

	return resp, response.GetHeaderVersion(), err
}
