package main

import "fmt"

// Quando passamos os parâmetros normalmente sem utilizar o '*', ponteiro, o que está acontecendo? Estou passando uma cópia do
// valor para essa função, o que isso quer dizer? Que a minha função não está mexendo na variável. Desta forma, não estou mexendo
// no endereço de memória desse cara, estou mexendo no valor dele e no caso fazendo uma cópia e deixando a variável inalterada
// Quando a gente chama uma função assim e passa uma variável normal pra ela a gente fala que estamos passando um parâmetro por
// cópia, estamos passando a cópia de um valor pra essa função
func inverterSinal(numero int) int {
	return numero * -1
}

// Quando a gente usa um ponteiro a gente está passando uma referência pra essa função, qualquer alteração feita nessa referência
// vai impactar a variável do lado de fora da função também
// Isso é muito importante que você entenda como funciona no Go porque existem diversas funções dentro de pacotes tanto padrões da
// linguagem como pacotes externos que precisam especificamente receber um ponteiro e outras que precisam especificamente receber
// uma variável normal, então é muito importante saber diferenciá-los e saber o impacto que passar de um jeito ou de outro pode
// trazer para o seu programa como um todo
// Em linhas gerais quando realmente você quer alterar a variável dentro de uma função e garantir que ela fique alterada nos outros
// lugares do seu código você usa o ponteiro, quando você não quer fazer essa alteração no seu código inteiro você só quer fazer
// uma cópia e manter o valor original por algum motivo você usa o parâmetro normal sem ser o ponteiro
func inverterSinalComPonteiro(numero *int) {
	// Aqui dentro vou alterar o valor que está armazenado nesse endereço de memória que o ponteiro está me passando
	// Relembrando que o ponteiro ele aponta para um endereço de memória
	// Então, nesse caso precisamos colocar o '*' antes e fazer o processo de desreferenciação, ou seja, o valor do número, vou
	// buscar no endereço de memória que está armazenado que valor está lá, e vou alterar esse valor
	// O interessante é que não vou precisar ter um retorno porque eu já vou estar fazendo a alteração diretamente na variável
	*numero = *numero * -1
}

func main() {
	// A primeira função inverterSinal ela não mexe no valor da variável propriamente dita, ela cria uma nova variável, faz uma
	// cópia desse valor e manda para a função, e depois da execução da função a primeira variável fica inalterada
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	// Quando a gente mexe com ponteiros a história é um pouco diferente porque eu criei aqui uma variável normal, quando eu passei
	// ela pra função, na verdade eu passei o endereço de memória onde essa variável está salva, e toda a alteração que eu fizer nesse
	// cara vai refletir diretamente nessa variável
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // '&' -> passando o endereço de memória onde essa variável está salva
	fmt.Println(novoNumero)
}
