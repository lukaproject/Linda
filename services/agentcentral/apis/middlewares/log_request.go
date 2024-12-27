package middlewares

import (
	"Linda/baselibs/abstractions/xlog"
	"net/http"
	"strings"
)

var logger = xlog.NewForPackage()

// LogRequest
// Log comming requests info.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			logger.Infof(
				"NewRequest Method %s, Proto %s, URL %s, UserAgent %s",
				r.Method, r.Proto, r.URL, r.UserAgent())
		}
		next.ServeHTTP(w, r)
	})
}
