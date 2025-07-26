// this is a inner call api,
// only be used between agent and agentcentral
package apis

import (
	"Linda/baselibs/abstractions/serviceskit/generator"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableInnerCall(r *mux.Router) {
	r.HandleFunc("/api/agent/innercall/nodeidgen", nodeIdGen).Methods(http.MethodGet)
	r.HandleFunc("/api/agent/innercall/bags/{bagName}/tasks/{taskName}", innerCallGetTask).Methods(http.MethodPost)
}

// nodeIdGen
// be used in agent start to fetch node id if not exists.
func nodeIdGen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(generator.GetInstance().NodeId()))
}

// innerCallGetTask
// be used in agent fetch task
func innerCallGetTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bagName := mux.Vars(r)["bagName"]
	taskName := mux.Vars(r)["taskName"]
	accessKey := r.Form.Get("accessKey")
	logger.Infof(
		"bagName=%s, taskName=%s, accessKey=%s",
		bagName, taskName, accessKey)
	taskModel := db.
		NewDBOperations().
		Tasks.
		GetByAccessKey(bagName, taskName, accessKey)
	w.Write(models.Serialize(taskModel))
}
