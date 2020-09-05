package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api-rest-mysql/middlewares"
)

// RouterAlumn .
func RouterAlumn(mux *http.ServeMux, s Storage) {
	alumn := newAlumn(s)
	mux.HandleFunc("/v1/alumns/search", middlewares.Log(middlewares.Authorization(alumn.getByID)))
	mux.HandleFunc("/v1/alumns/get-all", middlewares.Log(alumn.getAll))
	mux.HandleFunc("/v1/alumns/update", middlewares.Log(alumn.update))
	mux.HandleFunc("/v1/alumns/create", alumn.create)
	mux.HandleFunc("/v1/alumns/delete", alumn.delete)
}
