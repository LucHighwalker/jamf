package modulator

import (
	"fmt"

	"github.com/urfave/cli"

	"nor/editor"
)

func Command(wd string) *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add an existing module.",
		Action: func(c *cli.Context) error {
			s := editor.AddMiddleware(wd, "HolyShit.ItWorks()")
			fmt.Println(s)
			// editor.GetMiddleware(path.Join(norDir, "testingGround", "tryingmiddw", "src", "server.ts"))
			return nil
		},
	}
}
