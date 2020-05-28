package kafka

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
)

// Health type for kafka
type Health struct {
	Name     		string `json:"name"`
	Description     string `json:"description"`
	Health 			string `json:"health"`
	Timestamp       time.Time `json:"timestamp"`
}

// ProduceMessage This function can be used to send message on the topic
func ProduceMessage(health Health) error  {
	message, _ := json.Marshal(health)
	topic := fmt.Sprintf("%sdefault", kafkaConfiguration.Config.KafkaTopicPrefix)
	deliveryChan := make(chan kafka.Event)
	err := Producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(message)}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(deliveryChan)
	return nil
}