package apis

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func EnableHealthCheck(r *mux.Router) {
	r.HandleFunc("/api/healthcheck", healthCheck)
}

// healthCheck godoc
//
//	@Summary		health check
//	@Description	health check
//	@Accept			json
//	@Produce		plain
//	@Router			/healthcheck [post]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logrus.Debug("health check success!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
