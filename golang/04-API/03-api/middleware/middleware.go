package middleware

import (
	"net/http"

	"github.com/2212025424/api/authorization"
)

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		_, err := authorization.ValidateToken(token)

		if err != nil {
			forbidden(w, r)
			return
		}

		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Acceso denegado"))
}
