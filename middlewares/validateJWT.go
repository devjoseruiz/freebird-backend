package middlewares

import (
	"net/http"

	"gitlab.com/joseruizdev/freebird-backend/routers"
)

/*ValidateJWT validates the JWT provided in the request*/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error processing Token"+err.Error(),
				http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
