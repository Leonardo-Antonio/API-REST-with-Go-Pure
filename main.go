package main

import (
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api-rest-mysql/handler"
	"github.com/Leonardo-Antonio/api-rest-mysql/storage"
)

func main() {
	con := storage.NewMysql()
	alumn := storage.NewAlumn(con.Pool())
	mux := http.NewServeMux()

	handler.RouterAlumn(mux, alumn)

	log.Println("Corriendo en http://localhost:8080/v1/alumns/get-by-id")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("ha ocurrido un error %v", err)
	}
}
