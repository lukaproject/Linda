package xos

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/lukaproject/xerr"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func MkdirAll(path string, mode fs.FileMode) {
	if IsDir(path) {
		return
	}
	if !PathExists(path) {
		xerr.Must0(os.MkdirAll(path, mode))
	}
}

func ReadStringFromFile(path string) string {
	return string(xerr.Must(io.ReadAll(xerr.Must(os.Open(path)))))
}

// Touch
// a unix-like operation `touch xxxx`,
// to create an empty file in this path.
func Touch(path string) error {
	dir := filepath.Dir(path)
	MkdirAll(dir, fs.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		if os.IsExist(err) {
			return nil
		} else {
			return err
		}
	}
	f.Close()
	return nil
}
