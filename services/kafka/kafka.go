package kafka

import (
	"fmt"
	"os"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer for producing Messages
var Producer *kafka.Producer

// Consumer for consuming Messages
var Consumer *kafka.Consumer

// LoadKafka Configures Producer & Consumer with provided configuration
func LoadKafka() error {
	config := &kafka.ConfigMap{
        "metadata.broker.list": kafkaConfiguration.Config.KafkaBrokerList,
        "security.protocol":    "SASL_SSL",
        "sasl.mechanisms":      "SCRAM-SHA-256",
        "sasl.username":        kafkaConfiguration.Config.KafkaUsername,
        "sasl.password":        kafkaConfiguration.Config.KafkaPassword,
        "group.id":             kafkaConfiguration.Config.KafkaGroupID,
        "default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
    }
    
    p, err := kafka.NewProducer(config)
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
		return err
    }
    Producer = p
    fmt.Printf("Created Producer %v\n", p)
	  
    c, err := kafka.NewConsumer(config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
		return err
    }
    Consumer = c
    fmt.Printf("Created Consumer %v\n", c)

	return nil
}