# Overview
https://cwiki.apache.org/confluence/display/KAFKA/A+Guide+To+The+Kafka+Protocol#AGuideToTheKafkaProtocol-MetadataAPI

## Broker config
The implementation follows Kafka's [specifications](https://kafka.apache.org/documentation/#brokerconfigs_advertised.listeners) as close as possible. 

If advertised listeners are not set, the Metadata API returns the listeners list. If a listener has an empty hostname, OpenTalaria will return the IPv4 address of the first network interface of the host and print the IP address in an INFO log statement.

OpenTalaria does not support more than one listener, please see [#18](https://github.com/IBM/opentalaria/issues/18) for updates.

## Cluster config
if an environment variable `KAFKA_CLUSTER_ID` is not set OpenTalaria will generate a random UUID and set this as the cluster ID every time the process is restarted. 

# Caveats
* For now a rack ID cannot be assigned to the broker, since OpenTalaria will not support distributed clusters in the initial release. This will be added at later stages. To satisfy the protocol, the API will always return the rack as null.