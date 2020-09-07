package middlewares

import (
	"net/http"

	"github.com/Leonardo-Antonio/api-rest-mysql/certificates/authorization"
)

// Authorization middleware que simula un jwt
func Authorization(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbbiden(w, r)
			return
		}
		f(w, r)
	}
}

func forbbiden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"message-type": "Error", "message": "No autorizado"}`))
}
