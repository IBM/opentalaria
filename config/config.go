package config

import (
	"opentalaria/utils"

	"github.com/google/uuid"
)

type Config struct {
	Broker  *Broker
	Cluster *Cluster
}

type Cluster struct {
	ClusterID string
}

func NewConfig() (*Config, error) {
	config := Config{}

	broker, err := NewBroker()
	if err != nil {
		return &Config{}, err
	}

	config.Broker = broker

	clusterId, ok := utils.GetEnvVar("OT_CLUSTER_ID", "")
	if !ok {
		uid, err := uuid.NewV6()
		if err != nil {
			return &Config{}, err
		}

		clusterId = uid.String()
	}

	config.Cluster = &Cluster{
		ClusterID: clusterId,
	}

	return &config, nil
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
