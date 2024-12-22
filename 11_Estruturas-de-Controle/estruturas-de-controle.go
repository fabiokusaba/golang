package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init -> estou declarando uma variável e atribuindo um valor a ela, e estou avaliando se o valor dessa variável que acabei
	// de criar atende a uma certa condição que no caso é ser maior que zero
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
}
