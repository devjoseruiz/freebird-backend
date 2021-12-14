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

	router.HandleFunc("/profile/edit",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.EditProfile))).Methods("PUT")

	router.HandleFunc("/pick",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.SavePick))).Methods("POST")

	router.HandleFunc("/pick",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.ReturnPicks))).Methods("GET")

	router.HandleFunc("/pick",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.RemovePick))).Methods("DELETE")

	router.HandleFunc("/profile/avatar",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")

	router.HandleFunc("/profile/banner",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")

	router.HandleFunc("/profile/avatar",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("GET")

	router.HandleFunc("/profile/banner",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("GET")

	router.HandleFunc("/follow",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.SetRelationship))).Methods("POST")

	router.HandleFunc("/follow",
		middlewares.CheckDB(middlewares.ValidateJWT(routers.UnsetRelationship))).Methods("DELETE")

	PORT := helpers.GetConfVar("HOST_PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
