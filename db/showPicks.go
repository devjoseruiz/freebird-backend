package db

import (
	"context"
	"log"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ShowPicks returns the user timeline*/
func ShowPicks(ID string, page int64) ([]*models.ReturnPicks, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	picksCollection := dbConn.Collection("picks")

	var result []*models.ReturnPicks

	conditional := bson.M{"userid": ID}

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)

	cursor, err := picksCollection.Find(ctx, conditional, findOptions)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var document models.ReturnPicks

		err := cursor.Decode(&document)

		if err != nil {
			return result, false
		}

		result = append(result, &document)
	}

	return result, true
}
