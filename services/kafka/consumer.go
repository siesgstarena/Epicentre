package kafka

import (
	"fmt"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
)

// ConsumeMessage This function can be used to receive message on the topic
func ConsumeMessage()  {
	topic := fmt.Sprintf("%sdefault", kafkaConfiguration.Config.KafkaTopicPrefix)
	err := Consumer.Subscribe(topic, nil)
	if err != nil {
        panic(err)
	}
    for {
		msg, err := Consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}