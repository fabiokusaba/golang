package main

import (
	"log"
	"net/http"
)

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	// Criando uma rota -> primeiro parâmetro é a URI, e o segundo é uma função que vai receber esse cara e vai saber lidar com ela
	// Para ser uma função que vai ser reconhecida aqui ela precisa ter uma assinatura específica e receber dois parâmetros: o primeiro
	// parâmetro é do tipo http.ResponseWriter e o segundo parâmetro é um ponteiro para *http.Request
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Escrevendo uma mensagem -> obrigatoriamente essa mensagem precisa estar no formato slice de bytes
		w.Write([]byte("Olá Mundo!"))
	})

	// Uma forma de organizar melhor o código seria criar uma função que vai lidar com a requisição e a resposta e passar essa função aqui
	// no método HandleFunc
	http.HandleFunc("/usuarios", usuarios)

	// Criando um servidor HTTP em Go
	log.Fatal(http.ListenAndServe(":5000", nil))
}
