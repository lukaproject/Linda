package fakefileserver

import (
	"Linda/baselibs/multifs"
	"path/filepath"
	"testing"
)

func BuildDownloadURL(s FileServer, path string) string {
	return s.Addr() + filepath.Join(multifs.FILE_DOWNLOAD_PREFIX, path)
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
