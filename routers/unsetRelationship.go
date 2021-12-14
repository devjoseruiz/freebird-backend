package routers

import (
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*UnsetRelationship remove a follower relation in the DB*/
func UnsetRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relationship
	t.UserID = IDUser
	t.FollowingUserID = ID

	status, err := db.RemoveRelationship(t)

	if err != nil || status == false {
		http.Error(w, "Error removing relationship"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
