package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gitlab.com/joseruizdev/freebird-backend/helpers"
)

/*Handlers set the app port and listens to it*/
func Handlers() {
	router := mux.NewRouter()

	PORT := helpers.GetConfVar("HOST_PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
