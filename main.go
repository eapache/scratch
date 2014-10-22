package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	client, err := sarama.NewClient("client_id", []string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> connected")
	}
	defer client.Close()

	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	sarama.MaxRequestSize = 37 * 1024 * 1024
	producerConfig := sarama.NewProducerConfig()
	producerConfig.MaxMessageBytes = 1000000
	producerConfig.Compression = sarama.CompressionSnappy
	producer, err := sarama.NewProducer(client, producerConfig)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for {
		select {
		case producer.Input() <- &sarama.MessageToSend{Topic: "my_topic", Key: nil, Value: sarama.ByteEncoder(make([]byte, 1024))}:
		case err := <-producer.Errors():
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
