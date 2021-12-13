package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*SavePick allows to save a pick in the DB*/
func SavePick(w http.ResponseWriter, r *http.Request) {
	var message models.Pick
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Invalid data: "+err.Error(), 400)
		return
	}

	pick := models.SavePick{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	var status bool
	_, status, err = db.PickInsert(pick)

	if err != nil {
		http.Error(w, "Error at attempting to save pick: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Pick couldn't be saved", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
