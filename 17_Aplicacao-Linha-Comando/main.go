package main

import (
	"fmt"
	"log"
	"os"

	"linha-de-comando/app"
)

// A ideia vai ser que esse pacote main importar a aplicação de outro pacote, então vamos ter outro pacote que vai gerar a aplicação
// do nosso jeito com os comandos e ações que queremos que ela faça, vamos passar esse pacote aqui para o main que vai apenas executar
// essa aplicação
func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
