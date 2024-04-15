package agents

import (
	"Linda/protocol/models"
	"errors"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 4096,
	ReadBufferSize:  4096,
}

type Mgr interface {
	NewNode(nodeId string, w http.ResponseWriter, r *http.Request)
	RemoveNode(nodeId string) error
}

type agentsmgr struct {
	agents      map[string]Agent
	agentsRWMut sync.RWMutex
}

func (mgr *agentsmgr) NewNode(nodeId string, w http.ResponseWriter, r *http.Request) {
	defer func(w http.ResponseWriter, _ *http.Request) {
		if err := recover(); err != nil {
			logrus.Error(string(debug.Stack()))
			logrus.Errorf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write(models.Serialize(&models.HeartBeatStartResponse{
				Result: err.(error).Error(),
			})); err != nil {
				logrus.Error(err)
				panic(err)
			}
		}
	}(w, r)

	var (
		agent Agent = nil
		err   error = nil
	)

	{
		mgr.agentsRWMut.Lock()
		defer mgr.agentsRWMut.Unlock()
		if _, exist := mgr.agents[nodeId]; exist {
			panic(errors.New("nodeId exists"))
		}
		agent, err = NewAgent(nodeId, xerr.Must(upgrader.Upgrade(w, r, nil)), mgr)
		if err != nil {
			logrus.Error(err)
			return
		}
		mgr.agents[nodeId] = agent
	}
	agent.StartServe()
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
