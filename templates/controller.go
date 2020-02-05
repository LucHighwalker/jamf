package templates

import (
	"fmt"
	"nor/helper"
)

func Action(access, name, args, ret string) string {
	tmpl := `	%s %s(%s) %s {
		return {
			message: "%s called.",
		}
	}
`

	return fmt.Sprintf(tmpl, access, name, args, ret, name)
}

func Controller(name, actions string) []byte {
	tmpl := `class %sController {%s}

export default new %sController();
`

	n := helper.Capitalize(name)

	return []byte(fmt.Sprintf(tmpl, n, actions, n))
}
