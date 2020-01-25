package initializer

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/templates"
)

func Initialize(nd, wd string, c *cli.Context) {
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
