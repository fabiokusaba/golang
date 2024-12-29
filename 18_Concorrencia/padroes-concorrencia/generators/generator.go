package main

import (
	"fmt"
	"time"
)

func main() {
	// Um detalhe muito importante, essa nossa variável é agora um canal do tipo que só recebe
	canal := escrever("Olá Mundo!")

	// Imagine que eu queira imprimir dez mensagens na tela
	for i := 0; i < 10; i++ {
		// Imprimindo o valor que está chegando desse canal
		fmt.Println(<-canal)
	}
}

// Função escrever que recebe como parâmetro um texto do tipo string e retorna pra gente um canal do tipo string de uma direção só
func escrever(texto string) <-chan string {
	// Criando um canal do tipo string
	canal := make(chan string)

	// Aqui dentro vou ter uma goroutine que vai ficar num loop infinito que vai sempre jogar um valor para o canal
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Retornando o canal
	return canal
}
