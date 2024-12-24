package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// O que isso aqui está fazendo? Você está criando como se fosse uma função que tem esse nome salvar só que a diferença é que
// ela está grudada a uma certa estrutura, no caso essa estrutura é um usuario
// Basicamente está falando que todos usuarios tem um método chamado salvar
// Esse u seria uma variável que a gente pode usar pra referenciar outros campos do mesmo usuario que chamou esse método
// Os métodos seguem a mesma lógica das funções você pode especificar retorno, você pode passar parâmetros
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Imagine que eu tenha um método que eu vou querer que ele altere um valor aqui dentro do usuario, por exemplo posso ter um método
// fazerAniversario que vai adicionar +1 na idade do nosso usuario, como eu faria isso? Utilizando ponteiro, mas nesse caso não vamos
// precisar fazer o processo de desreferenciação, a diferença crucial é que como usuario é um ponteiro quando eu fizer essa alteração
// aqui fora ele vai sofrer isso também
// Isso é um jeito bem conveniente de você usar os ponteiros quando você precisa de métodos, quando você precisa de uma struct que você
// precisa alterar os dados dela é bem comum você usar o ponteiro para isso
// Via de regra quando você tem um método que vai fazer alguma alteração no seu struct você costuma passar a referência do struct usando
// o '*', e quando você não vai fazer nenhuma alteração você não precisa passar como ponteiro você pode passar diretamente como uma cópia
// que não vai ter problema
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
