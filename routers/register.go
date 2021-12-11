package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*Register is a function for registering new users into the system*/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid data: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email field is mandatory", 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "Password lenght must be equal "+
			"or greater than 8 characters", 400)
		return
	}

	_, found, _ := db.CheckIfUserExists(t.Email)

	if found == true {
		http.Error(w, "User already exists", 400)
		return
	}

	_, status, err := db.UserRegister(t)

	if err != nil {
		http.Error(w, "Error at user registering: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Couldn't complete user registering", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
