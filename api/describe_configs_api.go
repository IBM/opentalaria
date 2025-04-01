package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
)

func HandleDescribeConfigsRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	response := protocol.DescribeConfigsResponse{}

	response.Version = req.Header.RequestApiVersion

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	// Add the topics to the config response
	topics, err := req.Config.Plugin.ListTopics()
	if err != nil {
		return nil, 0, err
	}

	results := make([]protocol.DescribeConfigsResult, len(topics))

	for i, topic := range topics {
		results[i] = protocol.DescribeConfigsResult{
			ErrorCode:    int16(utils.ErrNoError),
			ResourceType: int8(utils.TOPIC_CONFIG_TYPE),
			ResourceName: *topic.Name,
		}
	}

	response.Results = results
	resp, err := protocol.Encode(&response)

	return resp, response.GetHeaderVersion(), err
}
