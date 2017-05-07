package main

import (
	"fmt"
	"os"
	"quake-log-parser/parser"

	cli "gopkg.in/urfave/cli.v2"
)

type game struct {
	TotalKills int            `json:"total_kills"`
	Players    map[int]string `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func main() {
	app := &cli.App{
		Name: "quake-log-parser",
		Commands: []cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
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
					err = parser.Parse()
					return
				},
			},
		},
	}

	app.Run(os.Args)
}
