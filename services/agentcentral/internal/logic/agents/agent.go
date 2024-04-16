package agents

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"
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

	lastHBTime time.Time
	lastSeqId  int64

	hbAgent chan *models.HeartBeatFromAgent
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

				xerr.Must0(
					hbconn.WriteMessage(
						ah.conn,
						&models.HeartBeatFromServer{
							SeqId: msg.SeqId,
						}),
				)
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

func NewAgent(nodeId string, conn *websocket.Conn) (Agent, error) {
	ah := &agentHolder{
		conn:      conn,
		nodeId:    nodeId,
		hbAgent:   make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId: -1,
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
	err = hbconn.WriteMessage(ah.conn, &models.HeartBeatStartResponse{
		Result: models.OK,
	})
	if err != nil {
		return nil, err
	}
	go ah.readHBLoop()
	return ah, nil
}
