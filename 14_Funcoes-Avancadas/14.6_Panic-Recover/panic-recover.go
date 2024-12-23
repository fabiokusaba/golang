package main

import "fmt"

// Existe uma maneira dessa função resolver o problema do panic e para isso usamos o recover
// Uma coisa importante de se mencionar é que se você estiver usando o recover em uma função que não está entrando em pânico não tem
// problema ele não vai influenciar, ele só vai ignorar isso aqui, ou seja, ele vai chamar o recover e se a sua função não estiver
// entrando em pânico o resultado em r vai ser nil então ele vai ignorar e seguir a vida
// O panic não é a melhor maneira de você tratar erros dentro do seu programa e você pode acabar vendo ele em alguns casos específicos
// se você estiver usando algum pacote externo ou mesmo algum pacote nativo no Go que tenha uma condição que faça o programa entrar em
// pânico é importante você saber como lidar com ele, então você usar o recover pra conseguir fazer o seu programa continuar sendo
// executado se for possível
// Importante frisar que o panic não é um erro, ele não é do tipo error que nós vimos, ele é um pouco diferente, ele tem as
// características diferentes dele
// Quando usamos o recover em Go normalmente fazemos da seguinte maneira:
func recuperarExecucao() {
	// Usando um if init declaramos uma variável r, chamamos a função recover, armazenamos o resultado dela em r, e caso esse
	// resultado seja diferente de nil quer dizer que a função foi recuperada com sucesso
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

// Vamos supor que em nosso programa, por algum motivo, exista uma regra em que a média do aluno nunca pode ser 6, se a média for 6
// meu programa vai entrar em pânico e não vai saber o que fazer
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// panic -> essa função vai interromper o fluxo normal do seu programa e vai parar tudo, então quando você chama a função panic
	// ela vai parar de executar o que está executando e vai começar a entrar em pânico, basicamente ela vai chamar todas as funções
	// que tem defer, se você não recuperar a função usando o recover o seu programa morre, ou seja, ele para de executar
	// É um pouco diferente do erro que tínhamos visto, o erro você pode retorná-lo e seu programa continua executando normalmente,
	// você pode fazer o tratamento dele e seguir o programa, o panic não, o panic ele mata a execução do programa
	// Então, se você tiver uma cláusula panic e você não estiver tentando recuperar ela usando o recover o seu programa morre e nada
	// do que era para ser executado depois disso vai ser executado, essa é a diferença principal entre um erro normal e o panic
	// Existe uma maneira de você recuperar um programa que está entrando em pânico e é justamente usando a cláusula recover
	panic("A média é 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
