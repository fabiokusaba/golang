package main

import (
	"fmt"
	"time"
)

func main() {
	// A primeira forma que podemos usar o for é como um while
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(1 * time.Second)
	}

	// Uma outra maneira seria usando uma sintaxe muito parecida com if init, uma sintaxe mais padrão de outras linguagens para usar o for
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(1 * time.Second)
	}

	// For com a cláusula range, essa cláusula a gente utiliza quando vamos iterar sobre alguma coisa, por exemplo um array, slice, string
	// map, etc.
	// O range retorna dois valores, o primeiro é o índice e o segundo é o valor. Se você não quiser usar o índice você precisa usar o _
	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// É possível você iterar sobre uma string, só que aqui temos uma peculiaridade. Quando iteramos sobre uma string o valor que ele nos
	// devolve é o código da tabela ASCII, então para transformar em string precisamos usar o pacote string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// Também é possível iteramos sobre um map e aqui vale uma observação que o mesmo não se aplica a structs, então não é possível iterar
	// sobre uma struct
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Uma curiosidade, para você fazer um loop infinito basta fazer da seguinte forma:
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(1 * time.Second)
	}
}
