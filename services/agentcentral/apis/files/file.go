package files

import (
	"Linda/baselibs/abstractions/xlog"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/logic/comm"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/lukaproject/xerr"
)

var (
	logger = xlog.NewForPackage()
)

const multipartFormMemoryLimit = (1 << 22)

func EnableFiles(r *mux.Router) {
	r.HandleFunc("/api/files/upload", upload).Methods(http.MethodPost)
	r.HandleFunc("/api/files/download/{block}/{fileName}", download).Methods(http.MethodGet)
}

// Upload file
//
//	@Summary		Upload file
//	@Description	Upload file, can only upload 1 file.
//	@Tags			files
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			fileName	formData	string	true	"file name"
//	@Param			block		formData	string	true	"file block"
//	@Param			file		formData	file	true	"this is a file"
//	@Success		200			{string}	string	"OK"
//	@Router			/files/upload [post]
func upload(w http.ResponseWriter, r *http.Request) {
	xerr.Must0(r.ParseMultipartForm(multipartFormMemoryLimit))
	xerr.Must0(r.ParseForm())
	filepath := path.Join(config.Instance().FileSaver.RootDir, r.Form.Get("block"), r.Form.Get("fileName"))
	logger.Debugf("upload file to %s", filepath)
	fileHeader, ok := r.MultipartForm.File["file"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, err := fileHeader[0].Open()
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	comm.GetFileSaverInstance().WriteWithReader(filepath, file)
}

// download file
//
//	@Summary		download file
//	@Description	download file
//	@Tags			files
//	@Accept			json
//	@Produce		octet-stream
//	@Param			fileName	path	string	true	"file name"
//	@Param			block		path	string	true	"block name"
//	@Router			/files/download/{block}/{fileName} [get]
func download(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/octet-stream")
	block, fileName := mux.Vars(r)["block"], mux.Vars(r)["fileName"]
	filepath := path.Join(config.Instance().FileSaver.RootDir, block, fileName)
	err := comm.GetFileSaverInstance().ReadFrom(filepath, w)
	if err != nil {
		if err == os.ErrNotExist {
			w.WriteHeader(http.StatusNotFound)
		} else {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}
