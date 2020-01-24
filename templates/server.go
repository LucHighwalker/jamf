package templates

import "fmt"

var DefaultMiddleware = `this.server.use(expressSanitizer());
		this.server.use(bodyParser.json());
		this.server.use(bodyParser.urlencoded({ extended: true }));
		this.server.use(cors());`

func Server(middleware, routes string) []byte {
	tmpl := `import * as express from "express";
import * as mongoose from "mongoose";
import * as bodyParser from "body-parser";
import * as cors from "cors";
import * as expressSanitizer from 'express-sanitizer';

class Server {
	public server;

	constructor() {
		this.server = express();
		this.connectDb();
		this.applyMiddleware();
		this.mountRoutes();
	}

	private connectDb(): void {
		const mongo = process.env.MONGO_URI;
		mongoose.connect(mongo, {
			useNewUrlParser: true,
			useCreateIndex: true
		});
		const db = mongoose.connection;
		db.on("error", console.error.bind(console, "MongoDB Connection error"));
	}

	private applyMiddleware(): void {
		%s
	}

	private mountRoutes(): void {
		%s
	}
}

export default new Server().server;`

	return []byte(fmt.Sprintf(tmpl, middleware, routes))
}
