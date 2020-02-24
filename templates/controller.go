package templates

import (
	"fmt"
	"jamf/helper"
	"strings"
)

// Action - Template for a controller's action
func Action(access, name, args, ret string) string {
	tmpl := `	%s %s(%s) %s {
		return {
			message: "%s called.",
		}
	}
`

	return fmt.Sprintf(tmpl, access, name, args, ret, name)
}

// Controller - Main template for a controller
func Controller(model, name, actions string) []byte {
	var m string
	if model != "" {
		m = fmt.Sprintf("import %s from \"../models/%s\"", helper.Capitalize(model), model)
	}

	tmpl := `class %sController {%s}

export default new %sController();
`

	n := helper.Capitalize(name)

	return []byte(fmt.Sprintf(tmpl, m, n, actions, n))
}

// Route - Template for each route
func Route(name, verb, action string, params, args []string) string {
	ptmpl := "\n\t\t\t\tconst {%s} = req.params;\n"
	pp := ""

	tmpl := `this.router.%s('/%s%s', async (req, res) => {
			try {%s
				const %s = %s.%s(%s);
				res.status(200).json({
					%s
				});
			} catch (error) {
				res.status(500).json({
					error: error.message
				});
			}
		});`

	p := strings.Join(params, "")
	a := strings.Join(args, ", ")

	if len(args) > 0 {
		pp = fmt.Sprintf(ptmpl, a)
	}

	return fmt.Sprintf(tmpl, verb, action, p, pp, action, name, action, a, action)
}

// Router - Main router template.
func Router(name, routes string) []byte {
	tmpl := `import * as express from 'express';

import %s from './%s.controller';

class %s {
	public router: express.Router;

	constructor() {
		this.router = express.Router();

		this.router.route('/');

		%s
	}
}

export default new %s().router;
`

	n := helper.Capitalize(name)

	return []byte(fmt.Sprintf(tmpl, name, name, n, routes, n))
}
