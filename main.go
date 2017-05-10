package main

import (
	"fmt"
	"os"
	"quake-log-parser/parser"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name: "quake-log-parser",
		Commands: []cli.Command{
			{
				Name:        "serve",
				Aliases:     []string{"s"},
				Usage:       "run the server",
				Description: "This start the server application",
				Action: func(c *cli.Context) (err error) {
					fmt.Printf("TODO")
					return
				},
			}, {
				Name:        "parse",
				Aliases:     []string{"p"},
				Usage:       "run the parser",
				Description: "This run a parser",
				Action: func(c *cli.Context) (err error) {
					fileStr := c.Args().Get(0)
					if fileStr == "" {
						fileStr = "games.log"
					}

					err = parser.Parse(fileStr)
					return
				},
			},
		},
	}

	app.Run(os.Args)
}
