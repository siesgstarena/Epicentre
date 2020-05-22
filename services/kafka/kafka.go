package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
	"time"
)

func config() *sarama.Config {
    config := sarama.NewConfig()
    config.Net.DialTimeout = 10 * time.Second
    config.Net.SASL.Enable = true
    config.Net.SASL.User = mainConfig.Config.KafkaUsername
    config.Net.SASL.Password = mainConfig.Config.KafkaPassword
    config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Return.Successes = true
    config.Version = sarama.V1_1_0_0
    return config
}