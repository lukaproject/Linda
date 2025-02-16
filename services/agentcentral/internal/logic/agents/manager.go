package agents

import (
	"Linda/baselibs/abstractions"
	"Linda/baselibs/codes/errno"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm"
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
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
	AddNodeToBag(nodeId, bagName string)
	FreeNode(nodeId string)

	GetNodeInfo(nodeId string) *models.NodeInfo

	CallAgent(nodeId string, callFunc func(agent Agent) error) error

	List(query abstractions.ListQueryPacker) (chan *models.NodeInfo, error)
}

type agentsmgr struct {
	agents      map[string]Agent
	agentsRWMut sync.RWMutex
}

func (mgr *agentsmgr) NewNode(nodeId string, w http.ResponseWriter, r *http.Request) {
	if agent, err := mgr.addNewNodeToMem(nodeId, w, r); err != nil {
		logger.Error(err)
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
	comm.NewLocker(&mgr.agentsRWMut).Run(func() {
		if _, exist := mgr.agents[nodeId]; exist {
			panic(errno.ErrNodeIdExists)
		}
		agent, err = NewAgent(nodeId, xerr.Must(upgrader.Upgrade(w, r, nil)))
		if err != nil {
			logger.Error(err)
			return
		}
		mgr.agents[nodeId] = agent
	})
	return
}

func (mgr *agentsmgr) RemoveNode(nodeId string) error {
	comm.NewLocker(&mgr.agentsRWMut).Run(func() {
		if _, ok := mgr.agents[nodeId]; ok {
			delete(mgr.agents, nodeId)
			db.NewDBOperations().NodeInfos.Delete(nodeId)
			logger.Debugf("node %s removed", nodeId)
		} else {
			logger.Debugf("node %s has been removed yet", nodeId)
		}
	})
	return nil
}

func (mgr *agentsmgr) AddNodeToBag(nodeId, bagName string) {
	comm.NewLocker(&mgr.agentsRWMut).Run(func() {
		if agent, exist := mgr.agents[nodeId]; exist {
			agent.Join(bagName)
		}
	})
}

func (mgr *agentsmgr) FreeNode(nodeId string) {
	mgr.agentsRWMut.Lock()
	defer mgr.agentsRWMut.Unlock()
	if agent, exist := mgr.agents[nodeId]; exist {
		agent.Free()
	} else {
		logger.Errorf("node %s not found", nodeId)
	}
}

func (mgr *agentsmgr) GetNodeInfo(nodeId string) *models.NodeInfo {
	mgr.agentsRWMut.Lock()
	defer mgr.agentsRWMut.Unlock()
	if agent, exist := mgr.agents[nodeId]; exist {
		return agent.GetInfo()
	}
	return nil
}

func (mgr *agentsmgr) CallAgent(nodeId string, callFunc func(agent Agent) error) error {
	mgr.agentsRWMut.RLock()
	defer mgr.agentsRWMut.RUnlock()
	if agentHolder, ok := mgr.agents[nodeId]; ok {
		return callFunc(agentHolder)
	} else {
		return errno.ErrAgentNotFound
	}
}

func (mgr *agentsmgr) List(queryPacker abstractions.ListQueryPacker) (chan *models.NodeInfo, error) {
	if queryPacker == nil {
		return nil, errors.New("query must be not nil")
	}

	return db.NewDBOperations().NodeInfos.List(queryPacker), nil
}

func NewMgr() Mgr {
	return &agentsmgr{
		agents: make(map[string]Agent),
	}
}
