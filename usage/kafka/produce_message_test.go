package kafka

import (
	"fmt"
	"testing"
	"github.com/sarama"
)

func TestProduceMessage(t *testing.T) {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close();err != nil {
			fmt.Printf("producer close err: %s\n", err.Error())
		}
	}()

	msg := &sarama.ProducerMessage{Topic:"test", Value: sarama.StringEncoder("This is a message")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("send message err: %s\n", err.Error())
	} else {
		fmt.Printf("send message successful! partition %d at offset %d\n", partition, offset)
	}
}
