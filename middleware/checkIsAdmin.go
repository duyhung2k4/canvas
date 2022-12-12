package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func AuthenticationAdmin(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			badRequest(w, r, err)
		} else {
			if claims["permission"] == "admin" {
				next.ServeHTTP(w, r)
			} else {
				notIsPermission(w, r, "Not is admin")
			}
		}
	})
}
