package app

import "github.com/urfave/cli"

// Esse pacote vai ter a nossa aplicação de linha de comando, tudo vai ser criado aqui
// Vamos ter uma função Gerar, perceba que aqui usamos o G maiúsculo porque a ideia é que o pacote main importe essa função, então
// dessa forma estamos tornando pública/visível essa função
func Gerar() *cli.App {
	// Esse método que chamamos vai retornar o que precisamos para a nossa função Gerar
	app := cli.NewApp()

	// Definindo algumas configurações da nossa aplicação
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	return app
}
