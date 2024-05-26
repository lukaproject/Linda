package apis

import (
	"net/http"

	"github.com/gorilla/mux"
)

func EnableTasks(r *mux.Router) {
	r.HandleFunc("/api/bags/{bagName}/tasks", addTask).Methods(http.MethodPost)
	r.HandleFunc("/api/bags/{bagName}/tasks/{taskName}", getTask).Methods(http.MethodGet)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	// TODO
}

func getTask(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	// TODO
}
