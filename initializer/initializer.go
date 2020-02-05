package initializer

import (
	"io/ioutil"
	"path"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/templates"
)

func Command(nd, bp, tp, wd string) *cli.Command {
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
			initialize(nd, bp, tp, wd, c)
			return nil
		},
	}
}

func initialize(nd, bp, tp, wd string, c *cli.Context) {
	name := c.Args().First()
	if name == "" {
		name = "norApp"
	}

	config(bp, tp)
	server(name, tp, c.Int("defPort"))

	finalize(tp, wd, name)
}

func config(bp, tp string) {
	configPath := path.Join(bp, "config")

	err := helper.CopyDir(configPath, tp)

	if err != nil {
		panic(err)
	}
}

func server(name, tp string, dp int) {
	index := templates.Index(dp)
	server := templates.Server(templates.DefaultImports, name, templates.DefaultMiddleware, "")

	srcPath := path.Join(tp, "src")
	helper.EnsureDirExists(srcPath)

	ioutil.WriteFile(path.Join(tp, "index.ts"), index, 0644)
	ioutil.WriteFile(path.Join(srcPath, "server.ts"), server, 0644)
}

func finalize(tp, wd, name string) {
	projPath := path.Join(wd, name)
	helper.EnsureDirExists(projPath)

	helper.CopyDir(tp, projPath)
}
