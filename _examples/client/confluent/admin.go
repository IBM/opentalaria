package main

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaConfig *kafka.ConfigMap
	admin       *kafka.AdminClient
	topicName   = "test-topic"
)

func init() {
	kafkaConfig = &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "test-client",
		"acks":              "all"}

	var err error
	admin, err = kafka.NewAdminClient(kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// get metadata
	defer admin.Close()

	// create topics
	topics := make([]kafka.TopicSpecification, 1)
	topics[0].Topic = topicName
	topics[0].NumPartitions = 1
	topics[0].ReplicationFactor = 1

	result, err := admin.CreateTopics(context.Background(), topics)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", result)

	// get metadata for topic
	metadata, err := admin.GetMetadata(&topicName, false, 1000)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", metadata)

}
