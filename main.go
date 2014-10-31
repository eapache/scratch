package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/Shopify/sarama"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	f, err := os.Create("profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	clientConfig := sarama.NewClientConfig()
	client, err := sarama.NewClient("client_id", []string{"192.168.100.67:6667"}, clientConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> connected")
	}
	defer client.Close()

	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	producerConfig := sarama.NewProducerConfig()
	producerConfig.Compression = sarama.CompressionSnappy
	producer, err := sarama.NewProducer(client, producerConfig)
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println(producer.Close())
	}()

	i := 0
	for {
		select {
		// default:
		case producer.Input() <- &sarama.MessageToSend{Topic: "my_topic", Key: nil, Value: sarama.ByteEncoder(make([]byte, 1024))}:
			// producer.QueueMessage("my_topic", nil, sarama.ByteEncoder(make([]byte, 1024)))
			i++
			if i >= 10000 {
				fmt.Println("batch")
				i = 0
			}
		case err := <-producer.Errors():
			fmt.Println(err)
		case <-sig:
			return
		}
	}
}
