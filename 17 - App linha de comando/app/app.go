package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// gerar vai retornar a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidores na internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs e Nomes de servidores na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
