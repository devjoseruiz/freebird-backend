package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*PickInsert saves a pick in the DB*/
func PickInsert(t models.SavePick) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	picksCollection := dbConn.Collection("picks")

	pick := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := picksCollection.InsertOne(ctx, pick)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
