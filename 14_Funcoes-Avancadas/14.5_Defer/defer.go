package main

import "fmt"

// Defer -> basicamente ela adia a execução de um determinado pedaço de código, pode ser uma função que você criou, uma função de
// algum pacote
// Dentre outras coisas o defer é muito útil quando estamos lidando com banco de dados, imagine que você tem uma função que vai
// manipular o banco de dados, então ela pode fazer uma consulta, inserir registros, atualizar, fazer diversas operações, e o que
// acontece? Quando você chama essa função normalmente você abre a conexão com o banco de dados e no meio de todas essas operações
// que você está fazendo pode ser que ocorra um erro e você tenha que dar um return antes do que você estava esperando, mas nesses
// casos independente do que aconteça você quer, por exemplo, fechar a conexão com o banco de dados, então pra você não ter que
// ficar colocando essa sentença para fechar a conexão em todos os lugares você poderia usar o defer
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	// Apesar desse cara estar na primeira linha da função ele será executado imediatamente antes do return indepedente de qual seja
	defer fmt.Println("Média calculada. Resultado será retornado!")

	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	// Se eu colocar a cláusula defer na frente da funcao1 ele primeiro vai executar a funcao2 e depois a funcao1, na prática isso
	// quer dizer que o defer em inglês quer dizer adiar, então essa cláusula está falando para o seu código para ele adiar a
	// execução dessa função até o último momento possível, no caso como estamos lidando com a função main, ela não tem retorno, o
	// último momento possível seria antes dela terminar
	// Se a gente estivesse lidando com funções que dão retorno o defer seria executado imediatamente antes do retorno ser dado
	// independente de qual seja
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
