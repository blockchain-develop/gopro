package kafka

import (
	"fmt"
	"github.com/sarama"
	"os"
	"os/signal"
	"testing"
)

func TestConsumeMessage(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close();err != nil {
			fmt.Printf("consumer close err: %s\n", err.Error())
		}
	}()

	partition, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partition.Close(); err != nil {
			fmt.Printf("partition close err: %s\n", err.Error())
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumerLoop:
	for {
		select {
		case msg := <- partition.Messages():
			fmt.Printf("consumed message offset: %d, message: %s\n", msg.Offset, string(msg.Value))
		case <- signals:
			break consumerLoop
		}
	}
}
