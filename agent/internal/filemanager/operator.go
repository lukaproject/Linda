package filemanager

import (
	"io/fs"
	"path/filepath"
)

type Operator struct {
}

func (o *Operator) ListFileNames(dirname string, files chan string) {
	go func(fch chan string) {
		filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				logger.Error(err)
			}
			if !info.IsDir() {
				fch <- path
			}
			return nil
		})
		close(fch)
	}(files)
}
