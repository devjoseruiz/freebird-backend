package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*EditProfile modifies the user profile*/
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid data: "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyProfile(t, IDUser)

	if err != nil {
		http.Error(w, "Error at attempting to modify data: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Profile couldn't be updated", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
