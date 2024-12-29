package main

import "fmt"

func main() {
	// Criando um canal de um tipo específico
	canal := make(chan string, 2)

	// Enviando um valor para o canal
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	// Criando uma variável que vai receber o valor desse canal
	mensagem := <-canal
	mensagem2 := <-canal

	// Exibindo em tela o valor da variável
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
