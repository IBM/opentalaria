package api

import (
	"log/slog"

	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/protocol"
)

func HandleMetadataRequest(req config.Request, apiVersion int16, opts ...any) ([]byte, int16, error) {
	metadataRequest := protocol.MetadataRequest{}
	_, err := protocol.VersionedDecode(req.Message, &metadataRequest, req.Header.RequestApiVersion)
	if err != nil {
		return nil, 0, err
	}

	response := protocol.MetadataResponse{}

	response.Version = req.Header.RequestApiVersion

	// TODO: handle throttle time
	response.ThrottleTimeMs = 0

	// TODO: we will have to handle multiple advertised listeners, this implementation is very naive and assumes OpenTalaria won't be run in cluster mode
	// Since cluster mode is not supported for now, we take the first AdvertisedListener as broker config.
	listener := req.Config.Broker.AdvertisedListeners[0]
	response.Brokers = append(response.Brokers, protocol.MetadataResponseBroker{
		NodeID: req.Config.Broker.BrokerID,
		Host:   listener.Host,
		Port:   listener.Port,
		Rack:   nil, // for now OpenTalaria does not support rack awareness.
	})

	response.ClusterID = &req.Config.Cluster.ClusterID
	response.ControllerID = req.Config.Broker.BrokerID

	var topicNames []string

	for i, topic := range metadataRequest.Topics {
		if topicNames == nil {
			topicNames = make([]string, len(metadataRequest.Topics))
		}

		topicNames[i] = *topic.Name
	}

	topics, err := req.Config.Plugin.ListTopics(topicNames)
	if err != nil {
		slog.Error("error listing topics", "err", err)
	}

	response.Topics = topics
	response.ClusterAuthorizedOperations = 0

	resp, err := protocol.Encode(&response)

	return resp, response.GetHeaderVersion(), err
}
