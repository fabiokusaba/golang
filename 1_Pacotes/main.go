package main

import (
	"fmt"
	"github.com/badoux/checkmail"

	"github.com/fabiokusaba/golang/1_Pacotes/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
