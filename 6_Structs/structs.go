package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	// Criando uma variável do tipo usuario
	var u usuario
	// Populando a variável com valores
	u.nome = "Davi"
	u.idade = 21
	// Exibindo o resultado
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// Declarando uma struct com inferência de tipo passando os valores na ordem em que se encontram
	u2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(u2)

	// Declarando uma struct com inferência de tipo, porém não passando todos os valores
	u3 := usuario{idade: 21}
	fmt.Println(u3)
}
