package main

import "fmt"

// Nada mais é do que uma função que vai ser executada antes da função main, a ordem não influência
// A grande diferença da função init para as outras é que você pode ter uma por arquivo, não uma por pacote então se você tiver um
// pacote que tem dez arquivos dentro, cada um deles pode ter a sua função init e antes de cada um desses arquivos ser executado a
// função init deles vai ser executada
// Ela pode ser usada para você fazer algum tipo de setup, inicializar alguma variável que vai ser usada pela função, alguma variável
// global
var n int

func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
