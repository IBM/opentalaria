package api

// // TODO: this is a placeholder function for now. We need to implement a backend that handles cluster topology in order to implement the API correctly and consume the messages.
// func (p ProduceAPI) GeneratePayload() ([]byte, error) {
// 	req := protocol.ProduceRequest{}
// 	_, err := protocol.VersionedDecode(p.GetRequest().Message, &req, p.GetRequest().Header.RequestApiVersion)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := protocol.ProduceResponse{
// 		Version: p.GetRequest().Header.RequestApiVersion,
// 	}

// 	for _, topic := range req.TopicData {
// 		topicResponse := protocol.TopicProduceResponse{}
// 		topicResponse.Version = resp.Version
// 		topicResponse.Name = topic.Name

// 		for _, partition := range topic.PartitionData {
// 			slog.Debug("Received records", "records", partition.Records)

// 			topicResponse.PartitionResponses = append(topicResponse.PartitionResponses, protocol.PartitionProduceResponse{
// 				Version:    resp.Version,
// 				Index:      partition.Index,
// 				ErrorCode:  int16(utils.ErrNoError),
// 				BaseOffset: partition.Records.BaseOffset,
// 				// TODO: this needs to be implemented, see documentation for details
// 				LogAppendTimeMs: -1,
// 				LogStartOffset:  0,
// 				// TODO: Don't forget to handle errors when the protocol is fully implemented
// 			})
// 		}

// 		resp.Responses = append(resp.Responses, topicResponse)
// 	}

// 	return protocol.Encode(&resp)
// }
