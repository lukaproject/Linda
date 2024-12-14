package filemanager

import (
	"Linda/agent/internal/config"
	"Linda/baselibs/abstractions/xio"
	"Linda/baselibs/abstractions/xos"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type Downloader struct {
}

func (d *Downloader) DownloadFromPublicURL(url, targetPath string) (err error) {
	resp, err := d.getClient().Do(d.newDefaultRequest(url))
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()

	xos.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	f, err := os.Create(targetPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer f.Close()
	return xio.Transport(resp.Body, f)
}

func (d *Downloader) DownloadFromInternal(block, fileName, targetPath string) (err error) {
	url := fmt.Sprintf("%s/files/download/%s/%s", config.Instance().AgentAPIUrl("http"), block, fileName)
	xos.MkdirAll(filepath.Dir(targetPath), fs.ModePerm)
	f, err := os.Create(targetPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer f.Close()
	resp, err := d.getClient().Do(d.newDefaultRequest(url))
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	return xio.Transport(resp.Body, f)
}

func (d *Downloader) newDefaultRequest(url string) *http.Request {
	req := xerr.Must(http.NewRequest("GET", url, nil))
	req.Header.Add("x-luka-source", "agent-file-downloader")
	return req
}

func (d *Downloader) getClient() *http.Client {
	return http.DefaultClient
}
