package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Declarando um variável do tipo array de inteiros com 5 posições
	var array1 [5]int
	// Populando um array acessando as suas posições
	array1[0] = 1
	fmt.Println(array1)

	// Declarando um array com inferência de tipo e passando os valores
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Declarando o array desta forma não significa que ele se tornou dinâmico apenas quer dizer que o seu tamanho vai
	// ser definido de acordo com a quantidade de elementos passados
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Declarando um slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// Adicionando um elemento ao slice com a função append
	slice = append(slice, 18)
	fmt.Println(slice)

	// Criando uma fatia do array especificando o intervalo
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Alterando um valor do array e refletindo na fatia, pois a fatia é uma referência para o array
	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Demonstrando os tipos através do pacote reflect
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
}
