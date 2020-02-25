package modeler

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/urfave/cli"

	"jamf/helper"
	"jamf/templates"
)

var tempPath string

// Command - Command for this module.
func Command(nd, tp, wd string) *cli.Command {
	return &cli.Command{
		Name:    "model",
		Aliases: []string{"m"},
		Usage:   "Create a model.",
		Action: func(c *cli.Context) error {
			tempPath = tp

			name := c.Args().First()
			args := c.Args().Tail()

			generateModel(wd, name, args)
			return nil
		},
	}
}

func generateModel(wd, name string, args []string) {
	fields, face := generateFields(args)

	model := templates.Model(name, fields)
	iface := templates.Interface(name, face)

	modelsPath := path.Join(tempPath, "models")
	interfacesPath := path.Join(tempPath, "interfaces")

	helper.EnsureDirExists(modelsPath)
	helper.EnsureDirExists(interfacesPath)

	ioutil.WriteFile(path.Join(modelsPath, fmt.Sprintf("%s.ts", name)), []byte(model), 0644)
	ioutil.WriteFile(path.Join(interfacesPath, fmt.Sprintf("%s.ts", name)), []byte(iface), 0644)

	helper.CopyDir(tempPath, path.Join(wd, "src"))
}

func generateFields(args []string) (string, string) {
	var fields string
	var face string
	for _, f := range args {
		split := strings.Split(f, ":")
		capType := helper.Capitalize(split[1])
		field := fmt.Sprintf("\t%s: {\n\t\ttype: %s,", split[0], capType)
		face = fmt.Sprintf("%s\n\t%s: %s;", face, split[0], capType)
		for i, s := range split {
			if i > 1 { // ignore name and type
				switch true {
				case s == "required":
					field = fmt.Sprintf("%s\n\t\trequired: true,", field)
					break

				case s == "unique":
					field = fmt.Sprintf("%s\n\t\tunique: true,", field)
					break

				case strings.Contains(s, "min"):
					ss := strings.Split(s, "=")
					field = fmt.Sprintf("%s\n\t\tmin: [%s, \"%s too small, %s minumum.\"],", field, ss[1], helper.Capitalize(split[0]), ss[1])
					break

				case strings.Contains(s, "max"):
					ss := strings.Split(s, "=")
					field = fmt.Sprintf("%s\n\t\tmax: [%s, \"%s too large, %s maximum.\"],", field, ss[1], helper.Capitalize(split[0]), ss[1])
					break

				case strings.Contains(s, "validate"):
					field = fmt.Sprintf("%s%s", field, generateValidation(split[0], s))
					break
				}
			}
		}
		field = fmt.Sprintf("%s\n\t},", field)
		fields = fmt.Sprintf("%s\n%s", fields, field)
	}
	return fmt.Sprintf("%s\n", fields), fmt.Sprintf("%s\n", face)
}

func generateValidation(f, v string) string {
	r, _ := regexp.Compile("(/[\\s\\S\\d]+/[a-z]*)")
	reg := r.Find([]byte(v))

	validate := fmt.Sprintf("\n\t\tvalidate: [\n\t\t\tfunction(%s: string) {", f)

	if reg != nil {
		validate = fmt.Sprintf("%s\n\t\t\t\tif (this.isModified(\"%s\")) return %s.test(%s);\n\t\t\t\treturn true;\n\t\t\t", validate, f, reg, f)
	}

	validate = fmt.Sprintf("%s},\n\t\t\t\"%s failed validation.\"\n\t\t]", validate, helper.Capitalize(f))

	return validate
}
