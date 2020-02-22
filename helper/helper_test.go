package helper

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCopy(t *testing.T) {
	os.RemoveAll("/tmp/nor_test")

	src := "/tmp/nor_test/something"
	dest := "/tmp/nor_test/something_copied"

	os.MkdirAll(src, 0755)
	os.MkdirAll(dest, 0755)

	ioutil.WriteFile(path.Join(src, "something.txt"), []byte("some data"), 0644)

	_, err := Copy(path.Join(src, "something.txt"), path.Join(dest, "something.txt"))

	if err != nil {
		t.Errorf("Error while Copying file: %s", err.Error())
	}

	dir, err := ioutil.ReadDir(dest)

	if err != nil {
		t.Errorf("Error while reading destination: %s", err.Error())
	}

	if len(dir) == 0 || dir[0].Name() != "something.txt" {
		t.Error("Failed to copy file.")
	}
}

func TestCopyDir(t *testing.T) {
	os.RemoveAll("/tmp/nor_test")

	src := "/tmp/nor_test/something"
	dest := "/tmp/nor_test/something_copied"

	os.MkdirAll(src, 0755)
	os.MkdirAll(dest, 0755)

	ioutil.WriteFile(path.Join(src, "something.txt"), []byte("some data"), 0644)
	ioutil.WriteFile(path.Join(src, "somethingElse.txt"), []byte("some other data"), 0644)
	ioutil.WriteFile(path.Join(src, "somethingMore.txt"), []byte("some more data"), 0644)

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
	os.RemoveAll("/tmp/nor_test")

	os.MkdirAll("/tmp/nor_test", 0755)

	if DoesDirExist("/tmp/nor_test/some_dir") == true {
		t.Error("False positive.")
	}

	os.MkdirAll("/tmp/nor_test/some_dir", 0755)

	if DoesDirExist("/tmp/nor_test/some_dir") == false {
		t.Error("False negative.")
	}
}

func TestEnsureDirExists(t *testing.T) {
	os.RemoveAll("/tmp/nor_test")

	EnsureDirExists("/tmp/nor_test/should_exist")

	stat, err := os.Stat("/tmp/nor_test/should_exist")

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
	os.RemoveAll("/tmp/nor_test")

	src := "/tmp/nor_test/something"

	os.MkdirAll(src, 0755)

	ioutil.WriteFile(path.Join(src, "something.txt"), []byte("some data"), 0644)

	if GetContent(path.Join(src, "something.txt")) == nil {
		t.Error("Failed to get content.")
	}
}
