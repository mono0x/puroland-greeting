package main

import (
	"log"
	"os"

	"github.com/mono0x/puroland-greeting/command"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "puroland-greeting"
	app.Commands = []cli.Command{
		command.NewServeCommand(),
		command.NewUpdateCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
