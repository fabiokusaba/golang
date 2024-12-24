package main

import (
	"fmt"
	"math"
)

// A interface só tem assinaturas de métodos, você não pode passar campos da mesma forma que a gente passa numa struct, você pode
// passar assinaturas de métodos que falam como esses métodos devem ser
// Interessante vermos a vantagem disso, tenho uma função que está recebendo uma interface que pode ser implementada por vários
// tipos diferentes, então como Go é muito rígido com a questão dos tipos usando a interface a gente consegue dar uma remediada
// nisso e deixar um pouco mais flexível
type forma interface {
	area() float64
}

// Agora que eu tenho uma interface do que eu criar duas funções para representar a área de cada uma das minhas figuras eu posso
// criar uma única função escreverArea que vai receber uma forma e eu vou escrever na tela o valor da área
// Perceba que eu tenho uma função que ao invés de receber uma estrutura concreta, como é o caso do struct, ela recebe uma estrutura
// abstrata que é o caso da interface
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// Imagine que dentro do nosso sistema vamos ter algumas formas geométricas e eu quero ter uma função que escreva na tela a área
// dessas figuras
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// Como é que eu consigo passar um retangulo como parâmetro dessa função e até mesmo um círculo? Eu preciso que essas minhas figuras
// tenham um método chamado area que não receba nenhum parâmetro e que retorne um float64
// A implementação de interfaces em Go é implícita, existe linguagens em que você tem que declarar a interface e depois declarar uma
// classe, no caso das linguagens orientadas a objetos, e falar que ela implementa essa interface manualmente
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
