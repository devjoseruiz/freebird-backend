package helpers

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword encrypts a given password using adaptive hashing algorithm*/
func EncryptPassword(passw string) (string, error) {
	cost := 16
	bytes, err := bcrypt.GenerateFromPassword([]byte(passw), cost)

	return string(bytes), err
}
