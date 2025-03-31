package api

// func (m CreateTopicsAPI) GeneratePayload() ([]byte, error) {
// 	req := protocol.CreateTopicsRequest{}
// 	_, err := protocol.VersionedDecode(m.GetRequest().Message, &req, m.GetRequest().Header.RequestApiVersion)

// 	resp := m.GenerateCreateTopicsResponse(m.GetRequest().Header.RequestApiVersion, req, err)

// 	return protocol.Encode(resp)
// }

// func (m CreateTopicsAPI) GenerateCreateTopicsResponse(version int16, req protocol.CreateTopicsRequest, err error) *protocol.CreateTopicsResponse {
// 	response := protocol.CreateTopicsResponse{}

// 	response.Version = version
// 	// TODO: handle throttle time
// 	response.ThrottleTimeMs = 0

// 	for _, topic := range req.Topics {
// 		err := m.Config.Plugin.AddTopic(topic)

// 		response.Topics = append(response.Topics, protocol.CreatableTopicResult{
// 			Version:   req.Version,
// 			Name:      topic.Name,
// 			ErrorCode: int16(err),
// 		})
// 	}

// 	return &response
// }
