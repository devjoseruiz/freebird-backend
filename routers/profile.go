package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*Profile returns the content of an user profile*/
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ID := r.URL.Query().Get("id")

	if len(ID) == 0 {
		http.Error(w, "User ID is mandatory", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "Error at attepting to search the user: "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
