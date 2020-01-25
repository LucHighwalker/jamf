package initializer

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/templates"
)

func InitCommand(nd, wd string) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a new NoR project.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "defPort",
				Value: 4200,
				Usage: "Default port for server to run on.",
			},
			&cli.IntFlag{
				Name:  "dbPort",
				Value: 27017,
				Usage: "Default port mongo runs on.",
			},
		},
		Action: func(c *cli.Context) error {
			initialize(nd, wd, c)
			return nil
		},
	}
}

func initialize(nd, wd string, c *cli.Context) {
	boilerPath := path.Join(nd, "boiler")
	tempPath := path.Join(nd, "__temp__")

	name := c.Args().First()
	if name == "" {
		name = "norApp"
	}

	config(boilerPath, tempPath)
	server(name, tempPath, c.Int("defPort"))
}

func config(bp, tp string) {
	configPath := path.Join(bp, "config")

	err := helper.CopyDir(configPath, tp)

	if err != nil {
		log.Fatal(err)
	}
}

func server(name, tp string, dp int) {
	index := templates.Index(dp)
	server := templates.Server(name, templates.DefaultMiddleware, "")

	srcPath := path.Join(tp, "src")
	helper.EnsureDirExists(srcPath)

	ioutil.WriteFile(path.Join(tp, "index.ts"), index, 0644)
	ioutil.WriteFile(path.Join(srcPath, "server.ts"), server, 0644)
}
