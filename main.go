package main

import (
	"log"
	"os"

	"github.com/jonDufty/recipes/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "recipes"
	app.Usage = "Command line for running sub components"
	app.Commands = []*cli.Command{
		{
			Name:   "initdb",
			Usage:  "Initialise the database",
			Action: cmd.InitDB,
		},
		{
			Name:   "migrate",
			Usage:  "Run database migrations",
			Action: cmd.Migrate,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Print("Error running command: " + err.Error())
		os.Exit(1)
	}

}
