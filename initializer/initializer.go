package initializer

import (
	"io/ioutil"
	"log"
	"path"

	"nor/helper"
	"nor/templates"
)

func Initialize(nd, wd string) {
	boilerPath := path.Join(nd, "boiler")
	tempPath := path.Join(nd, "__temp__")

	config(boilerPath, tempPath)
	server(tempPath)
}

func config(bp, tp string) {
	configPath := path.Join(bp, "config")

	err := helper.CopyDir(configPath, tp)

	if err != nil {
		log.Fatal(err)
	}
}

func server(tp string) {
	index := templates.Index(4200)
	server := templates.Server(templates.DefaultMiddleware, "")

	srcPath := path.Join(tp, "src")
	helper.EnsureDirExists(srcPath)

	ioutil.WriteFile(path.Join(tp, "index.ts"), index, 0644)
	ioutil.WriteFile(path.Join(srcPath, "server.ts"), server, 0644)
}
