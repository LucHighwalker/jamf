package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Copies a single file to destination.
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

// Copies all files in a directory to destination.
func CopyDir(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		Copy(path.Join(src, f.Name()), path.Join(dst, f.Name()))
	}
	return nil
}

// Ensures the given directory exists, otherwise creates it.
// https://siongui.github.io/2017/03/28/go-create-directory-if-not-exist/
func EnsureDirExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
