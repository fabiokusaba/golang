package main

import "fmt"

// Isso aqui está falando para a minha função que ela pode receber um parâmetro que atenda a essa interface aqui, só que como
// não tem nada dentro dessa interface tudo atende, então uma string atenderia, um int atenderia, um booleano, qualquer coisa
// atenderia essa interface
// Então, isso é um jeito que você tem em Go pra usar tanto parâmetro quanto retorno como qualquer coisa que você quiser como
// sendo genérico e você não precisar ligar para o tipo
// Isso é bom? Depende, tem alguns casos, algumas funções que realmente fazem sentido você receber uma interface assim genérica
// para que você possa interagir com qualquer coisa, um exemplo claro disso é o Println
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
