package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api-rest-mysql/middlewares"
)

// RouterAlumn .
func RouterAlumn(mux *http.ServeMux, s AlumnStorage) {
	alumn := newAlumn(s)
	mux.HandleFunc("/v1/alumns/search", middlewares.Log(middlewares.Authorization(alumn.getByID)))
	mux.HandleFunc("/v1/alumns/get-all", middlewares.Log(alumn.getAll))
	mux.HandleFunc("/v1/alumns/update", middlewares.Log(middlewares.Authorization(alumn.update)))
	mux.HandleFunc("/v1/alumns/create", middlewares.Log(middlewares.Authorization(alumn.create)))
	mux.HandleFunc("/v1/alumns/delete", middlewares.Log(middlewares.Authorization(alumn.delete)))
}

// RouterLogin .
func RouterLogin(mux *http.ServeMux, s UserStorage) {
	user := newUser(s)
	mux.HandleFunc("/v1/user", middlewares.Log(user.login))
	mux.HandleFunc("/v1/user/create", middlewares.Log(user.create))
	mux.HandleFunc("/v1/user/update", middlewares.Log(middlewares.Authorization(user.update)))
}
