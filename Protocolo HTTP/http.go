package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP E UM PROTOCOLO DE COMUNICACAO - BASE DA COMUNICACAO WEB

	// CLIENTE (FAZ A REQUISICAO) - SERVIDOR (PROCESSA REQUISICAO E ENVIA  RESPOSTA)

	// REQUEST - RESPONSE

	// ROTAS
	// URI - IDENTIFICADOR DO RECURSO
	// METODO - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola Mundo!"))
	})

	http.HandleFunc("/estudos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Roadmap de estudos!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
