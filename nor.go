package main

import (
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/initializer"
	"nor/models"
	"nor/modulator"
)

var gopath string
var boilerPath string

var binDir, _ = os.Executable()
var norDir = strings.TrimRight(binDir, "/nor") + "/nor"
var workDir, _ = os.Getwd()
var tempPath = path.Join("/tmp", "nor")

var nor = cli.NewApp()

func clearTemp() {
	os.RemoveAll(tempPath)
	helper.EnsureDirExists(tempPath)
}

func getGopath() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	boilerPath = path.Join(gopath, "bin", "norb")
}

func checkBoiler() {
	exists := helper.DoesDirExist(boilerPath)
	if !exists {
		getBoiler()
	}
}

func getBoiler() {
	cmd := exec.Command("git", "clone", "https://github.com/LucHighwalker/norb.git", boilerPath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func info() {
	nor.Name = "Node on Rails"
	nor.Usage = "Like Ruby on Rails but NodeJS"
	author := cli.Author{Name: "Luc Highwalker", Email: "email@luc.gg"}
	nor.Authors = []*cli.Author{&author}
	nor.Version = "0.0.1"
}

func commands() {
	nor.Commands = []*cli.Command{
		initializer.InitCommand(norDir, boilerPath, workDir),
		modulator.Command(norDir, boilerPath, tempPath, workDir),
		{
			Name:    "controller",
			Aliases: []string{"c"},
			Usage:   "Create a new controller.",
			Action: func(c *cli.Context) error {
				// generate controller
				return nil
			},
		},
		models.ModelsCommand(norDir, tempPath, workDir),
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
	getGopath()
	checkBoiler()
	clearTemp()
	info()
	commands()
	err := nor.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
