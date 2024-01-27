package app

import (
	"github.com/urfave/cli"
)

// Detect will return the server IP address searched
func Detect() *cli.App {
	app := *cli.NewApp()
	app.Name = "IP Detector"
	app.Usage = "IP Detector will return the server IP address searched"

	return &app
	
}