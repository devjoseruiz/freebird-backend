package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*Email value contained in Email for all the endpoints*/
var Email string

/*IDUser is the ID returned by the model*/
var IDUser string

/*ProcessToken process the Token for extracting its content*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	var envPrivate = helpers.GetConfVar("JWT_PRIVATE")
	privateKey := []byte(envPrivate)

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid Token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkParsed, err := jwt.ParseWithClaims(tk, claims,
		func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

	if err == nil {
		_, found, _ := db.CheckIfUserExists(claims.Email)

		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, found, IDUser, nil
	}

	if !tkParsed.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
