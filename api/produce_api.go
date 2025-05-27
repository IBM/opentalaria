package api

import (
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleProduceRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	produceRequest := protocol.ProduceRequest{}
	_, err := protocol.VersionedDecode(req.Message, &produceRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	produceResponse, err := req.Config.Plugin.Produce(produceRequest)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.ProduceResponse{}

	response.Version = req.Header.RequestApiVersion

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	topicResponses := make([]protocol.TopicProduceResponse, 0)

	for topic, resp := range produceResponse {
		partResponses := make([]protocol.PartitionProduceResponse, len(resp))
		for i, r := range resp {
			partResponses[i] = protocol.PartitionProduceResponse{
				Version:    response.Version,
				Index:      r.PartitionIndex,
				ErrorCode:  int16(r.Error),
				BaseOffset: int64(r.BaseOffset),
				// TODO: this needs to be implemented, see documentation for details
				LogAppendTimeMs: -1,
				LogStartOffset:  0,
				// TODO: Don't forget to handle errors when the protocol is fully implemented
			}
		}

		topicResponses = append(response.Responses, protocol.TopicProduceResponse{
			Name:               topic,
			PartitionResponses: partResponses,
		})
	}

	response.Responses = topicResponses

	resp, err := protocol.Encode(&response)
	return resp, response.GetHeaderVersion(), err
}
