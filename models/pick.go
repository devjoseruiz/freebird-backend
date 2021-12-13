package models

/*Pick captures from Body the provided message*/
type Pick struct {
	Message string `bson:"message" json:"message"`
}
