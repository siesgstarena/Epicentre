package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
	"time"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"crypto/tls"
)

// Producer producer
type Producer struct {
	AsyncProducer sarama.AsyncProducer
}

//Load Kafka
func Load(){
	brokersList:= mainConfig.Config.KafkaBrokerList
	brokers:=strings.Split(brokersList, ",")
	topic := mainConfig.Config.KafkaTopicPrefix + "default"

    client, err := sarama.NewClient(brokers, Config())
    if err != nil {
        panic(err)
    }
    fmt.Println("New Client")

	producer := &Producer{
	   AsyncProducer: GetAsyncProducer(),
	}
	ProduceMessageAsync(producer)
	consumer, err := sarama.NewConsumer(brokers, Config())
	if err != nil {
		panic(err)
	}
	defer producer.AsyncProducer.Close() // handle error yourself
	defer func() {
        if err := consumer.Close(); err != nil {
            panic(err)
        }
	}()
	
    partitions, err := client.Partitions(topic)
    if err != nil {
        panic(err)
    }

    messages := make(chan *sarama.ConsumerMessage)
    for _, part := range partitions {
        go consumePartition(consumer, topic, part, messages)
    }

    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    consumed := 0
	ConsumerLoop: 
	for {
		select {
		case msg := <-messages:
			fmt.Printf("Consumed message offset %d\n", msg.Offset)
			fmt.Println(msg.Value)
			consumed++
		case <-signals:
			fmt.Println("break")
			break ConsumerLoop
		}
	}

    fmt.Printf("Consumed: %d\n", consumed)
}

func consumePartition(consumer sarama.Consumer, topic string, partition int32, out chan *sarama.ConsumerMessage) {
    partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := partitionConsumer.Close(); err != nil {
            panic(err)
        }
    }()

    for msg := range partitionConsumer.Messages() {
        out <- msg
    }
} 

// Config Returns config needed for kafka
func Config() *sarama.Config {
    config := sarama.NewConfig()
    config.Net.DialTimeout = 10 * time.Second
    config.Net.SASL.Enable = true
    config.Net.SASL.User = mainConfig.Config.KafkaUsername
    config.Net.SASL.Password = mainConfig.Config.KafkaPassword
    config.Net.TLS.Enable = true
    config.Net.TLS.Config = &tls.Config{
        InsecureSkipVerify: true,
        ClientAuth:         0,
	}
    // config.Version = sarama.V1_0_0_0
    return config
}