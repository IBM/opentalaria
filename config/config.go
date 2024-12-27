package config

type Config struct {
	Broker  *Broker
	Cluster *Cluster
}

type Broker struct {
	BrokerID int32
	Rack     *string
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listeners
	Listeners []Listener
	// https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#advertised-listeners
	AdvertisedListeners []Listener
}

type Listener struct {
	Host string
	Port int32
	// If the listener name is a security protocol, like PLAINTEXT,SSL,SASL_PLAINTEXT,SASL_SSL,
	// the name will be set as SecurityProtocol. Otherwise the name should be mapped in listener.security.protocol.map.
	// see https://docs.confluent.io/platform/current/installation/configuration/broker-configs.html#listener-security-protocol-map.
	SecurityProtocol SecurityProtocol
	ListenerName     string
}

type Cluster struct {
	ClusterID string
}
