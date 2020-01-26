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
			s := editor.AddImport(wd, "import someshit from \"something\";")
			fmt.Println(s)
			return nil
		},
	}
}
