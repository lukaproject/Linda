package testenv

import (
	"Linda/protocol/models"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HealthCheck(url string) bool {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return false
	}

	return string(b) == models.OK
}
