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
