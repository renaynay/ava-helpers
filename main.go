package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/renaynay/ava-helpers/generate"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "AVA Account Generator"
	app.Commands = []cli.Command{
		generate.GenerateCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}