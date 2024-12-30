package config

type Config struct {
	Broker  *Broker
	Cluster *Cluster
}

type Cluster struct {
	ClusterID string
}

/**
 * Unit test helpers
 */

// MockCluster generates a mock object used for unit testing.
func MockCluster() *Cluster {
	return &Cluster{
		ClusterID: "abc",
	}
}

// MockConfig generates a mock object used for unit testing.
func MockConfig() *Config {
	config := Config{}

	config.Cluster = MockCluster()
	config.Broker = MockBroker()

	return &config
}
