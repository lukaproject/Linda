package agents

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type Agent interface {
	// 启动 Agent 守护协程，维持与 agent 的心跳
	StartServe()

	// 安排task
	AssignTask(taskId string) error
}

type agentHolder struct {
	conn   *websocket.Conn
	nodeId string
	bagId  string

	lastHBTime      time.Time
	lastSeqId       int64
	maxRunningTasks int

	hbAgent     chan *models.HeartBeatFromAgent
	tasksClient taskqueueclient.Client
}

func (ah *agentHolder) StartServe() {
	go ah.serveLoop()
}

func (ah *agentHolder) AssignTask(taskId string) (err error) {
	return
}

func (ah *agentHolder) serveLoop() {
	defer ah.recoverWSPanic()

	for {
		select {
		case msg := <-ah.hbAgent:
			{
				logrus.Infof("seqid %d", msg.SeqId)
				ah.lastSeqId = msg.SeqId
				ah.lastHBTime = time.Now()
				hbFromServer := ah.packHeartBeatResponse(msg)
				xerr.Must0(hbconn.WriteMessage(ah.conn, hbFromServer))
			}
		case <-time.After(5 * time.Second):
			{
				logrus.Errorf("hb timeout, nodeId %s", ah.nodeId)
				mgrInstance.RemoveNode(ah.nodeId)
				return
			}
		}
	}
}

func (ah *agentHolder) readHBLoop() {
	defer ah.recoverWSPanic()

	for {
		hb := &models.HeartBeatFromAgent{}
		xerr.Must0(hbconn.ReadMessage(ah.conn, hb))
		ah.hbAgent <- hb
	}
}

func (ah *agentHolder) recoverWSPanic() {
	if err := recover(); err != nil {
		logrus.Error(string(debug.Stack()), err)
		mgrInstance.RemoveNode(ah.nodeId)
	}
}

func (ah *agentHolder) packHeartBeatResponse(hbFromAgent *models.HeartBeatFromAgent) (hb *models.HeartBeatFromServer) {
	hb = &models.HeartBeatFromServer{
		SeqId: hbFromAgent.SeqId,
	}
	return hb
}

func NewAgent(nodeId string, conn *websocket.Conn) (Agent, error) {
	ah := &agentHolder{
		conn:        conn,
		nodeId:      nodeId,
		hbAgent:     make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId:   -1,
		tasksClient: taskqueueclient.NewRedisTaskQueueClient(config.TestConfig().Redis),
	}
	hbStart := &models.HeartBeatStart{}
	err := hbconn.ReadMessage(ah.conn, hbStart)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Debugf("nodeId %s, BagId %s", nodeId, hbStart.Node.BagId)
	ah.bagId = hbStart.Node.BagId
	ah.lastHBTime = time.Now()
	ah.maxRunningTasks = max(hbStart.Node.MaxRunningTasks, 1)
	err = hbconn.WriteMessage(ah.conn, &models.HeartBeatStartResponse{
		Result: models.OK,
	})
	if err != nil {
		return nil, err
	}
	go ah.readHBLoop()
	return ah, nil
}
