package agents

import (
	"Linda/services/agentcentral/internal/logic/comm"
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		WriteBufferSize: 4096,
		ReadBufferSize:  4096,
	}
)

type Mgr interface {
	NewNode(nodeId string, w http.ResponseWriter, r *http.Request)
	RemoveNode(nodeId string) error
}

type agentsmgr struct {
	agents      map[string]Agent
	agentsRWMut sync.RWMutex
}

func (mgr *agentsmgr) NewNode(nodeId string, w http.ResponseWriter, r *http.Request) {
	if agent, err := mgr.addNewNodeToMem(nodeId, w, r); err != nil {
		logrus.Error(err)
		return
	} else {
		agent.StartServe()
	}
}

func (mgr *agentsmgr) addNewNodeToMem(
	nodeId string,
	w http.ResponseWriter,
	r *http.Request,
) (agent Agent, err error) {
	comm.NewRLocker(&mgr.agentsRWMut).Run(func() {
		// logrus.Infof("validate new node %s", nodeId)
		if _, exist := mgr.agents[nodeId]; exist {
			panic(errors.New("nodeId exists"))
		}
		agent, err = NewAgent(nodeId, xerr.Must(upgrader.Upgrade(w, r, nil)))
		if err != nil {
			logrus.Error(err)
			return
		}
		mgr.agents[nodeId] = agent
	})
	return
}

func (mgr *agentsmgr) RemoveNode(nodeId string) error {
	mgr.agentsRWMut.Lock()
	defer mgr.agentsRWMut.Unlock()
	delete(mgr.agents, nodeId)
	logrus.Debugf("node %s removed", nodeId)
	return nil
}

func NewMgr() Mgr {
	return &agentsmgr{
		agents: make(map[string]Agent),
	}
}
