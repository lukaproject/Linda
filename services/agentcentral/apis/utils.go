package apis

import (
	"Linda/protocol/models"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
				logrus.Fatalf("Panic no error value, value is %v", err)
			}
		}
	}
}

func processError(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Debugf("record not found, write 404 NotFound, err is %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	logrus.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(models.Serialize(map[string]any{
		"errormsg": err.Error(),
	}))
}
