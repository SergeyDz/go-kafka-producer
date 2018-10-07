package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SergeyDz/go-kafka-producer/config"
	model "github.com/SergeyDz/go-kafka-producer/model"
	"github.com/Shopify/sarama"
)

var (
	settings config.Config
)

func main() {

	// init settings
	settings = config.NewConfig()

	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	for {

		model := model.Metrics{Timestamp: time.Now().Format("2006-01-02 15:04:05"), Container: "fake", CPU: "50%", Memory: "33%"}
		msg, _ := json.Marshal(model)
		publish(string(msg), settings.Topic, producer)

		time.Sleep(settings.Timeout)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// sync producer
	prd, err := sarama.NewSyncProducer(strings.Split(settings.KafkaBrokers, ","), config)

	return prd, err
}

func publish(message string, topic string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
}
