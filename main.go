package main

import (
	"log"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/handlers"
)

func main() {
	if db.CheckConn() == false {
		log.Fatal("Error at connecting to DB")
		return
	}

	handlers.Handlers()
}
