// this is a inner call api,
// only be used between agent and agentcentral
package apis

import (
	"Linda/services/agentcentral/internal/logic/agents"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableInnerCall(r *mux.Router) {
	r.HandleFunc("/api/agent/innercall/nodeidgen", nodeIdGen).Methods(http.MethodGet)
}

func nodeIdGen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(agents.GenNodeId()))
}
