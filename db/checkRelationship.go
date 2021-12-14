package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckRelationship checks a relation between two users*/
func CheckRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	relationshipsCollection := dbConn.Collection("relationships")

	conditional := bson.M{
		"userid":          t.UserID,
		"followinguserid": t.FollowingUserID,
	}

	var result models.Relationship

	err := relationshipsCollection.FindOne(ctx, conditional).Decode(&result)

	if err != nil {
		return false, err
	}

	return true, nil
}
