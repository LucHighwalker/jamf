package initializer

import (
	"go/build"
	"os"
	"path"
	"testing"
)

func boilerPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return path.Join(gopath, "bin", "jamfb")
}

func doFilesExist(files []string) (string, bool) {
	for _, dir := range files {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return dir, false
		}
	}
	return "", true
}

func values() (string, string, string) {
	os.RemoveAll("/tmp/jamf_test")
	os.RemoveAll("/tmp/jamf_test_result")

	bp := boilerPath()
	tp := "/tmp/jamf_test/initializer"
	wd := "/tmp/jamf_test_result/initializer"

	os.MkdirAll(tp, 0755)
	os.MkdirAll(wd, 0755)

	return bp, tp, wd
}

func TestInitialize(t *testing.T) {
	bp, tp, wd := values()

	initialize("testing", bp, tp, wd, 4200)

	missing, exists := doFilesExist([]string{
		"/tmp/jamf_test_result/initializer/testing/src",
		"/tmp/jamf_test_result/initializer/testing/src/server.ts",
		"/tmp/jamf_test_result/initializer/testing/.dockerignore",
		"/tmp/jamf_test_result/initializer/testing/.env.example",
		"/tmp/jamf_test_result/initializer/testing/.gitignore",
		"/tmp/jamf_test_result/initializer/testing/docker-compose.override.yml.example",
		"/tmp/jamf_test_result/initializer/testing/docker-compose.yml",
		"/tmp/jamf_test_result/initializer/testing/Dockerfile",
		"/tmp/jamf_test_result/initializer/testing/index.ts",
		"/tmp/jamf_test_result/initializer/testing/node.sh",
		"/tmp/jamf_test_result/initializer/testing/nodemon.json",
		"/tmp/jamf_test_result/initializer/testing/package-lock.json",
		"/tmp/jamf_test_result/initializer/testing/package.json",
		"/tmp/jamf_test_result/initializer/testing/README.md",
		"/tmp/jamf_test_result/initializer/testing/tsconfig.json",
		"/tmp/jamf_test_result/initializer/testing/tslint.json",
	})

	if !exists {
		t.Errorf("Missing initialized file: %s", missing)
	}
}

func TestConfig(t *testing.T) {
	bp, tp, _ := values()

	config(bp, tp)

	missing, exists := doFilesExist([]string{
		"/tmp/jamf_test/initializer/.dockerignore",
		"/tmp/jamf_test/initializer/.env.example",
		"/tmp/jamf_test/initializer/.gitignore",
		"/tmp/jamf_test/initializer/node.sh",
		"/tmp/jamf_test/initializer/nodemon.json",
		"/tmp/jamf_test/initializer/package-lock.json",
		"/tmp/jamf_test/initializer/package.json",
		"/tmp/jamf_test/initializer/README.md",
		"/tmp/jamf_test/initializer/tsconfig.json",
		"/tmp/jamf_test/initializer/tslint.json",
	})

	if !exists {
		t.Errorf("Missing config file: %s", missing)
	}
}

func TestServer(t *testing.T) {
	_, tp, _ := values()

	server("testing", tp, 4200)

	missing, exists := doFilesExist([]string{
		"/tmp/jamf_test/initializer/src",
		"/tmp/jamf_test/initializer/src/server.ts",
		"/tmp/jamf_test/initializer/index.ts",
	})

	if !exists {
		t.Errorf("Missing server file: %s", missing)
	}
}

func TestDockerize(t *testing.T) {
	_, tp, _ := values()

	dockerize("testing", tp, 4200)

	missing, exists := doFilesExist([]string{
		"/tmp/jamf_test/initializer/Dockerfile",
		"/tmp/jamf_test/initializer/docker-compose.yml",
		"/tmp/jamf_test/initializer/docker-compose.override.yml.example",
	})

	if !exists {
		t.Errorf("Missing docker file: %s", missing)
	}
}
