package comm

import (
	"Linda/baselibs/abstractions/xos"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

const blockSize = 1 << 12

type FileSaver interface {
	WriteWithReader(filepath string, reader io.Reader) error
	ReadFrom(filepath string, writer io.Writer) error
}

type localFileSaver struct{}

func (lfs *localFileSaver) WriteWithReader(filepath string, reader io.Reader) error {
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
	p := make([]byte, blockSize)
	for {
		n, err := reader.Read(p)
		if n > 0 {
			n1 := xerr.Must(f.Write(p[:n]))
			if n1 != n {
				logrus.Warnf("write size %d != read size %d", n1, n)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}
	return nil
}

func (lfs *localFileSaver) ReadFrom(filepath string, writer io.Writer) error {
	if !xos.PathExists(filepath) {
		logrus.Errorf("%s not exist", filepath)
		return os.ErrNotExist
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	p := make([]byte, blockSize)
	for {
		n, err := f.Read(p)
		if n > 0 {
			n1 := xerr.Must(writer.Write(p[:n]))
			if n1 != n {
				logrus.Warnf("write size %d != read size %d", n1, n)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}
	return nil
}

func NewLocalFileSaver() FileSaver {
	return &localFileSaver{}
}
