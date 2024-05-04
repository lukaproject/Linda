package utils

import (
	"io/fs"
	"os"

	"github.com/lukaproject/xerr"
)

func MkdirAll(path string, mode fs.FileMode) {
	if IsDir(path) {
		return
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			xerr.Must0(os.MkdirAll(path, mode))
		} else {
			panic(err)
		}
	}
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
