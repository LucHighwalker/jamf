package editor

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"nor/helper"
)

func AddImports(wd, i string) string {
	server := path.Join(wd, "src", "server.ts")
	imports := getImports(server)
	if i != "" {
		imports = fmt.Sprintf("%s\n%s", imports, i)
	}
	return imports
}

func AddMiddleware(wd, mn, mw string) string {
	server := path.Join(wd, "src", "server.ts")
	middleWare := getMiddleware(server)
	if mw != "" {
		middleWare = fmt.Sprintf("%s\n\t\tthis.server.use(%sController.%s);", middleWare, mn, mw)
	}
	return middleWare
}

func AddRoute(wd, r, mn string) string {
	server := path.Join(wd, "src", "server.ts")
	routes := getRoutes(server)
	if r != "" {
		routes = fmt.Sprintf("%s\n\t\tthis.server.use(\"/%s\", %sRoutes)", routes, r, mn)
	}
	return routes
}

func getMiddleware(src string) string {
	content := helper.GetContent(src)

	re, _ := regexp.Compile("(private applyMiddleware\\(\\): void {[\\s\\S\\d]+)(private mountRoutes)")
	match := re.Find(content)

	return isolateUse(string(match))
}

func getRoutes(src string) string {
	content := helper.GetContent(src)

	re, _ := regexp.Compile("(private mountRoutes\\(\\): void {[\\s\\S\\d]+)")
	match := re.Find(content)

	return isolateUse(string(match))
}

func getImports(src string) string {
	content := helper.GetContent(src)

	re := regexp.MustCompile(`(?m:(import [* a-zA-Z"./;\-']+))`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	result := []string{}
	for _, m := range matches {
		result = append(result, m[0])
	}
	return strings.Join(result, "\n")
}

func isolateUse(str string) string {
	re := regexp.MustCompile(`(?m:(this.server.use[(a-zA-Z./"',){ };:]+))`)
	matches := re.FindAllStringSubmatch(str, -1)
	result := []string{}
	for _, m := range matches {
		result = append(result, m[0])
	}
	return strings.Join(result, "\n\t\t")
}
