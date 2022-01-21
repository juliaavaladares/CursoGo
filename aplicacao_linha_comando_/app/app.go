package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Funcao gerar retorna a aplicacao de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços na internet",
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

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
