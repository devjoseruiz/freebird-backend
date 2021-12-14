package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*ListUsers return an users list*/
func ListUsers(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Page number must be integer and greater than zero", http.StatusBadRequest)
		return
	}

	pageNumber := int64(pageTemp)

	result, status := db.ShowAllUsers(IDUser, pageNumber, search, userType)

	if status == false {
		http.Error(w, "Error searching users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
