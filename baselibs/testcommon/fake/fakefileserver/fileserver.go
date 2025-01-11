package fakefileserver

import (
	"Linda/baselibs/multifs"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/lukaproject/xerr"
)

var (
	portsList = []int{
		10000,
		10001,
		10002,
		10003,
		10004,
		10005,
	}
)

type FileServer interface {
	Init()

	Start()
	Shutdown()
	Addr() string

	AddFileContent(path, content string) error
	RemovePath(path string) error
}

type fileServer struct {
	t                *testing.T
	root             string
	schema           string
	host             string
	port             int
	localFileService *multifs.FileService
}

func (fs *fileServer) Init() {
	fs.root = fs.t.TempDir()
	fs.schema = "http"
	fs.host = "127.0.0.1"
	fs.port = portsList[rand.Intn(len(portsList))]
	fs.localFileService = multifs.New(
		multifs.NewFileServiceInput{
			Port:    fs.port,
			BaseDir: fs.root,
			Type:    multifs.FileServiceType_Local,
		})
	fs.t.Cleanup(fs.Shutdown)
	log.Printf("addr is %s://%s:%d", fs.schema, fs.host, fs.port)
}

func (fs *fileServer) AddFileContent(path, content string) error {
	interPath := filepath.Join(fs.root, path)
	dir := filepath.Dir(interPath)
	_ = os.MkdirAll(dir, os.ModePerm)

	f, err := os.Create(interPath)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func (fs *fileServer) RemovePath(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	return os.RemoveAll(path)
}

func (fs *fileServer) Start() {
	fs.localFileService.Start()
}

func (fs *fileServer) Shutdown() {
	xerr.Must0(fs.localFileService.Shutdown(context.Background()))
}

func (fs *fileServer) Addr() string {
	return fmt.Sprintf("%s://%s:%d", fs.schema, fs.host, fs.port)
}
