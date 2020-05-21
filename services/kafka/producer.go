package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
	"fmt"
    "strings"
)

// ProduceMessageAsync Produces message
func ProduceMessageAsync(producer *Producer){

	topic := mainConfig.Config.KafkaTopicPrefix + "default"
	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	message := &sarama.ProducerMessage{
	   Topic: topic,
	   Value: sarama.StringEncoder("Hello Go! Async Message"),
	}
	producer.AsyncProducer.Input() <- message
} 

// GetAsyncProducer returns producer
func GetAsyncProducer() sarama.AsyncProducer {

	brokers:= mainConfig.Config.KafkaBrokerList
	brokerList:=strings.Split(brokers, ",")
 
	producer, err := sarama.NewAsyncProducer(brokerList, Config())
	if err != nil {
	   fmt.Println("Failed to start Sarama producer:", err)
	}
	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
	   for err := range producer.Errors() {
		fmt.Println("Failed to write access log entry:", err)
	   }
	}()
 
	return producer
 }