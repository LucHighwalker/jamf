package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/urfave/cli"

	"nor/helper"
)

func ModelsCommand(nd, wd string) *cli.Command {
	return &cli.Command{
		Name:    "model",
		Aliases: []string{"m"},
		Usage:   "Create a model.",
		Action: func(c *cli.Context) error {
			args := c.Args().Tail()

			fmt.Println(generateFields(args))
			return nil
		},
	}
}

func generateFields(args []string) string {
	var fields string
	for _, f := range args {
		split := strings.Split(f, ":")
		field := fmt.Sprintf("\t%s: {\n\t\ttype: %s,", split[0], helper.Capitalize(split[1]))
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
					field = fmt.Sprintf("%s\n\t\tmin: [%s, \"%s too small, %s minumum.],", field, ss[1], helper.Capitalize(split[0]), ss[1])
					break

				case strings.Contains(s, "max"):
					ss := strings.Split(s, "=")
					field = fmt.Sprintf("%s\n\t\tmax: [%s, \"%s too large, %s maximum.],", field, ss[1], helper.Capitalize(split[0]), ss[1])
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
	return fields
}

func generateValidation(f, v string) string {
	r, _ := regexp.Compile("(/[\\s\\S\\d]+/[a-z]*)")
	reg := r.Find([]byte(v))

	validate := fmt.Sprintf("\n\t\tvalidate: [\n\t\t\tfunction(%s: string) {", f)

	if reg != nil {
		validate = fmt.Sprintf("%s\n\t\t\t\tif (this.isModified(\"%s\") return %s.test(%s);\n\t\t\t\treturn true;\n\t\t\t", validate, f, reg, f)
	}

	validate = fmt.Sprintf("%s},\n\t\t\t\"%s failed validation.\"\n\t\t]", validate, helper.Capitalize(f))

	return validate
}
