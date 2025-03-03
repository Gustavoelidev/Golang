package main

import (
	servidor "crud/Servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// CRUD - Create/Read/UPDATE/DELETE
	// Create - POST
	// READ - GET
	// UP - PUT
	// DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuario", servidor.CriarUsuario).Methods(http.MethodPost)

	println("Escutando na porta 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))

}
