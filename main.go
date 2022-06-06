package main

import (
	"log"
	"os"

	"github.com/jonDufty/recipes/auth"
	"github.com/jonDufty/recipes/cmd"
	"github.com/jonDufty/recipes/cookbook"
	"github.com/jonDufty/recipes/graph"
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
			Name:  "migrate",
			Usage: "Run database migrations",
			Subcommands: []*cli.Command{
				{
					Name:   "up",
					Usage:  "Migrate up",
					Action: cmd.MigrateUp,
				},
				{
					Name:   "down",
					Usage:  "Migrate down",
					Action: cmd.MigrateDown,
				},
			},
		},
		{
			Name:  "serve",
			Usage: "Start app server",
			Subcommands: []*cli.Command{
				{
					Name:   "auth",
					Usage:  "Start auth microservice",
					Action: auth.ServeHandler,
				},
				{
					Name:   "cookbook",
					Usage:  "Start cookbook microservice",
					Action: cookbook.ServeHandler,
				},
				{
					Name:   "graph",
					Usage:  "Start Graphql server",
					Action: graph.ServeGraph,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Print("Error running command: " + err.Error())
		os.Exit(1)
	}

}
