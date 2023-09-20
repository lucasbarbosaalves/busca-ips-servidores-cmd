package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando" // Nome da aplicação
	app.Usage = "Busca IPs e Nomes de Servidor na Internet" // Descricao da aplicação

	flags := []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "cv-lucasalves.vercel.app",
				},
			}

	// Definindo os comandos
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na Internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}


	func buscarIps(c *cli.Context)  {
		host := c.String("host") // Pegando o valor do host

		// Pegando função nativa do pacote net
		ips, erro := net.LookupIP(host) // LookupIP retorna um slice de IPs e um erro
		if erro != nil {
			// tratamento de erro
			log.Fatal(erro)
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

	func buscarServidores(c *cli.Context) {
		host := c.String("host") // Pegando o valor do host

		servidores, erro := net.LookupNS(host) // LookupNS retorna um slice de servidores e um erro
		if erro != nil {
			log.Fatal(erro)
		}

		for _, servidor := range servidores {
			fmt.Println(servidor.Host) // Pegando o nome do servidor
		}
	}
