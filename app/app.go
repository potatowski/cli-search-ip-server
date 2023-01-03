package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return a cli ready to executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI"
	app.Usage = "Search IPs and Servername in Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs in Internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "sn",
			Usage:  "Search servername in Internet",
			Flags:  flags,
			Action: searchServer,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range servers {
		fmt.Println(s.Host)
	}
}
