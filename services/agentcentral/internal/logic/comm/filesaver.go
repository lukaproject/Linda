package comm

import (
	"Linda/baselibs/abstractions/xio"
	"Linda/baselibs/abstractions/xos"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type FileSaver interface {
	WriteWithReader(filepath string, reader io.Reader) error
	ReadFrom(filepath string, writer io.Writer) error
}

type localFileSaver struct{}

func (lfs *localFileSaver) WriteWithReader(filepath string, reader io.Reader) (err error) {
	if !xos.PathExists(path.Dir(filepath)) {
		err := os.MkdirAll(path.Dir(filepath), fs.ModePerm)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = xio.Transport(reader, f)
	return
}

func (lfs *localFileSaver) ReadFrom(filepath string, writer io.Writer) (err error) {
	if !xos.PathExists(filepath) {
		logrus.Errorf("%s not exist", filepath)
		return os.ErrNotExist
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = xio.Transport(f, writer)
	return
}

func NewLocalFileSaver() FileSaver {
	return &localFileSaver{}
}
