package templates

import "fmt"

// Dockerfile - Main Dockerfile template
func Dockerfile(port int) []byte {
	tmpl := `FROM node:11
ARG ENVIRONMENT

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install
COPY . /usr/src/app

RUN npm install -g ts-node typescript nodemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait

RUN chmod +x /wait
RUN chmod +x ./node.sh

EXPOSE %d

CMD /wait && ./node.sh`

	return []byte(fmt.Sprintf(tmpl, port))
}

// DockerCompose - Main DockerCompose template.
func DockerCompose(name string, port int) []byte {
	tmpl := `version: "2"
services:
	%s:
		container_name: %s
		restart: always
		build: .
		ports:
			- "%s:%s"
		depends_on:
			- mongo
		links:
			- mongo
		environment:
			- ENVIRONMENT=prod
			- HOST_URL=domain.com
			- PORT=%s
			- MONGO_URI=mongodb://mongo:27017/%s
			- JWT_SECRET=averysecretsecret
			- MAILGUN_API_KEY=amailgunapi
			- EMAIL_DOMAIN=email.domain.com
			- WAIT_HOSTS=mongo:27017
		volumes:
			- "./src:/usr/src/app/src"
		logging:
			driver: "json-file"
			options:
				max-size: "1k"
				max-file: "3"
	mongo_%s:
		container_name: mongo_%s
		image: mongo
		volumes:
			- /data/db
		ports:
			- "27017:27017"
		logging:
			driver: "json-file"
			options:
				max-size: "1k"
				max-file: "3"
	portainer_%s:
		container_name: portainer_%s
		image: portainer/portainer
		command: -H unix:///var/run/docker.sock
		restart: always
		ports:
			- 9000:9000
		volumes:
			- /var/run/docker.sock:/var/run/docker.sock
			- portainer_data:/data
		logging:
			driver: "json-file"
			options:
				max-size: "1k"
				max-file: "3"

volumes:
	portainer_data:`

	return []byte(fmt.Sprintf(tmpl, name, name, port, port, port, name, name, name, name, name))
}

// DockerComposeOverride - Template for docker compose override.
func DockerComposeOverride(name string, port int) []byte {
	tmpl := `version: "2"
services:
	%s:
		environment: 
			- ENVIRONMENT=dev
			- HOST_URL=domain.com
			- PORT=%s
			- MONGO_URI=mongodb://mongo:27017/%s
			- JWT_SECRET=averysecretsecret
			- MAILGUN_API_KEY=amailgunapi
			- EMAIL_DOMAIN=email.domain.com`

	return []byte(fmt.Sprintf(tmpl, name, port, name))
}
