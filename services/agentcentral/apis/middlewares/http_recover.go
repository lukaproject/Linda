package middlewares

import (
	"Linda/protocol/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// 用于放在各种 handlerFunc 中，对抛出的异常进行处理
func HTTPRecover(next http.Handler) http.Handler {
	recoverFunc := func(w http.ResponseWriter, _ *http.Request) {
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverFunc(w, r)
		next.ServeHTTP(w, r)
	})
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
