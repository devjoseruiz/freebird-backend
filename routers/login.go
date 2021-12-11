package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/jwt"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*Login performs user sign-in*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User and/or password invalid: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email field is mandatory", 400)
		return
	}

	user, exists := db.LoginAttempt(t.Email, t.Password)

	if exists == false {
		http.Error(w, "User and/or password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)

	if err != nil {
		http.Error(w, "Error attempting to generate a token: "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{Token: jwtKey}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
