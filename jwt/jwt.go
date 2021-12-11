package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*GenerateJWT generates an encrypted JWT*/
func GenerateJWT(t models.User) (string, error) {
	var envPrivate = helpers.GetConfVar("JWT_PRIVATE")
	privateKey := []byte(envPrivate)

	payload := jwt.MapClaims{
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24 * 14).Unix(),
		"email":    t.Email,
		"name":     t.Name,
		"surname":  t.Surname,
		"birth":    t.Birth,
		"bio":      t.Bio,
		"location": t.Location,
		"web":      t.Web,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
