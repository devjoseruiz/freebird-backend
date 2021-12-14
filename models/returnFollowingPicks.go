package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReturnFollowingPicks is the structure in what picks are returned*/
type ReturnFollowingPicks struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID          string             `bson:"userid" json:"userId,omitempty"`
	FollowingUserID string             `bson:"followinguserid" json:"followingUserId"`
	Pick            struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
	}
}
