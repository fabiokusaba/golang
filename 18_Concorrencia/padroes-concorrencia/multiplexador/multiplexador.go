package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// Aqui estou passando dois canais porque o retorno da função escrever é um canal e a função multiplexar precisa receber dois
	// canais para funcionar
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	// Lendo do canal através do for
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// A ideia do multiplexador é: vou chamar essa função escrever mais de uma vez então vou ter mais de um canal, e a minha
// ideia é juntar esses dois canais em um único canal pra centralizar a comunicação num lugar só aqui dentro do main para
// isso vou fazer uma função multiplexar, essa função vai receber dois canais como entrada e vai retornar um canal só como
// saída
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	// Criando o canal de saída
	canalDeSaida := make(chan string)

	// Para fazer esse processo é bem simples, vou ter uma chamada para uma goroutine anônima
	go func() {
		// Aqui dentro vamos ter um for infinito
		for {
			// Aqui dentro vamos ter um select que vai ver qual dos dois canais tem uma mensagem disponível pra ser lida,
			// então se for do canal 1 ele vai jogar para o canal de saída, se for do canal 2 ele também vai jogar para o
			// canal de saída, então ele vai fazer uma espécie de encaminhamento de mensagens
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	// Ao final retornamos o nosso canal de saída, então estou centralizando a comunicação de dois canais dentro de um canal só
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(2000)))
		}
	}()

	return canal
}
