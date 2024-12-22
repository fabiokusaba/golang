package main

import "fmt"

// Retorno nomeado -> isso quer dizer que a função vai retornar uma variável chamada soma e uma variável chamada subtração
// A ideia do retorno nomeado é justamente isso, você coloca só a cláusula return e o Go já sabe o que você está retornando
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
