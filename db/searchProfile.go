package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*SearchProfile searchs for an user profile in the DB*/
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	usersCollection := dbConn.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	conditional := bson.M{"_id": objID}

	err := usersCollection.FindOne(ctx, conditional).Decode(&profile)

	if err != nil {
		return profile, err
	}

	profile.Password = ""

	return profile, nil
}
