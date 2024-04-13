package agents

import (
	"Linda/protocol/models"
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type Agent interface {
	StartServe()
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

func (ah *agentHolder) serveLoop() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case msg := <-ah.hbAgent:
			{
				logrus.Infof("seqid %d", msg.SeqId)
				xerr.Must0(
					ah.conn.WriteMessage(
						websocket.BinaryMessage,
						models.Serialize(&models.HeartBeatFromServer{
							SeqId: msg.SeqId,
						})))
			}
		case <-tick.C:
			{
				logrus.Errorf("hb timeout, nodeId %s", ah.nodeId)
				return
			}
		}
	}
}

func (ah *agentHolder) readHBLoop() {
	for {
		msgType, body, err := ah.conn.ReadMessage()

		if err != nil {
			logrus.Error(err)
			break
		}
		if msgType != websocket.BinaryMessage {
			logrus.Errorf("msgType is invalid, %d", msgType)
			return
		}

		hb := &models.HeartBeatFromAgent{}
		models.Deserialize(body, hb)
		ah.hbAgent <- hb
	}
}

func NewAgent(nodeId string, conn *websocket.Conn) (Agent, error) {
	ah := &agentHolder{
		conn:      conn,
		nodeId:    nodeId,
		hbAgent:   make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId: -1,
	}

	msgType, body, err := ah.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	if msgType != websocket.BinaryMessage {
		return nil, errors.New("msgType is not binary")
	}

	hbStart := &models.HeartBeatStart{}
	models.Deserialize(body, hbStart)
	logrus.Debugf("nodeId %s, BagId %s", nodeId, hbStart.Node.BagId)
	ah.bagId = hbStart.Node.BagId
	ah.lastHBTime = time.Now()
	err = ah.conn.WriteMessage(
		websocket.BinaryMessage,
		models.Serialize(&models.HeartBeatStartResponse{
			Result: models.OK,
		}))
	if err != nil {
		return nil, err
	}
	go ah.readHBLoop()
	return ah, nil
}
