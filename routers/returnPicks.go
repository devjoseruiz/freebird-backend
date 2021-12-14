package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*ReturnPicks returns the user timeline*/
func ReturnPicks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")

	if len(ID) == 0 {
		http.Error(w, "ID is mandatory", http.StatusBadRequest)
		return
	}

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

	pageInt64 := int64(pageNumber)

	response, status := db.ShowPicks(ID, pageInt64)

	if status == false {
		http.Error(w, "Error at reading picks", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
