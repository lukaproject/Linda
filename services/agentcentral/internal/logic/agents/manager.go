package agents

import (
	"Linda/baselibs/abstractions"
	"Linda/baselibs/abstractions/xctx"
	"Linda/baselibs/codes/errno"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

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

	// AddNodeToBag
	// 将 Node 加入 Bag 中.
	// 有可能会返回 error, 用于检查 Node 是否被其他的 Bag 占用
	AddNodeToBag(nodeId, bagName string) (err error)
	FreeNode(nodeId string)

	GetNodeInfo(nodeId string) *models.NodeInfo

	CallAgent(nodeId string, callFunc func(agent Agent) error) error

	List(query abstractions.ListQueryPacker) (chan *models.NodeInfo, error)
	WaitForFileListResponse(nodeId, operationId string, timeout time.Duration) (models.FileListResponse, error)
	WaitForFileGetResponse(nodeId, operationId string, timeout time.Duration) (*models.FileGetResponse, error)
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
	xctx.NewLocker(&mgr.agentsRWMut).Run(func() {
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
	xctx.NewLocker(&mgr.agentsRWMut).Run(func() {
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

func (mgr *agentsmgr) AddNodeToBag(nodeId, bagName string) (err error) {
	xctx.NewLocker(&mgr.agentsRWMut).Run(func() {
		if agent, exist := mgr.agents[nodeId]; exist {
			err = agent.Join(bagName)
		}
	})
	return
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

func (mgr *agentsmgr) WaitForFileListResponse(nodeId, operationId string, timeout time.Duration) (models.FileListResponse, error) {
	agent, exists := mgr.agents[nodeId]
	if !exists {
		return models.FileListResponse{}, fmt.Errorf("node %s not found", nodeId)
	}

	ah := agent.(*agentHolder)
	ah.responsesListMutex.Lock()
	responseChan, exists := ah.pendingFileListOps[operationId]
	ah.responsesListMutex.Unlock()

	if !exists {
		return models.FileListResponse{}, fmt.Errorf("operation %s not found", operationId)
	}

	select {
	case response := <-responseChan:
		return models.FileListResponse{
			OperationId: response.OperationId,
			Files:       response.Files,
			Error:       response.Error,
		}, nil
	case <-time.After(timeout):
		// Clean up pending operation
		ah.responsesListMutex.Lock()
		delete(ah.pendingFileListOps, operationId)
		ah.responsesListMutex.Unlock()
		return models.FileListResponse{}, fmt.Errorf("timeout waiting for response")
	}
}

func (mgr *agentsmgr) WaitForFileGetResponse(nodeId, operationId string, timeout time.Duration) (*models.FileGetResponse, error) {
	agent, exists := mgr.agents[nodeId]
	if !exists {
		return nil, fmt.Errorf("node %s not found", nodeId)
	}

	ah := agent.(*agentHolder)
	ah.responsesGetMutex.Lock()
	responseChan, exists := ah.pendingFileGetOps[operationId]
	ah.responsesGetMutex.Unlock()

	if !exists {
		return nil, fmt.Errorf("operation %s not found", operationId)
	}

	select {
	case response := <-responseChan:
		return &response, nil
	case <-time.After(timeout):
		// Clean up pending operation
		ah.responsesGetMutex.Lock()
		delete(ah.pendingFileGetOps, operationId)
		ah.responsesGetMutex.Unlock()
		logger.Errorf("timeout waiting for file get response for operation %s on node %s", operationId, nodeId)
		return nil, fmt.Errorf("timeout waiting for response")
	}
}

func NewMgr() Mgr {
	return &agentsmgr{
		agents: make(map[string]Agent),
	}
}
