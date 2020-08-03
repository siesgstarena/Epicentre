package kafka

import (
	"encoding/json"
	"fmt"
	// "time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
)

// GithubData can be parsed using this type
type GithubData struct {
	Action		string		`json:"action,omitempty"`
	Issue		issue		`json:"issue,omitempty"`
	Repository	repository	`json:"repository,omitempty"`
	Sender		sender		`json:"sender,omitempty"`
}

type issue struct {
	URL		string	`json:"url,omitempty"`
	Number	int		`json:"number,omitempty"`
}

type repository struct {
	ID			int		`json:"id,omitempty"`
	FullName	string	`json:"full_name,omitempty"`
}

type sender struct {
	Login	string	`json:"login,omitempty"`
}

// ProduceMessage This function can be used to send message on the topic
func ProduceMessage(data GithubData) error  {
	message, _ := json.Marshal(data)
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