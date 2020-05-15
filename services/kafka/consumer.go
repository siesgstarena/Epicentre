package kafka

import (
	"fmt"
	"os"
    "os/signal"
    "syscall"
	kafkaConfiguration "github.com/siesgstarena/epicentre/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ConsumeMessage This function can be used to receive message on the topic
func ConsumeMessage() error  {
	topic := fmt.Sprintf("%sdefault", kafkaConfiguration.Config.KafkaTopicPrefix)
	sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	err := Consumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}
    run := true
    counter := 0
    commitAfter := 1000
    for run == true {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        case ev := <-Consumer.Events():
            switch e := ev.(type) {
            case kafka.AssignedPartitions:
                Consumer.Assign(e.Partitions)
            case kafka.RevokedPartitions:
                Consumer.Unassign()
            case *kafka.Message:
                fmt.Printf("%% Message on %s: %s\n", e.TopicPartition, string(e.Value))
                counter++
                if counter > commitAfter {
                    Consumer.Commit()
                    counter = 0
                }
            case kafka.PartitionEOF:
                fmt.Printf("%% Reached %v\n", e)
            case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
				return e
            }
        }
    }
    fmt.Printf("Closing consumer\n")
	Consumer.Close()
	return nil
}