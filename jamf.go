package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"

	"jamf/controllator"
	"jamf/helper"
	"jamf/initializer"
	"jamf/modeler"
	"jamf/modulator"
)

var gopath string
var boilerPath string

var binDir, _ = os.Executable()

var workDir, _ = os.Getwd()
var tempPath = path.Join("/tmp", "jamf")

var jamf = cli.NewApp()

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
	jamf.Name = "Just Another Mvc Framework"
	jamf.Usage = "It's just that."
	author := cli.Author{Name: "Luc Highwalker", Email: "email@luc.gg"}
	jamf.Authors = []*cli.Author{&author}
	jamf.Version = "0.0.1"
}

func commands() {
	fmt.Println(binDir)

	jamf.Commands = []*cli.Command{
		initializer.Command(binDir, boilerPath, tempPath, workDir),
		modulator.Command(binDir, boilerPath, tempPath, workDir),
		controllator.Command(workDir, tempPath),
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
	err := jamf.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
