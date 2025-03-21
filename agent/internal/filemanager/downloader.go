package filemanager

import (
	"Linda/baselibs/abstractions/xio"
	"Linda/baselibs/abstractions/xos"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lukaproject/xerr"
)

type Downloader struct {
}

func (d *Downloader) DownloadFromPublicURL(url, targetPath string) (err error) {
	resp, err := d.getClient().Do(d.newDefaultRequest(url))
	if err != nil {
		logger.Error(err)
		return
	}
	defer resp.Body.Close()

	xos.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	f, err := os.Create(targetPath)
	if err != nil {
		logger.Error(err)
		return
	}
	defer f.Close()
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
