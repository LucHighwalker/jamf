package editor

import (
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

func AddMiddleware(wd, mw string) string {
	serverPath := path.Join(wd, "src", "server.ts")
	middleWare := getMiddleware(serverPath)
	middleWare = middleWare + "\n\t\tthis.server.use(" + mw + ");"
	return middleWare
}

func AddRoute(wd, r string) string {
	serverPath := path.Join(wd, "src", "server.ts")
	routes := getRoutes(serverPath)
	routes = routes + "\n\t\tthis.server.use(" + r + ");"
	return routes
}

func getContent(src string) []byte {
	content, err := ioutil.ReadFile(src)

	if err != nil {
		panic(err)
	}

	return content
}

func getMiddleware(src string) string {
	content := getContent(src)

	re, _ := regexp.Compile("(private applyMiddleware\\(\\): void {[\\s\\S\\d]+)(private mountRoutes)")
	match := re.Find(content)

	return isolateUse(string(match))
}

func getRoutes(src string) string {
	content := getContent(src)

	re, _ := regexp.Compile("(private mountRoutes\\(\\): void {[\\s\\S\\d]+)")
	match := re.Find(content)

	return isolateUse(string(match))
}

func isolateUse(str string) string {
	re := regexp.MustCompile(`(?m:(this.server.use[(a-zA-Z."',){ };:]+))`)
	matches := re.FindAllStringSubmatch(str, -1)
	result := []string{}
	for _, m := range matches {
		result = append(result, m[0])
	}
	return strings.Join(result, "\n\t\t")
}
