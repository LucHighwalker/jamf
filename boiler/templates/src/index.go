package templates

import "fmt"

func Index(defaultPort int) string {
	if defaultPort == -1 {
		defaultPort = 4200
	}

	tmpl := `import * as dotenv from 'dotenv';
	dotenv.config();
	
	import server from './server';
	
	const port = process.env.PORT || %s;
	
	server.listen(port, (error: Error) => {
		if (error) {
			return console.log(error);
		}
	
		return console.log(` + "`server is listening on ${port}`" + `);
	});`

	return fmt.Sprintf(tmpl, defaultPort)
}
