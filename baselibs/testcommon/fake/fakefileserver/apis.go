package fakefileserver

import (
	"path/filepath"
	"testing"
)

const (
	PREFIX = "/files/"
)

func BuildDownloadURL(s FileServer, path string) string {
	return s.Addr() + filepath.Join(PREFIX, path)
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
