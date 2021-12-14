package routers

import (
	"io"
	"net/http"
	"os"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*GetBanner returns the user banner*/
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "ID is mandatory", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	fileOpen, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, fileOpen)

	if err != nil {
		http.Error(w, "Error returning the image", http.StatusBadRequest)
	}
}
