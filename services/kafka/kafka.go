package kafka

import (
	"fmt"
	"os"
    "os/signal"
    "syscall"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

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
    topic := fmt.Sprintf("%sdefault", kafkaConfiguration.Config.KafkaTopicPrefix)
    p, err := kafka.NewProducer(config)
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
		return err
    }
    fmt.Printf("Created Producer %v\n", p)
	deliveryChan := make(chan kafka.Event)
	fmt.Println(deliveryChan);
	  
	sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    c, err := kafka.NewConsumer(config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
		return err
    }
    fmt.Printf("Created Consumer %v\n", c)
    err = c.Subscribe(topic, nil)

	return nil
}