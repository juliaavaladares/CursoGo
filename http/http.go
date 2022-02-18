package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP -> protocolo de comunicacao - base da comunicaçao WEB
	// Cliente (faz requisiçao/request)-> Servidor (processa requisição e envia resposta/response)

	//Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT e DELET

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oi oi pessoas!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando pagina de usuarios"))
	})

	log.Fatal(
		http.ListenAndServe(":5000", nil),
	)

}
