package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
)

// NewProducer Kafka
func NewProducer() (sarama.SyncProducer, error) {
	brokers:= mainConfig.Config.KafkaBrokerList
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