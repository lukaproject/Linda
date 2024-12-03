package fakefileserver

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
)

var (
	portsList = []int16{
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
	t    *testing.T
	root string

	schema string
	host   string
	port   int16

	server   *http.Server
	listener net.Listener
}

func (fs *fileServer) Init() {
	fs.root = fs.t.TempDir()
	fs.schema = "http"
	fs.host = "127.0.0.1"
	fs.port = portsList[rand.Intn(len(portsList))]
	serveAddr := fmt.Sprintf(":%d", fs.port)
	fs.listener = xerr.Must(net.Listen("tcp", serveAddr))

	mux := http.NewServeMux()
	mux.Handle(
		PREFIX,
		http.StripPrefix(
			PREFIX,
			http.FileServer(http.Dir(fs.root))))
	mux.HandleFunc("/upload", uploadFile)

	fs.server = &http.Server{
		Addr:    serveAddr,
		Handler: mux,
	}
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
	go func() {
		log.Printf("file server is started in port %d", fs.port)
		if err := fs.server.Serve(fs.listener); err != nil {
			log.Println(err)
		}
	}()
	// wait for file server serve.
	<-time.After(10 * time.Millisecond)
}

func (fs *fileServer) Shutdown() {
	xerr.Must0(fs.server.Shutdown(context.Background()))
}

func (fs *fileServer) Addr() string {
	return fmt.Sprintf("%s://%s:%d", fs.schema, fs.host, fs.port)
}
