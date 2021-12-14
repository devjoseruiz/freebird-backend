package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"gitlab.com/joseruizdev/freebird-backend/db"
	"gitlab.com/joseruizdev/freebird-backend/models"
)

/*UploadAvatar allows the user to upload a profile image*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")

	var split = strings.Split(handler.Filename, ".")
	var extension = split[len(split)-1]

	var saveFile string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error uploading image: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error saving the image: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension

	status, err = db.ModifyProfile(user, IDUser)

	if err != nil || status == false {
		http.Error(w, "Error updating user profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
