package utils

type ConfigResourceType int8

// This enum is used in the DescribeConfigs API
const (
	BROKER_CONFIG_TYPE ConfigResourceType = iota
	BROKER_LOGGER_CONFIG_TYPE
	TOPIC_CONFIG_TYPE
	UNKNOWN_CONFIG_TYPE
)
