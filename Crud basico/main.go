package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - Create/Read/UPDATE/DELETE

	router := mux.NewRouter()
	println("Escutando na porta 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))

}
