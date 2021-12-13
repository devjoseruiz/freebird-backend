package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RemovePick allows to remove a pick*/
func RemovePick(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	picksCollection := dbConn.Collection("picks")

	objID, _ := primitive.ObjectIDFromHex(ID)

	conditional := bson.M{"_id": objID, "userid": UserID}

	_, err := picksCollection.DeleteOne(ctx, conditional)

	return err
}
