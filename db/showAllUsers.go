package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/models"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ShowAllUsers returns all users related with current user*/
func ShowAllUsers(ID string, page int64, search string, searchType string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	usersCollection := dbConn.Collection("users")

	var result []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := usersCollection.Find(ctx, query, findOptions)

	if err != nil {
		return result, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)

		if err != nil {
			return result, false
		}

		var r models.Relationship
		r.UserID = ID
		r.FollowingUserID = s.ID.Hex()

		include = false

		found, err = CheckRelationship(r)

		if searchType == "new" && found == false {
			include = true
		}

		if searchType == "follow" && found == true {
			include = true
		}

		if r.FollowingUserID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Email = ""
			s.Bio = ""
			s.Web = ""
			s.Location = ""
			s.Banner = ""

			result = append(result, &s)
		}
	}

	err = cursor.Err()

	if err != nil {
		return result, false
	}

	cursor.Close(ctx)
	return result, true
}
