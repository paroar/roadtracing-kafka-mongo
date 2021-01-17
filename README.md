# Kafka-mongo

Kafka-mongo is a kafka consumer container ready to insert into mongo database.

## Configuration

Required environment variables and their default values.

```docker-compose
KAFKA_SERVER=kafka:9092
KAFKA_TOPIC=topic
MONGO_HOST=mongodb:27017
MONGO_DB=db
MONGO_COLLECTION=collection
MONGO_USERNAME=username
MONGO_PASSWORD=password
```

## Notes

The authorization check is on admin db.