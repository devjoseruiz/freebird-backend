package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckIfUserExists checks if an user exists in the DB*/
func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	usersCollection := dbConn.Collection("users")

	conditional := bson.M{"email": email}

	var result models.User

	objID := result.ID.Hex()
	err := usersCollection.FindOne(ctx, conditional).Decode(&result)

	if err != nil {
		return result, false, objID
	}

	return result, true, objID
}
