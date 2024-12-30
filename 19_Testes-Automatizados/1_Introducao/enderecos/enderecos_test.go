package enderecos

import "testing"

// Vamos refatorar a nossa função para que ela lide com vários cenários de testes e para fazer isso vamos utilizar os structs,
// o struct vai representar um cenário de teste e ele vai conter dois campos: um vai ser o valor que a gente vai estar passando
// para a função TipoDeEndereco e o outro é o valor que eu espero receber dessa função
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Essa aqui é a sintaxe da assinatura de uma função de teste, então ela tem o seguinte formato: começa obrigatoriamente com Test
// seguido do nome da função que você vai testar, você não precisa fazer com o mesmo nome da função que você vai testar mas é uma
// boa prática fazê-lo
// Além disso, a função recebe um parâmetro muito específico que é um ponteiro do tipo T que está dentro do pacote testing, pacote
// padrão para você desenvolver testes automatizados
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}

func TestTipoDeEnderecoComDiversosCenarios(t *testing.T) {
	cenariosDeTestes := []cenarioDeTeste{
		{enderecoInserido: "Rua ABC", retornoEsperado: "Rua"},
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "Avenida"},
		{enderecoInserido: "Rodovia Julio Budisk", retornoEsperado: "Rodovia"},
		{enderecoInserido: "Estrada Velho Porto", retornoEsperado: "Estrada"},
		{enderecoInserido: "Praças das Rosas", retornoEsperado: "Tipo Inválido!"},
	}

	for _, cenario := range cenariosDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
