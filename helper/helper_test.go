package helper

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCopy(t *testing.T) {
	os.RemoveAll("/tmp/jamf_test")

	src := "/tmp/jamf_test/helper"
	dest := "/tmp/jamf_test/helper_copied"

	os.MkdirAll(src, 0755)
	os.MkdirAll(dest, 0755)

	ioutil.WriteFile(path.Join(src, "helper.txt"), []byte("some data"), 0644)

	_, err := Copy(path.Join(src, "helper.txt"), path.Join(dest, "helper.txt"))

	if err != nil {
		t.Errorf("Error while Copying file: %s", err.Error())
	}

	dir, err := ioutil.ReadDir(dest)

	if err != nil {
		t.Errorf("Error while reading destination: %s", err.Error())
	}

	if len(dir) == 0 || dir[0].Name() != "helper.txt" {
		t.Error("Failed to copy file.")
	}
}

func TestCopyDir(t *testing.T) {
	os.RemoveAll("/tmp/jamf_test")

	src := "/tmp/jamf_test/helper"
	dest := "/tmp/jamf_test/helper_copied"

	os.MkdirAll(src, 0755)
	os.MkdirAll(dest, 0755)

	ioutil.WriteFile(path.Join(src, "helper.txt"), []byte("some data"), 0644)
	ioutil.WriteFile(path.Join(src, "helperElse.txt"), []byte("some other data"), 0644)
	ioutil.WriteFile(path.Join(src, "helperMore.txt"), []byte("some more data"), 0644)

	err := CopyDir(src, dest)

	if err != nil {
		t.Errorf("Error while Copying directory: %s", err.Error())
	}

	dir, err := ioutil.ReadDir(dest)

	if err != nil {
		t.Errorf("Error while reading destination: %s", err.Error())
	}

	if len(dir) != 3 {
		t.Error("Failed to copying directory.")
	}
}

func TestDoesDirExist(t *testing.T) {
	os.RemoveAll("/tmp/jamf_test")

	os.MkdirAll("/tmp/jamf_test", 0755)

	if DoesDirExist("/tmp/jamf_test/some_dir") == true {
		t.Error("False positive.")
	}

	os.MkdirAll("/tmp/jamf_test/some_dir", 0755)

	if DoesDirExist("/tmp/jamf_test/some_dir") == false {
		t.Error("False negative.")
	}
}

func TestEnsureDirExists(t *testing.T) {
	os.RemoveAll("/tmp/jamf_test")

	EnsureDirExists("/tmp/jamf_test/should_exist")

	stat, err := os.Stat("/tmp/jamf_test/should_exist")

	if err != nil {
		t.Errorf("Error ensuring directory existence: %s", err.Error())
	}

	if !stat.IsDir() {
		t.Error("Expected directory.")
	}
}

func TestCapitalize(t *testing.T) {
	if Capitalize("string") != "String" {
		t.Error("Failed capitalizing string.")
	}
}

func TestGetContent(t *testing.T) {
	os.RemoveAll("/tmp/jamf_test")

	src := "/tmp/jamf_test/helper"

	os.MkdirAll(src, 0755)

	ioutil.WriteFile(path.Join(src, "helper.txt"), []byte("some data"), 0644)

	if GetContent(path.Join(src, "helper.txt")) == nil {
		t.Error("Failed to get content.")
	}
}
