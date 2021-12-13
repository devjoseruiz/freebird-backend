package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ModifyProfile allows to modify an own user profile*/
func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	usersCollection := dbConn.Collection("users")

	userData := make(map[string]interface{})

	if len(u.Name) > 0 {
		userData["name"] = u.Name
	}

	if len(u.Surname) > 0 {
		userData["surname"] = u.Surname
	}

	userData["birth"] = u.Birth

	if len(u.Avatar) > 0 {
		userData["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		userData["banner"] = u.Banner
	}

	if len(u.Bio) > 0 {
		userData["bio"] = u.Bio
	}

	if len(u.Location) > 0 {
		userData["location"] = u.Location
	}

	if len(u.Web) > 0 {
		userData["web"] = u.Web
	}

	updatedUserData := bson.M{"$set": userData}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := usersCollection.UpdateOne(ctx, filter, updatedUserData)

	if err != nil {
		return false, err
	}

	return true, nil
}
