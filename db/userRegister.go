package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UserRegister insert a document in MongoDB for a new registered user*/
func UserRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	usersCollection := dbConn.Collection("users")

	u.Password, _ = helpers.EncryptPassword(u.Password)

	result, err := usersCollection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
