package xos

import (
	"io/fs"
	"os"

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
