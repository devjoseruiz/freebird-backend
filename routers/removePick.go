package routers

import (
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*RemovePick allows to remove a pick*/
func RemovePick(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")

	if len(ID) == 0 {
		http.Error(w, "ID is mandatory", http.StatusBadRequest)
		return
	}

	err := db.RemovePick(ID, IDUser)

	if err != nil {
		http.Error(w, "Error at removing a pick: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
