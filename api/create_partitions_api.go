package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
)

func HandleCreatePartitionsRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	createPartitionsRequest := protocol.CreatePartitionsRequest{}
	_, err := protocol.VersionedDecode(req.Message, &createPartitionsRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.CreatePartitionsResponse{}

	response.Version = req.Header.RequestApiVersion
	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	for _, t := range createPartitionsRequest.Topics {
		errResponse := utils.ErrNoError

		topics, err := req.Config.Plugin.ListTopics([]string{t.Name})
		if err != nil || len(topics) == 0 {
			errResponse = utils.ErrInvalidTopic
		}

		topic := topics[0]

		// New topic partitions cannot be less than the already defined topic partitions.
		// This is a limitation of the vanilla Kafka broker, but we are implementing it to maintain compatibility.
		if len(topic.Partitions) > int(t.Count) {
			errResponse = utils.ErrInvalidPartitions
		} else {
			// validation passed, try to update the partitions
			if !createPartitionsRequest.ValidateOnly {
				err := req.Config.Plugin.CreatePartitions(t.Name, int(t.Count))
				if err != nil {
					errResponse = utils.ErrInvalidTopic
				}
			}
		}

		response.Results = append(response.Results, protocol.CreatePartitionsTopicResult{
			Name:         t.Name,
			ErrorCode:    int16(errResponse),
			ErrorMessage: utils.PtrTo(errResponse.Error()),
		})
	}

	resp, err := protocol.Encode(&response)

	return resp, response.GetHeaderVersion(), err
}
