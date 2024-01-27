package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Detect will return the server IP address 
// or the server name that is being searched
func Detect() *cli.App {
	app := *cli.NewApp()
	app.Name = "IP Detector"
	app.Usage = "IP Detector will return the server IP address searched"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Return IP addresses",
			Flags:  flags,
			Action: getIP,
		},
		{
			Name:	"server",
			Usage:	"Return the server name",
			Flags:	flags,
			Action:	getServers,
		},
	}

	return &app
}

func getIPs(c *cli.Context) {
	host := c.String("host")

	ips, _error := net.LookupIP(host)

	if _error != nil {
		log.Fatal(_error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

funct getServers(c *cli.Context) {
	host := c.String("host")

	servers, _error := net.LookupNS(host)

	if _error != nil {
		log.Fatal(_error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}