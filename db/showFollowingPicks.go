package db

import (
	"context"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ShowFollowingPicks return picks from the followed users*/
func ShowFollowingPicks(ID string, page int) ([]models.ReturnFollowingPicks, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbName := helpers.GetConfVar("DATABASE_NAME")
	dbConn := MongoConn.Database(dbName)
	relationshipsCollection := dbConn.Collection("relationships")

	skip := (page - 1) * 20

	conditional := make([]bson.M, 0)

	conditional = append(conditional, bson.M{"$match": bson.M{"userid": ID}})
	conditional = append(conditional, bson.M{
		"$lookup": bson.M{
			"from":         "picks",
			"localField":   "followinguserid",
			"foreignField": "userid",
			"as":           "pick",
		},
	})

	conditional = append(conditional, bson.M{"$unwind": "$pick"})
	conditional = append(conditional, bson.M{"$sort": bson.M{"pick.date": -1}})
	conditional = append(conditional, bson.M{"$skip": skip})
	conditional = append(conditional, bson.M{"$limit": 20})

	cursor, err := relationshipsCollection.Aggregate(ctx, conditional)
	var result []models.ReturnFollowingPicks
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
