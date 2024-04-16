package apis

import (
	"Linda/protocol/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

// 用于放在各种 handlerFunc 中，对抛出的异常进行处理
func httpRecover(w http.ResponseWriter, _ *http.Request) {
	if e := recover(); e != nil {
		switch err := e.(type) {
		case error:
			{
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(models.Serialize(map[string]any{
					"errormsg": err.Error(),
				}))
			}
		default:
			{
				logrus.Fatalf("Panic no error value, value is %v", err)
			}
		}
	}
}
