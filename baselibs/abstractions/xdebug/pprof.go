package xdebug

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/gorilla/mux"
)

func EnablePprof(r *mux.Router) {
	r.PathPrefix("/debug/pprof/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, "/cmdline") {
				pprof.Cmdline(w, r)
			} else if strings.HasSuffix(r.URL.Path, "/profile") {
				pprof.Profile(w, r)
			} else if strings.HasSuffix(r.URL.Path, "/trace") {
				pprof.Trace(w, r)
			} else if strings.HasSuffix(r.URL.Path, "/symbol") {
				pprof.Symbol(w, r)
			} else {
				pprof.Index(w, r)
			}
		})
}
