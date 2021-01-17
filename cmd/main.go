package main

import (
	"github.com/paroar/roadtracing-kafka-mongo/internal/kafka"
	"github.com/paroar/roadtracing-kafka-mongo/internal/mongo"
)

func main() {

	mongo.InitialiseMongo()

	kafka.NewConsumer()
	kafka.ReceiveFromKafka()
}
