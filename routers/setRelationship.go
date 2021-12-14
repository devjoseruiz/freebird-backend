package routers

import (
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*SetRelationship performs the creation of relationships between users*/
func SetRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) == 0 {
		http.Error(w, "ID is mandatory", http.StatusBadRequest)
		return
	}

	var t models.Relationship
	t.UserID = IDUser
	t.FollowingUserID = ID

	status, err := db.InsertRelationship(t)

	if err != nil || status == false {
		http.Error(w, "Error creating relationship"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
