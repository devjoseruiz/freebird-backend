package models

/*Relationship set a relation between an user and the users that he follows*/
type Relationship struct {
	UserID          string `bson:"userid" json:"userId"`
	FollowingUserID string `bson:"followinguserid" json:"followingUserId"`
}
