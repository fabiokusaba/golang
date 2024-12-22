package main

import "fmt"

// Funções anônimas -> basicamente é uma função que não tem nome e pra você chamar a execução dela é um pouco diferente
func main() {
	// Chamando uma função, mas não declarando ela explicitamente
	// Desta forma estou falando para o Go que isso aqui é uma função anônima, ela realiza essa execução, então declara
	// ela pra mim e assim que você terminar de declarar já executa
	// É possível passar parâmetros para ela e também que ela nos retorne valores
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá Mundo!")

	fmt.Println(retorno)
}
