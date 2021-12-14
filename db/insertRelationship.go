package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*InsertRelationship creates a relation between an user and the users he follows*/
func InsertRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	relationshipsCollection := dbConn.Collection("relationships")

	_, err := relationshipsCollection.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
