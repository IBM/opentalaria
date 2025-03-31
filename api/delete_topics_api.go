package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleDeleteTopics(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	deleteTopicsRequest := protocol.DeleteTopicsRequest{}

	_, err := protocol.VersionedDecode(req.Message, &deleteTopicsRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.DeleteTopicsResponse{
		Version: req.Header.RequestApiVersion,
	}

	// v5< specific code. In v6+ we have to iterate over req.Topics
	for _, topic := range deleteTopicsRequest.TopicNames {
		err := req.Config.Plugin.DeleteTopic(topic)

		response.Responses = append(response.Responses, protocol.DeletableTopicResult{
			Version:   req.Header.RequestApiVersion,
			Name:      &topic,
			ErrorCode: int16(err),
		})
	}

	resp, err := protocol.Encode(&response)

	return resp, response.GetHeaderVersion(), err
}
