package enderecos

import "strings"

// A ideia da nossa função é o seguinte: você vai passar um endereço, por exemplo Rua, Avenida, etc, e ela vai te retornar qual é
// o tipo desse endereço
// Avenida Paulista -> Avenida / Rua Tal -> Rua / X x -> Tipo Indefinido
// Como essa função vai ser exportada precisamos lembrar que ela precisa iniciar com a letra maiúscula
func TipoDeEndereco(endereco string) string {
	// Criando um slice com os tipos de endereços considerados válidos
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	// Pegando a primeira palavra desse endereço, basicamente aqui estamos dividindo a nossa string em um slice de acordo com o
	// separador que passamos e pegando a palavra que está contida na primeira posição
	// Rua dos Bobos -> ["Rua", "dos", "Bobos"]
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	// Verificando se a palavra está contida no slice de tipos válidos
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	// Verificando que se o endereço tem um tipo válido, se sim vou retornar a primeira palavra do endereço porque ele é o próprio
	// tipo
	if enderecoTemUmTipoValido {
		// Essa função Title do pacote strings vai pegar a primeira letra de cada palavra e vai deixar em maiúsculo
		return strings.Title(primeiraPalavraDoEndereco)
	}

	// Caso contrário, retorno que o tipo é inválido
	return "Tipo Inválido!"
}
