package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Node on Rails"
	app.Usage = "Like Ruby on Rails but NodeJS"
	author := cli.Author{Name: "Luc Highwalker", Email: "email@luc.gg"}
	app.Authors = []*cli.Author{&author}
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize a new NoR project.",
			Action: func(c *cli.Context) error {
				// initialize
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add an existing module.",
			Action: func(c *cli.Context) error {
				// add module
				return nil
			},
		},
		{
			Name:    "controller",
			Aliases: []string{"c"},
			Usage:   "Create a new controller.",
			Action: func(c *cli.Context) error {
				// generate controller
				return nil
			},
		},
		{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "Generate a model.",
			Action: func(c *cli.Context) error {
				// Generate model
				return nil
			},
		},
		{
			Name:    "struct",
			Aliases: []string{"s"},
			Usage:   "Generate a structure (model/controller).",
			Action: func(c *cli.Context) error {
				// Generate a model struct
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
