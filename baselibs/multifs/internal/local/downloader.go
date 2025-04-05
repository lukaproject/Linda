package local

import (
	"Linda/baselibs/abstractions/xio"
	"Linda/baselibs/abstractions/xos"
	"net/http"
	"os"
	"path"
)

type FileDownloader struct {
	BaseDir string
}

func (fs *FileDownloader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fpath := r.URL.Path
	logger.Infof("download file %s", fpath)

	absPath := path.Join(fs.BaseDir, fpath)
	if !xos.PathExists(absPath) {
		logger.Warnf("file %s not found", fpath)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f, err := os.Open(absPath)
	if err != nil {
		logger.Errorf("read file %s failed, err is %v", fpath, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if err = xio.Transport(f, w); err != nil {
		logger.Errorf("transfer file %s failed, err is %v", fpath, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
