package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Copy - Copies a single file to destination.
// https://opensource.com/article/18/6/copying-files-go
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// CopyDir - Copies all files in a directory to destination.
func CopyDir(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	EnsureDirExists(dst)

	for _, f := range files {
		if f.IsDir() {
			EnsureDirExists(path.Join(dst, f.Name()))
			err := CopyDir(path.Join(src, f.Name()), path.Join(dst, f.Name()))
			if err != nil {
				return err
			}
		} else if f.Name() != ".json" {
			_, err := Copy(path.Join(src, f.Name()), path.Join(dst, f.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DoesDirExist - Returns whether or not the directory exists.
func DoesDirExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

// EnsureDirExists - Creates directory if it doesn't exist.
// https://siongui.github.io/2017/03/28/go-create-directory-if-not-exist/
func EnsureDirExists(dir string) {
	if !DoesDirExist(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// Capitalize - Capitalizes the first letter in the string.
func Capitalize(s string) string {
	return strings.Title(s)
}

// GetContent - Returns the bytes contained in a given file.
func GetContent(src string) []byte {
	content, err := ioutil.ReadFile(src)

	if err != nil {
		panic(err)
	}

	return content
}
