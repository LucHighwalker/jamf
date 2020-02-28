package main

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestClearTemp(t *testing.T) {
	clearTemp()
	ioutil.WriteFile(path.Join(tempPath, "something.txt"), []byte("Some data."), 0644)
	clearTemp()
	dir, err := ioutil.ReadDir(tempPath)

	if err != nil {
		t.Errorf("Got error reading temp. Error: %s", err.Error())
	}

	if len(dir) > 0 {
		t.Error("Could not clear temp.")
	}
}

func TestGetGopath(t *testing.T) {
	getGopath()
	if gopath == "" || boilerPath == "" {
		t.Error("Could not get gopath.")
	}
}

func TestCheckBoiler(t *testing.T) {
	t.SkipNow()
}

func TestGetBoiler(t *testing.T) {
	t.SkipNow()
}

func TestInfo(t *testing.T) {
	info()
	if jamf.Name != "Just Another Mvc Framework" {
		t.Errorf("jamf.Name is false. Got: [%s]", jamf.Name)
	}
}

func TestCommands(t *testing.T) {
	commands()
	if len(jamf.Commands) == 0 {
		t.Error("Commands did not initialize.")
	}
}

func TestMain(t *testing.T) {
	t.SkipNow()
}
