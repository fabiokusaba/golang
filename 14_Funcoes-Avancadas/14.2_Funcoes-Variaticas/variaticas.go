package main

import "fmt"

// Função variática -> basicamente é uma função que pode receber n parâmetros, você não precisa especificar a quantidade de parâmetros
// que ela vai receber
// A limitação é que você não pode ter mais de um parâmetro variádico por função e também ele precisa obrigatoriamente ser o último
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo!", 1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
}
