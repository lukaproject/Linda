package filemanager

import (
	"io/fs"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Operator struct {
}

func (o *Operator) ListFileNames(dirname string, files chan string) {
	go func(fch chan string) {
		filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				logrus.Error(err)
			}
			if !info.IsDir() {
				fch <- path
			}
			return nil
		})
		close(fch)
	}(files)
}
