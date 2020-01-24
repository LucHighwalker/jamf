package initializer

import (
	"log"
	"path"

	"nor/helper"
)

func Initialize(nd, wd string) {
	configPath := path.Join(nd, "boiler", "config")
	tempPath := path.Join(nd, "__temp__")

	err := helper.CopyDir(configPath, tempPath)

	if err != nil {
		log.Fatal(err)
	}
}
