package main

import "fmt"

// Funções recursivas -> basicamente são funções que chamam elas mesmas, então nesse caso uma função recursiva é como se ela
// dependesse dela mesma para funcionar
// Sequência de Fibonacci -> sequência de números inteiros, começando normalmente por 0 e 1, na qual, cada termo subsequente
// corresponde à soma dos dois anteriores
// Uma coisa muito importante das funções recursivas é que elas tem uma condição de parada, então você tem que especificar
// para o seu programa o momento em que ele vai parar de chamar a função porque se não você pode ficar numa execução infinita
// e gerar o que chamamos de estouro de pilha
// Estouro de pilha -> cada execução que você vai fazendo dessa função ela vai colocando em uma pilha e se você tiver muitas
// execuções ele vai criar uma pilha muito grande, se você nunca parar de chamar essa função vai ser uma pilha gigantesca e
// vai ocorrer o que a gente chama de estouro de pilha, o seu sistema não vai conseguir suportar ela porque ele vai depender
// de muitas execuções, ele vai ficar esperando muitas execuções e vai dar o estouro de pilha, em inglês é chamado de stack
// overflow
// Uma coisa a se comentar é que as funções recursivas não são recomendadas se você tiver números muito grandes, muitas execuções
// sendo feitas porque você pode ter um estouro de pilha devido a quantidade de execuções que você está fazendo
func fibonacci(posicao uint) uint {
	// A primeira coisa que vamos especificar vai ser a nossa condição de parada
	if posicao <= 1 {
		// Se a posição for 1 ou 0, então vou retornar o próprio valor da posição
		return posicao
	}

	// Se não, vou continuar chamando a função fibonacci até que em algum momento a nossa condição de parada seja atingida
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// Convertendo a posição para um uint (unsigned int) porque a posição não pode ser um número negativo
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	// Imprimindo a sequência de Fibonacci
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
