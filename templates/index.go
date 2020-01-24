package templates

import "fmt"

func Index(defaultPort int) []byte {
	if defaultPort == -1 {
		defaultPort = 4200
	}

	tmpl := `import * as dotenv from 'dotenv';
dotenv.config();

import server from './src/server';

const port = process.env.PORT || %d;

server.listen(port, (error: Error) => {
	if (error) {
		return console.log(error);
	}

	return console.log(` + "`server is listening on ${port}`" + `);
});`

	return []byte(fmt.Sprintf(tmpl, defaultPort))
}
