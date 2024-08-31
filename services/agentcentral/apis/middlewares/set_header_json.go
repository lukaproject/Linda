package middlewares

import (
	"net/http"
)

// SetHeaderJSON
// a middleware to set write header as content-type: json/application
func SetHeaderJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
