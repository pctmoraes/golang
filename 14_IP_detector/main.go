package main

import (
	"ip-detector/app"
	"log"
	"os"
)

func main() {
	app := app.Detect()

	if _error := app.Run(os.Args); _error != nil {
		log.Fatal(_error)
	}
}