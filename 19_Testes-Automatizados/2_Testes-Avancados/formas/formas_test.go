package formas

import (
	"math"
	"testing"
)

// Eu quero que essa função teste a área tanto do círculo quanto do retângulo, então para isso podemos usar uma maneira um pouco diferente
// para escrever o nosso teste, questão de organização
func TestArea(t *testing.T) {
	// Usado quando você quer separar as coisas, vamos ter duas funções que vão representar dois métodos que estão sendo testados
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			// Uma outra possibilidade que podemos usar aqui é o Fatalf, da mesma forma que o Errorf ele também vai falar que o seu teste
			// não passou só que ele vai parar aqui, então imagina que você tem várias condições sendo avaliadas o Errorf vai dar erro em
			// uma falando que o seu teste quebrou só que ele continua executando as outras, o Fatalf não, então assim que ele encontrar
			// que essa função foi chamada ele vai dar que o seu teste não passou e vai parar de executar ele
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
