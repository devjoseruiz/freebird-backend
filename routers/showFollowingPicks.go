package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*ShowFollowingPicks return picks of all followed users*/
func ShowFollowingPicks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	page := r.URL.Query().Get("page")
	if len(page) == 0 {
		http.Error(w, "Page is mandatory", http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Page must be a number and greater than zero", http.StatusBadRequest)
		return
	}

	result, fine := db.ShowFollowingPicks(IDUser, pageNumber)

	if fine == false {
		http.Error(w, "Error reading picks", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
