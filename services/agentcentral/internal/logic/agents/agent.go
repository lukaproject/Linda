package agents

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type Agent interface {
	// 启动 Agent 守护协程，维持与 agent 的心跳
	StartServe()
}

type agentHolder struct {
	conn   *websocket.Conn
	nodeId string
	bagId  string

	lastHBTime time.Time
	lastSeqId  int64

	hbAgent chan *models.HeartBeatFromAgent

	mgr Mgr
}

func (ah *agentHolder) StartServe() {
	go ah.serveLoop()
}

func (ah *agentHolder) serveLoop() {
	tick := time.NewTicker(5 * time.Second)
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
		case <-tick.C:
			{
				logrus.Errorf("hb timeout, nodeId %s", ah.nodeId)
				ah.mgr.RemoveNode(ah.nodeId)
				return
			}
		}
	}
}

func (ah *agentHolder) readHBLoop() {
	for {
		hb := &models.HeartBeatFromAgent{}

		if err := hbconn.ReadMessage(ah.conn, hb); err != nil {
			logrus.Error(err)
			continue
		}

		ah.hbAgent <- hb
	}
}

func NewAgent(nodeId string, conn *websocket.Conn, mgr Mgr) (Agent, error) {
	ah := &agentHolder{
		conn:      conn,
		nodeId:    nodeId,
		hbAgent:   make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId: -1,
		mgr:       mgr,
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
