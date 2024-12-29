package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando canais
	canal1, canal2 := make(chan string), make(chan string)

	// Criando goroutines para enviar valores para os canais
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// Criando um loop infinito para receber os valores e exibir em tela
	for {
		// Solucionando o problema do tempo utilizando select, aqui estamos falando que: caso o canal1 esteja pronto para receber um
		// valor exiba esse valor em tela, caso o canal2 esteja pronto para receber exiba o valor em tela
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// Nesse trecho de código abaixo temos o problema do tempo, ou seja, estamos desperdiçando tempo porque precisamos aguardar que
		// ambos recebam valores para prosseguir com o loop mesmo que um já esteja pronto
		// mensagemCanal1 := <-canal1
		// fmt.Println(mensagemCanal1)

		// mensagemCanal2 := <-canal2
		// fmt.Println(mensagemCanal2)
	}
}
