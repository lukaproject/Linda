package apis

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/logic"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 4096,
	ReadBufferSize:  4096,
}

func EnableHeartBeat(r *mux.Router) {
	r.HandleFunc("/api/agent/heartbeat/{nodeId}", heartbeat)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {

	defer func(w http.ResponseWriter, _ *http.Request) {
		if err := recover(); err != nil {
			logrus.Error(string(debug.Stack()))
			logrus.Errorf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write(models.Serialize(&models.HeartBeatStartResponse{
				Result: err.(error).Error(),
			})); err != nil {
				panic(err)
			}
		}
	}(w, r)

	nodeId := mux.Vars(r)["nodeId"]
	conn := xerr.Must(upgrader.Upgrade(w, r, nil))
	xerr.Must0(logic.AgentsMgr().NewNode(nodeId, conn))
}
