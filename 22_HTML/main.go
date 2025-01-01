package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	// Basicamente nessa linha estou jogando na nossa variável templates todos os arquivos html que vamos criar
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Criando um usuário
		u := usuario{"Fulano", "fulano@email.com"}

		// Lendo arquivo html e renderizando na tela -> executar um template específico que está dentro dessa variável templates
		// Esse cara aqui recebe três parâmetros: o primeiro parâmetro é o ResponseWriter que está vindo da função, o segundo parâmetro é o nome
		// do template e o terceiro parâmetro por enquanto vamos passar como nil mas aqui seria se a gente tivesse algum dado para passar para
		// esse template
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
