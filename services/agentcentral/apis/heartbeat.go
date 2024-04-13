package apis

import (
	"Linda/services/agentcentral/internal/logic"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func EnableHeartBeat(r *mux.Router) {
	r.HandleFunc("/api/agent/heartbeat/{nodeId}", heartbeat)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	logrus.Infof("connection start from %s", nodeId)
	logic.AgentsMgr().NewNode(nodeId, w, r)
}
