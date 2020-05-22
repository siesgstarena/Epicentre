package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
    "strings"
)

// NewProducer Kafka
func NewProducer() (sarama.SyncProducer, error) {
	brokersList:= mainConfig.Config.KafkaBrokerList
	brokers:=strings.Split(brokersList, ",")
    producer, err := sarama.NewSyncProducer(brokers, config())
	return producer, err
}

// PrepareMessage Kafka
func PrepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}