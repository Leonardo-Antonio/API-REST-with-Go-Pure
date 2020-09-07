package main

import (
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api-rest-mysql/certificates/authorization"
	"github.com/Leonardo-Antonio/api-rest-mysql/handler"
	"github.com/Leonardo-Antonio/api-rest-mysql/storage"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados")
	}
	con := storage.NewMysql()
	AlumnStore := storage.NewAlumn(con.Pool())
	UserStore := storage.NewUser(con.Pool())
	mux := http.NewServeMux()

	handler.RouterAlumn(mux, AlumnStore)
	handler.RouterLogin(mux, UserStore)

	log.Println("Corriendo en http://localhost:8080/v1/alumns/get-by-id")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("ha ocurrido un error %v", err)
	}
}
