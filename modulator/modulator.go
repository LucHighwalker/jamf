package modulator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/urfave/cli"

	"nor/editor"
	"nor/helper"
	"nor/templates"
)

var tempPath string

type moduleConfig struct {
	Imports      []string
	Routes       string
	Middleware   string
	Dependencies []string
}

func Command(nd, bp, tp, wd string) *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add a module.",
		Action: func(c *cli.Context) error {
			tempPath = tp

			addModule(nd, bp, wd, c.Args().First())
			return nil
		},
	}
}

func addModule(nd, bp, wd, module string) {
	moduleDest := path.Join(wd, "src", module)
	if helper.DoesDirExist(moduleDest) {
		fmt.Printf("Module [%s] is already installed.\nAborting...\n", module)
		return
	}

	modulePath := path.Join(bp, "modules", module)

	name := path.Base(wd)

	if !helper.DoesDirExist(modulePath) {
		fmt.Printf("Module (%s) does not exist.\n", module)
		return
	}

	var config moduleConfig
	jsonContent := helper.GetContent(path.Join(modulePath, ".json"))
	json.Unmarshal(jsonContent, &config)

	helper.CopyDir(modulePath, tempPath)

	imports := generateImports(config.Imports, module)
	generateServer(tempPath, wd, module, imports, config.Middleware, config.Routes, name)

	helper.CopyDir(tempPath, path.Join(wd, "src"))

	fmt.Printf("Added module [%s]\n", module)
	for _, d := range config.Dependencies {
		fmt.Printf("[%s] reuires [%s]\n", module, d)
		addModule(nd, bp, wd, d)
	}
}

func generateImports(im []string, mn string) string {
	var imports string

	for _, i := range im {
		imp := mn + helper.Capitalize(i)
		imports = fmt.Sprintf("%s\nimport %s from \"./%s/%s.%s\";", imports, imp, mn, mn, i)
	}

	return imports
}

func generateServer(tp, wd, mn, i, mw, r, name string) {
	imports := editor.AddImports(wd, i)
	middleware := editor.AddMiddleware(wd, mn, mw)
	routes := editor.AddRoute(wd, r, mn)

	server := templates.Server(imports, name, middleware, routes)

	ioutil.WriteFile(path.Join(tp, "server.ts"), server, 0644)
}
