// this is a inner call api,
// only be used between agent and agentcentral
package apis

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/logic/agents"
	"Linda/services/agentcentral/internal/logic/tasks"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableInnerCall(r *mux.Router) {
	r.HandleFunc("/api/agent/innercall/nodeidgen", nodeIdGen).Methods(http.MethodGet)
	r.HandleFunc("/api/agent/innercall/bags/{bagName}/tasks/{taskName}", innerCallGetTask).Methods(http.MethodGet)
}

// nodeIdGen
// be used in agent start to fetch node id if not exists.
func nodeIdGen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(agents.GenNodeId()))
}

// innerCallGetTask
// be used in agent fetch task
func innerCallGetTask(w http.ResponseWriter, r *http.Request) {
	bagName := mux.Vars(r)["bagName"]
	taskName := mux.Vars(r)["taskName"]
	taskModel := tasks.
		GetBagsMgrInstance().
		GetTasksMgr(bagName).GetTask(taskName)
	w.Write(models.Serialize(taskModel))
}
