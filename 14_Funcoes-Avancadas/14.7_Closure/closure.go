package main

import "fmt"

// Closure -> basicamente são funções que referenciam variáveis que estão fora do seu corpo
// Função que não recebe nenhum parâmetro e que retorna uma função
func closure() func() {
	// Declarando uma variável e atribuindo valor usando a inferência de tipo
	texto := "Dentro da função closure"

	// Função que não recebe nenhum parâmetro e que também não retorna nada, apenas printa na tela o valor que está na variável
	// texto
	funcao := func() {
		fmt.Println(texto)
	}

	// Retornando a outra função
	return funcao
}

// Agora vamos chamar a função closure dentro da função main
func main() {
	// Declarando uma variável texto
	texto := "Dentro da função main"
	// Exibindo o valor de texto
	fmt.Println(texto)

	// Criando uma variável que vai receber o retorno da minha função closure, como o retorno da função closure é uma outra função
	// essa variável aqui vai ser do tipo função
	funcaoNova := closure()
	// Então, vou chamar essa minha funcaoNova
	funcaoNova()

	// Quando rodamos a nossa aplicação percebemos que ele imprimiu o texto que está dentro da função closure, isso aqui é uma coisa
	// diferente da função closure, porque quando eu criei a função eu estava referenciando essa variável que estava aqui dentro,
	// então quando eu executar ela essa referência não vai ser perdida mesmo que dentro da função main, que está chamando essa função,
	// eu tenha também uma outra variável chamada texto, basicamente o Go vai manter a referência inicial que a gente tinha passado
	// Basicamente a função closure é isso, é como se ela tivesse uma "memória", ela tem uma lembrança da onde ela veio então ela sabe
	// que quando você der um Println no texto você vai estar sempre se referindo a esse texto que está dentro da função closure que é
	// o texto que foi passado para ela quando ela foi declarada
}
