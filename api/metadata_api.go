package api

// func (m MetadataAPI) GeneratePayload() ([]byte, error) {
// 	req := protocol.MetadataRequest{}
// 	_, err := protocol.VersionedDecode(m.GetRequest().Message, &req, m.GetRequest().Header.RequestApiVersion)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := m.GenerateMetadataResponse()
// 	return protocol.Encode(response)
// }

// func (m MetadataAPI) GenerateMetadataResponse() *protocol.MetadataResponse {
// 	// For now the returned data is mock, just so we can continue developing the rest of the APIs.
// 	// Once we have a more robust project architecture, this struct will be populated with the real
// 	// cluster metadata.
// 	response := protocol.MetadataResponse{}

// 	response.Version = m.GetRequest().Header.RequestApiVersion
// 	// TODO: handle throttle time
// 	response.ThrottleTimeMs = 0

// 	// TODO: we will have to handle multiple advertised listeners, this implementation is very naive and assumes OpenTalaria won't be run in cluster mode
// 	// Since cluster mode is not supported for now, we take the first AdvertisedListener as broker config.
// 	listener := m.Config.Broker.AdvertisedListeners[0]
// 	response.Brokers = append(response.Brokers, protocol.MetadataResponseBroker{
// 		NodeID: m.Config.Broker.BrokerID,
// 		Host:   listener.Host,
// 		Port:   listener.Port,
// 		Rack:   nil, // for now OpenTalaria does not support rack awareness.
// 	})

// 	response.ClusterID = &m.Config.Cluster.ClusterID
// 	response.ControllerID = m.Config.Broker.BrokerID

// 	topics, err := m.Config.Plugin.ListTopics()
// 	if err != nil {
// 		slog.Error("error listing topics", "err", err)
// 	}

// 	response.Topics = topics
// 	response.ClusterAuthorizedOperations = 0

// 	return &response
// }
