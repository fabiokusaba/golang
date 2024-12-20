package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // Declarando o tipo explicitamente
	variavel2 := "Variável 2"           // Declarando o tipo implicitamente

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( // Declaração múltipla de variáveis
		variavel3 string = "lalala"
		variavel4 string = "lelele"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6" // Declaração múltipla de variáveis com inferência de tipo
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // Trocando valores de variáveis
	fmt.Println(variavel5, variavel6)
}
