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

func main() {
	info()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
