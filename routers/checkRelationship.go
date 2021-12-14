package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*CheckRelationship checks if a relation between two users exists*/
func CheckRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relationship
	t.UserID = IDUser
	t.FollowingUserID = ID

	var response models.ResultCheckRelationship

	status, err := db.CheckRelationship(t)

	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
