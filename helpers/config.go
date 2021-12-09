package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*GetConfVar loads the .env file and returns the value of a given key*/
func GetConfVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
