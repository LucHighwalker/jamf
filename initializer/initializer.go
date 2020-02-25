package initializer

import (
	"io/ioutil"
	"path"

	"github.com/urfave/cli"

	"jamf/helper"
	"jamf/templates"
)

// Command - Command for this module.
func Command(nd, bp, tp, wd string) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a new jamf project.",
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
		name = "jamfApp"
	}

	config(bp, tp)
	dockerize(name, tp, 4200)
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

	helper.EnsureDirExists(path.Join(tp, "src"))
	ioutil.WriteFile(path.Join(tp, "index.ts"), index, 0644)
	ioutil.WriteFile(path.Join(tp, "src", "server.ts"), server, 0644)
}

func dockerize(name, tp string, dp int) {
	docker := templates.Dockerfile(dp)
	dockerComp := templates.DockerCompose(name, dp)
	dockerCompO := templates.DockerComposeOverride(name, 4200)

	ioutil.WriteFile(path.Join(tp, "Dockerfile"), docker, 0644)
	ioutil.WriteFile(path.Join(tp, "docker-compose.yml"), dockerComp, 0644)
	ioutil.WriteFile(path.Join(tp, "docker-compose.override.yml.example"), dockerCompO, 0644)
}

func finalize(tp, wd, name string) {
	projPath := path.Join(wd, name)
	helper.EnsureDirExists(projPath)

	helper.CopyDir(tp, projPath)
}
