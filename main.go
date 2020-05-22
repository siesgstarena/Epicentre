package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"github.com/siesgstarena/epicentre/config"
	"github.com/siesgstarena/epicentre/logger"
	routes "github.com/siesgstarena/epicentre/router"
	"github.com/siesgstarena/epicentre/services/kafka"
	"gopkg.in/Shopify/sarama.v1"
	"github.com/siesgstarena/epicentre/services/mongo"
)

func main() {

	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	err = logger.LoadLogger(*config.Config)
	if err != nil {
		panic(err)
	}
	logger.Log.Info("Logger Installed Successfully")

	brokersList:= config.Config.KafkaBrokerList
	brokers:=strings.Split(brokersList, ",")
	topic := config.Config.KafkaTopicPrefix + "default"

	producer, err := kafka.NewProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	kafka.Subscribe(topic, consumer)

	msg := kafka.PrepareMessage(topic, "Hello Sarama !")

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
	}

	mongo.LoadMongo()

	router := gin.Default()

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)

	// defer mongo.Client.Disconnect(*mongo.Ctx)
}
