package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um canal especificando o seu tipo
	canal := make(chan string)

	// Chamando a função como uma goroutine passando um texto e o canal que criamos
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// Recebendo um valor do canal, então estou esperando que o canal receba um valor
	// mensagem := <-canal
	// Exibindo em tela o valor recebido pelo canal
	// fmt.Println(mensagem)

	// Utilizando um for infinito para exibir os valores, então o que ele vai fazer é esperar sempre um valor chegar nesse canal para
	// exibir em tela
	// for {
	// 	// Aqui além de pegar dados do canal posso obter um outro retorno para verificar se o canal está aberto
	// 	mensagem, aberto := <-canal
	// 	// Estou verificando que se o canal não estiver aberto eu vou sair desse loop
	// 	if !aberto {
	// 		break // Comando para sair de um loop
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Refatorando o código acima, isso aqui está fazendo o seguinte: para cada mensagem recebida no canal, enquanto ele estiver aberto,
	// vou exibir o seu valor em tela
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

// Adicionamos como parâmetro da nossa função que além dela receber o texto também irá receber um canal do tipo string
func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		// Ao invés de escrever o texto em tela vou enviar esse dado pelo canal
		canal <- texto // Aqui estamos dizendo que o canal está recebendo o texto
		time.Sleep(time.Second)
	}

	// Fechando o canal, então o que estou falando aqui é que depois que terminar esse loop de cinco vezes eu vou fechar o canal, então
	// ele não vai mais enviar dados e nem receber dados
	close(canal)
}
