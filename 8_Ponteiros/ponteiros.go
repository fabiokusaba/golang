package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	// Estou criando uma variável e atribuindo um valor a ela, e depois estou criando uma outra variável e atribuindo o
	// valor da primeira variável a ela.
	var variavel1 int = 10
	var variavel2 int = variavel1
	// Sem nenhuma surpresa temos duas variáveis com o mesmo valor
	fmt.Println(variavel1, variavel2)

	// Alterando o valor de uma das variáveis
	variavel1++
	// Percebemos aqui que mudou apenas o valor da primeira variável
	fmt.Println(variavel1, variavel2)

	// Declarando uma variável do tipo inteiro e uma variável do tipo ponteiro para inteiro
	var variavel3 int // Guarda um valor inteiro
	var ponteiro *int // Guarda o endereço de memória de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)  // Valor da variável e endereço de memória da variável
	fmt.Println(variavel3, *ponteiro) // Valor da variável e valor da variável que está no endereço de memória

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
}
