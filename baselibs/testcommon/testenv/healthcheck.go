package testenv

import (
	"Linda/baselibs/abstractions/xlog"
	"Linda/protocol/models"
	"io"
	"net/http"
)

var logger = xlog.NewForPackage()

func HealthCheck(url string) bool {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		logger.Error(err)
		return false
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return false
	}

	return string(b) == models.OK
}
