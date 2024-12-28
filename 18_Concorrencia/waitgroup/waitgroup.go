package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Criando WaitGroup
	var waitGroup sync.WaitGroup

	// Adicionando goroutines que ele precisa esperar terminar
	waitGroup.Add(2)

	// Criando goroutines
	go func() {
		// Chamando a função escrever
		escrever("Olá Mundo!")
		// Quando a função terminar ela chama o método Done, basicamente esse método vai tirar 1 do contador
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	// Justamente falando para o nosso programa main esperar a contagem das goroutines chegar em zero
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
