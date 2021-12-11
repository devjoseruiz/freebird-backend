package middlewares

import (
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/db"
)

/*CheckDB is a middleware that checks DB connection status*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConn() == false {
			http.Error(w, "DB connection lost", 500)
		}

		next.ServeHTTP(w, r)
	}
}
