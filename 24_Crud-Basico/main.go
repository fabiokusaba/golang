package main

import (
	"fmt"
	"log"
	"net/http"

	"crud/server"

	"github.com/gorilla/mux"
)

func main() {
	// Primeira coisa que fazemos é criar um router
	router := mux.NewRouter()

	// Criando funções para tratar requisições de rota com determinado método
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost) // Especificando o método com gorilla mux
	router.HandleFunc("/usuarios", server.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete)

	// Utilizar o pacote http para subir o servidor junto com o router
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
