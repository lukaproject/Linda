package multifs

import (
	"Linda/baselibs/multifs/internal/local"
	"Linda/protocol/xhttp"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lukaproject/xerr"
)

const (
	FILE_DOWNLOAD_PREFIX = "/files/"
	FILE_UPLOAD_PREFIX   = "/upload/"
)

type FileService struct {
	Port    int
	BaseDir string

	server   *http.Server
	listener net.Listener
}

func (fs *FileService) Start() {
	go func() {
		logger.Infof("file server is started in port %d", fs.Port)
		if err := fs.server.Serve(fs.listener); err != nil {
			logger.Info(err)
		}
	}()
	// wait for file server serve.
	<-time.After(10 * time.Millisecond)
}

func (fs *FileService) Shutdown(ctx context.Context) error {
	return fs.server.Shutdown(ctx)
}

type FileServiceType int

const (
	FileServiceType_Local FileServiceType = iota
)

type NewFileServiceInput struct {
	Port    int
	BaseDir string
	Type    FileServiceType
}

func New(input NewFileServiceInput) *FileService {
	if input.Type == FileServiceType_Local {
		return newLocalFileService(input)
	}
	panic("not implementation")
}

func newLocalFileService(input NewFileServiceInput) *FileService {
	fs := &FileService{
		Port:    input.Port,
		BaseDir: input.BaseDir,
	}
	serveAddr := fmt.Sprintf(":%d", fs.Port)
	fs.listener = xerr.Must(net.Listen("tcp", serveAddr))

	mux := http.NewServeMux()
	mux.Handle(
		FILE_DOWNLOAD_PREFIX,
		http.StripPrefix(
			FILE_DOWNLOAD_PREFIX,
			&local.FileDownloader{
				BaseDir: fs.BaseDir,
			}))
	mux.Handle(FILE_UPLOAD_PREFIX,
		http.StripPrefix(
			FILE_UPLOAD_PREFIX,
			&local.FileUploader{
				BaseDir: fs.BaseDir,
			}))

	mux.HandleFunc("/api/healthcheck", xhttp.Healthcheck)

	fs.server = &http.Server{
		Addr:    serveAddr,
		Handler: mux,
	}

	return fs
}
