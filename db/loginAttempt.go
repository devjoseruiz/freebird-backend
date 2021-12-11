package db

import (
	"gitlab.com/joseruizdev/freebird-backend/models"
	"golang.org/x/crypto/bcrypt"
)

/*LoginAttempt compares a provided password with the password stored in DB*/
func LoginAttempt(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfUserExists(email)

	if found == false {
		return user, false
	}

	userPassword := []byte(password)
	passwordFromDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordFromDB, userPassword)

	if err != nil {
		return user, false
	}

	return user, true
}
