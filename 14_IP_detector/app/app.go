package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Detect will return the server IP address searched
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
		
	}

	return &app
}

func getIP(c *cli.Context) {
	host := c.String("host")

	ips, _error := net.LookupIP(host)

	if _error != nil {
		log.Fatal(_error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}