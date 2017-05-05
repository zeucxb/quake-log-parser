package main

import (
	"os"
	"quake-log-parser/actions"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name: "quake-log-parser",
		Commands: []cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "run the server",
				Description: "This start the server application",
				Action:      actions.Run,
			}, {
				Name:        "parse",
				Aliases:     []string{"p"},
				Usage:       "run the parser",
				Description: "This run a parser",
				Action:      actions.Parse,
			},
		},
	}

	app.Run(os.Args)
}
