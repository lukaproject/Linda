package agents

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/logic/comm"
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
}

type agentHolder struct {
	conn    *websocket.Conn
	nodeId  string
	bagName string

	lastHBTime      time.Time
	lastSeqId       int64
	maxRunningTasks int

	hbAgent     chan *models.HeartBeatFromAgent
	tasksClient taskqueueclient.Client
}

func (ah *agentHolder) StartServe() {
	go ah.serveLoop()
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
				ah.processFinishedTask(msg)
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
	numOfRestResource := ah.maxRunningTasks - len(hbFromAgent.RunningTaskNames)
	for i := 0; i < numOfRestResource; i++ {
		taskName, err := ah.tasksClient.Deque(ah.bagName)
		if err != nil {
			logrus.Error(err)
			break
		}
		if hb.ScheduledTaskNames == nil {
			hb.ScheduledTaskNames = make([]string, 0)
		}
		hb.ScheduledTaskNames = append(hb.ScheduledTaskNames, taskName)
	}
	if hb.ScheduledTaskNames != nil {
		logrus.Infof("taskNames scheduled to %s is %v", ah.nodeId, hb.ScheduledTaskNames)
		ah.processScheduledTask(hb.ScheduledTaskNames)
	} else {
		logrus.Infof("no task scheduled to %s", ah.nodeId)
	}
	return hb
}

func (ah *agentHolder) processFinishedTask(msg *models.HeartBeatFromAgent) (err error) {
	if len(msg.FinishedTaskNames) > 0 {
		go comm.
			GetAsyncWorksInstance().
			PersistFinishedTasks(ah.bagName, msg.FinishedTaskNames)
	}
	return
}

func (ah *agentHolder) processScheduledTask(scheduledTaskNames []string) (err error) {
	if len(scheduledTaskNames) > 0 {
		go comm.
			GetAsyncWorksInstance().
			PersistScheduledTasks(ah.bagName, scheduledTaskNames, ah.nodeId)
	}
	return
}

func NewAgent(nodeId string, conn *websocket.Conn) (Agent, error) {
	ah := &agentHolder{
		conn:        conn,
		nodeId:      nodeId,
		hbAgent:     make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId:   -1,
		tasksClient: taskqueueclient.NewRedisTaskQueueClient(config.Instance().Redis),
	}
	hbStart := &models.HeartBeatStart{}
	err := hbconn.ReadMessage(ah.conn, hbStart)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Infof("nodeId %s, BagName %s", nodeId, hbStart.Node.BagName)
	ah.bagName = hbStart.Node.BagName
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
