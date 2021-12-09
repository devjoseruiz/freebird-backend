package db

import (
	"context"
	"log"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDbURI = helpers.GetConfVar("MONGODB_URI")
var clientOptions = options.Client().ApplyURI(mongoDbURI)

/*MongoConn is the MongoDB connection object*/
var MongoConn = Connect()

/*Connect is a function that allows to stablish a connection with MongoDB*/
func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection to DB was successful")

	return client
}

/*CheckConn checks the connection to the MongoDB server*/
func CheckConn() int {
	err := MongoConn.Ping(context.TODO(), nil)

	if err != nil {
		return 1
	}

	return 0
}
