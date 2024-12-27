package apis

import (
	"Linda/baselibs/abstractions/xlog"
	"Linda/protocol/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var (
	logger = xlog.NewForPackage()
)

// 用于放在各种 handlerFunc 中，对抛出的异常进行处理
func httpRecover(w http.ResponseWriter, _ *http.Request) {
	if e := recover(); e != nil {
		switch err := e.(type) {
		case error:
			{
				processError(w, err)
			}
		default:
			{
				logger.Fatalf("Panic no error value, value is %v", err)
			}
		}
	}
}

func processError(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Debugf("record not found, write 404 NotFound, err is %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	logger.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(models.Serialize(map[string]any{
		"errormsg": err.Error(),
	}))
}
