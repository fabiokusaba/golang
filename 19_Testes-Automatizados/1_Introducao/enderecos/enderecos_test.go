package enderecos

import "testing"

// Essa aqui é a sintaxe da assinatura de uma função de teste, então ela tem o seguinte formato: começa obrigatoriamente com Test
// seguido do nome da função que você vai testar, você não precisa fazer com o mesmo nome da função que você vai testar mas é uma
// boa prática fazê-lo
// Além disso, a função recebe um parâmetro muito específico que é um ponteiro do tipo T que está dentro do pacote testing, pacote
// padrão para você desenvolver testes automatizados
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua ABC"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
