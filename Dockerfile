FROM golang:1.15.6

WORKDIR /app

COPY . .

RUN go mod vendor
RUN go build -o ./bin/kafka-mongo ./cmd/main.go

ENV KAFKA_SERVER=kafka:9092 \
    KAFKA_TOPIC=topic \
    MONGO_HOST=mongodb:27017 \
    MONGO_DB=db \
    MONGO_COLLECTION=collection \
    MONGO_USERNAME=username \
    MONGO_PASSWORD=password

CMD [ "./bin/kafka-mongo" ]