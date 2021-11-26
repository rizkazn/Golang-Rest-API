package middleware

import (
	"net/http"

	"github.com/rizkazn/gorestfull/helpers"
)

func Validate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		headerToken := r.Header.Get("Authtoken")

		// if !strings.Contains(headerToken, "Bearer") {
		// 	helpers.Response(w, "Please Login First", 401, false)
		// 	return
		// }

		// token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, roles := helpers.ValidateToken(headerToken)

		if !checkToken {
			helpers.Response(w, "Please Login Again to Access This Page", 400, false)
			return
		}

		if roles != "admin" {
			helpers.Response(w, "Unauthorized Access", 400, true)
			return
		}

		next.ServeHTTP(w, r)
	}
}
