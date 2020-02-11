package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"

	"nor/controllator"
	"nor/helper"
	"nor/initializer"
	"nor/modeler"
	"nor/modulator"
)

var gopath string
var boilerPath string

var binDir, _ = os.Executable()

// var norDir = strings.TrimRight(binDir, "/nor") + "/nor"
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
	fmt.Println(binDir)

	nor.Commands = []*cli.Command{
		initializer.Command(binDir, boilerPath, tempPath, workDir),
		modulator.Command(binDir, boilerPath, tempPath, workDir),
		controllator.Command(workDir),
		modeler.Command(binDir, tempPath, workDir),
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
