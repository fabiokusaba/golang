package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Esse pacote vai ter a nossa aplicação de linha de comando, tudo vai ser criado aqui
// Vamos ter uma função Gerar, perceba que aqui usamos o G maiúsculo porque a ideia é que o pacote main importe essa função, então
// dessa forma estamos tornando pública/visível essa função
func Gerar() *cli.App {
	// Esse método que chamamos vai retornar o que precisamos para a nossa função Gerar
	app := cli.NewApp()

	// Definindo algumas configurações da nossa aplicação
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	// Comandos da aplicação
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs de endereços na Internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos Servidores na Internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	// Obtendo o valor da flag 'host'
	host := c.String("host")

	// Busca por IPs utilizando o pacote net
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	// Como nossa variável ips é um slice eu posso estar iterando sobre ela
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	// Name server -> nome do servidor
	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
