package filemanager

import (
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type fileDownloader struct{}

func (d *fileDownloader) Download(url, targetPath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(targetPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	blockSize := 1 << 12
	p := make([]byte, blockSize)
	for {
		n, err := resp.Body.Read(p)
		if err != nil {
			logrus.Error(err)
			return err
		}
		n1, err := f.Write(p[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			logrus.Error(err)
			return err
		}
		if n1 != n {
			logrus.Errorf("write not equal, read %d, write %d", n, n1)
			break
		}
	}
	defer f.Close()
	return
}
