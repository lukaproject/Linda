package fakefileserver

import (
	"Linda/baselibs/multifs"
	"path"
	"testing"
)

func BuildDownloadURL(s FileServer, filePath string) string {
	return s.Addr() + path.Join(multifs.FILE_DOWNLOAD_PREFIX, filePath)
}

func NewT(t *testing.T) FileServer {
	fs := &fileServer{
		t: t,
	}
	fs.Init()
	return fs
}

func StartT(t *testing.T) FileServer {
	fs := &fileServer{
		t: t,
	}
	fs.Init()
	fs.Start()
	return fs
}
