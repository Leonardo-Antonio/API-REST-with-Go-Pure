package middlewares

import (
	"log"
	"net/http"
)

// Log .. date and peticion of the query
func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Petición: %q, Method: %q\n", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Authorization middleware que simula un jwt
func Authorization(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "token-jwt" {
			// respinse -> prohibido o acceso denegado
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
