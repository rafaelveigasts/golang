package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação CLI"
	app.Usage = "Busca de Ip e nome de servidor"
	app.Version = "0.0.1"

	flags := []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca de endereços IPs na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca de nomes de servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context){
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}