package mongo

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/paroar/roadtracing-kafka-mongo/internal/types"
	"gopkg.in/mgo.v2"
)

type mongoStoreSession struct {
	session *mgo.Session
}

var mongoConnection = mongoStoreSession{}

var (
	hosts      = os.Getenv("MONGO_HOST")
	database   = os.Getenv("MONGO_DB")
	collection = os.Getenv("MONGO_COLLECTION")
	username   = os.Getenv("MONGO_USERNAME")
	password   = os.Getenv("MONGO_PASSWORD")
)

// InitialiseMongo initial connection to mongo
func InitialiseMongo() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: "admin",
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	mongoConnection.session = session
}

// SavePositionToMongo saves a Position into mongo
func SavePositionToMongo(posString string) {

	col := mongoConnection.session.DB(database).C(collection)

	var _pos types.Position
	b := []byte(posString)
	err := json.Unmarshal(b, &_pos)
	if err != nil {
		panic(err)
	}

	errMongo := col.Insert(_pos)
	if errMongo != nil {
		panic(errMongo)
	}

	log.Printf("Saved to MongoDB : %s", posString)
}
