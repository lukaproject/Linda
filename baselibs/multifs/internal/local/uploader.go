package local

import (
	"Linda/baselibs/abstractions/xio"
	"Linda/baselibs/abstractions/xos"
	"net/http"
	"os"
	"path"
)

type FileUploader struct {
	BaseDir string
}

func (fs *FileUploader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fs.uploadFile(w, r)
}

func (fs *FileUploader) uploadFile(w http.ResponseWriter, r *http.Request) {
	fpath := r.URL.Path
	logger.Infof("upload file %s", fpath)
	absPath := path.Join(fs.BaseDir, fpath)

	fBody, fHeader, err := r.FormFile("file")
	if err != nil {
		logger.Errorf("file %s read file failed, err is %v", fpath, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Infof("file size is %d, filename from header is %s", fHeader.Size, fHeader.Filename)
	xos.MkdirAll(path.Dir(absPath), os.ModePerm)
	wFile, err := os.Create(absPath)
	if err != nil {
		logger.Errorf("create file, err is %v", err)
		if os.IsExist(err) {
			wFile, err = os.Open(absPath)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = nil
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	defer wFile.Close()

	if err = xio.Transport(fBody, wFile); err != nil {
		logger.Errorf("transfer file %s failed, err is %v", fpath, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
