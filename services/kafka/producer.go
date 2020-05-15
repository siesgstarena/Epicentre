package kafka

import (
	"fmt"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProduceMessage This function can be used to send message on the topic
func ProduceMessage(message string) error  {
	topic := fmt.Sprintf("%sdefault", kafkaConfiguration.Config.KafkaTopicPrefix)
	deliveryChan := make(chan kafka.Event)
	fmt.Println(message)
    for i := 0; i < 10; i++ {
        value := fmt.Sprintf("[%d] %s", i+1, message)
		err := Producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}, deliveryChan)
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
		fmt.Println(value)
	}
	close(deliveryChan)
	return nil
}