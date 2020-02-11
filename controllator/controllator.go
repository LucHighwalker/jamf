package controllator

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/templates"
)

func Command(wd string) *cli.Command {
	return &cli.Command{
		Name:    "controller",
		Aliases: []string{"c"},
		Usage:   "Initialize a new controller.",
		Action: func(c *cli.Context) error {
			generateController(wd, c)
			return nil
		},
	}
}

func generateController(wd string, c *cli.Context) {
	name := c.Args().First()
	actions, hasActions := generateActions(c.Args().Tail())

	if hasActions {
		actions = fmt.Sprintf("\n%s", actions)
	}

	controller := templates.Controller(name, actions)
	controllerPath := path.Join(wd, "src", name)

	helper.EnsureDirExists(controllerPath)
	ioutil.WriteFile(path.Join(controllerPath, fmt.Sprintf("%s.controller.ts", name)), controller, 0644)
}

func generateActions(actions []string) (string, bool) {
	result := []string{}
	for _, act := range actions {
		split := strings.Split(act, ":")
		name := split[0]
		args := []string{}
		access := "public"

		for i, a := range split {
			if i > 0 {
				switch true {

				case a == "private":
					access = a
					break

					// case strings.Contains(a, "ret"):
					// 	s = strings.Split(a, "=")
					// 	break

				case strings.Contains(a, "="):
					s := strings.Split(a, "=")
					args = append(args, fmt.Sprintf("%s: %s", s[0], helper.Capitalize(s[1])))
					break

				}
			}
		}

		result = append(result, templates.Action(access, name, strings.Join(args, ", "), ""))
	}
	return strings.Join(result, "\n"), len(result) > 0
}
