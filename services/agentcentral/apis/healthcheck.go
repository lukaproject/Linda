package apis

import (
	"net/http"

	"github.com/gorilla/mux"
)

func EnableHealthCheck(r *mux.Router) {
	r.HandleFunc("/api/healthcheck", healthCheck).Methods(http.MethodPost)
}

// healthCheck godoc
//
//	@Summary		health check
//	@Description	health check
//	@Accept			json
//	@Produce		plain
//	@Router			/healthcheck [post]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Debug("health check success!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
