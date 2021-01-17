package kafka

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/paroar/roadtracing-kafka-mongo/internal/mongo"
)

var (
	hostKafka = os.Getenv("KAFKA_SERVER")
	topic     = os.Getenv("KAFKA_TOPIC")
	consumer  *kafka.Consumer
)

// NewConsumer creates a new kafka consumer
func NewConsumer() {

	for {
		fmt.Println("Creating Consumer")
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": hostKafka,
			"group.id":          "group-id-1",
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Consumer created")
			consumer = c
			break
		}
	}
}

// ReceiveFromKafka starts kafka consumer
func ReceiveFromKafka() {

	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			job := string(msg.Value)
			mongo.SavePositionToMongo(job)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
}
