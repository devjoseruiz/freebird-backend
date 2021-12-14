package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*RemoveRelationship remove a relationship in the DB*/
func RemoveRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	relationshipsCollection := dbConn.Collection("relationships")

	_, err := relationshipsCollection.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
