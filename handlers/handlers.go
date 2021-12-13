package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gitlab.com/joseruizdev/freebird-backend/helpers"
	"gitlab.com/joseruizdev/freebird-backend/middlewares"
	"gitlab.com/joseruizdev/freebird-backend/routers"
)

/*Handlers set the app port and listens to it*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register",
		middlewares.CheckDB(routers.Register)).Methods("POST")

	router.HandleFunc("/login",
		middlewares.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/profile",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.Profile))).Methods("GET")

	PORT := helpers.GetConfVar("HOST_PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
