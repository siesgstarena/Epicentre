package kafka

import (
	"gopkg.in/Shopify/sarama.v1"
	mainConfig "github.com/siesgstarena/epicentre/config"
	"time"
	"crypto/tls"
)

func config() *sarama.Config {
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
    config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
    return config
}