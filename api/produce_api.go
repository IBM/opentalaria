package api

import (
	"fmt"
	"log/slog"

	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
)

func HandleProduceRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	produceRequest := protocol.ProduceRequest{}
	_, err := protocol.VersionedDecode(req.Message, &produceRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.ProduceResponse{}

	response.Version = req.Header.RequestApiVersion

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	for _, topic := range produceRequest.TopicData {
		topicResponse := protocol.TopicProduceResponse{}
		topicResponse.Version = response.Version
		topicResponse.Name = topic.Name

		for _, partition := range topic.PartitionData {
			slog.Debug("Received records", "records", fmt.Sprintf("%+v", partition.Records))

			topicResponse.PartitionResponses = append(topicResponse.PartitionResponses, protocol.PartitionProduceResponse{
				Version:    response.Version,
				Index:      partition.Index,
				ErrorCode:  int16(utils.ErrNoError),
				BaseOffset: partition.Records.BaseOffset,
				// TODO: this needs to be implemented, see documentation for details
				LogAppendTimeMs: -1,
				LogStartOffset:  0,
				// TODO: Don't forget to handle errors when the protocol is fully implemented
			})
		}

		response.Responses = append(response.Responses, topicResponse)
	}

	resp, err := protocol.Encode(&response)
	return resp, response.GetHeaderVersion(), err
}
