package agents

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
)

type Agent interface {
	// 启动 Agent 守护协程，维持与 agent 的心跳
	StartServe()

	// 把当前Agent加入某个bag
	Join(bagName string)

	// 把当前Agent设置为空闲
	Free()

	// 上传文件到Node，但是不是及时的会有一定的延迟
	AddFilesUploadToNode(files []models.FileDescription)

	GetInfo() *models.NodeInfo
}

type agentHolder struct {
	conn       *websocket.Conn
	nodeId     string
	nodeName   string
	nodeStates *nodeStates

	lastHBTime      time.Time
	lastSeqId       int64
	maxRunningTasks int

	hbAgent     chan *models.HeartBeatFromAgent
	tasksClient taskqueueclient.Client

	noUploadFiles    []models.FileDescription
	noUploadFilesMut sync.Mutex
}

func (ah *agentHolder) GetInfo() *models.NodeInfo {
	return &models.NodeInfo{
		BagName:         ah.nodeStates.GetBagName(),
		MaxRunningTasks: ah.maxRunningTasks,
		NodeName:        ah.nodeName,
	}
}

func (ah *agentHolder) StartServe() {
	go ah.serveLoop()
}

func (ah *agentHolder) AddFilesUploadToNode(files []models.FileDescription) {
	ah.noUploadFilesMut.Lock()
	defer ah.noUploadFilesMut.Unlock()
	ah.noUploadFiles = append(ah.noUploadFiles, files...)
}

func (ah *agentHolder) Join(bagName string) {
	ah.nodeStates.Join(bagName)
}

func (ah *agentHolder) Free() {
	logger.Infof("free node %s", ah.nodeId)
	ah.nodeStates.Free()
}

func (ah *agentHolder) serveLoop() {
	defer ah.recoverWSPanic()
	for {
		select {
		case msg := <-ah.hbAgent:
			{
				ah.heartBeatProcess(msg)
			}
		case <-time.After(15 * time.Second):
			{
				logger.Errorf("hb timeout, nodeId %s", ah.nodeId)
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

func (ah *agentHolder) heartBeatProcess(msg *models.HeartBeatFromAgent) {
	var err error = nil
	func() {
		defer xerr.Recover(&err)
		ah.lastSeqId = msg.SeqId
		ah.lastHBTime = time.Now()
		hbFromServer := ah.packHeartBeatResponse(msg)
		ah.processFinishedTask(ah.nodeStates.BagName, msg)
		xerr.Must0(hbconn.WriteMessage(ah.conn, hbFromServer))
	}()
	if err != nil {
		logger.Error(err)
	}
}

func (ah *agentHolder) recoverWSPanic() {
	if err := recover(); err != nil {
		logger.Error(string(debug.Stack()), err)
		mgrInstance.RemoveNode(ah.nodeId)
	}
}

func (ah *agentHolder) packHeartBeatResponse(
	hbFromAgent *models.HeartBeatFromAgent,
) (hb *models.HeartBeatFromServer) {
	hb = &models.HeartBeatFromServer{
		SeqId: hbFromAgent.SeqId,
	}
	if ah.nodeStates.IsOnGoingStates() {
		bagName, state := ah.nodeStates.GetBagNameWithState()
		logger.Infof("ongoing, state=%s, bagName=%s", state, bagName)
		if state == node_STATES_JOINING {
			hb.JoinBag = &models.JoinBag{
				BagName: bagName,
			}
			if hbFromAgent.Node.BagName == bagName {
				ah.nodeStates.JoinFinished(hbFromAgent.Node.BagName)
			}
		} else if state == node_STATES_FREEING {
			logger.Infof("nodeId %s is freeing from %s", ah.nodeId, bagName)
			hb.FreeNode = &models.FreeNode{}
			if hbFromAgent.Node.BagName == emptyBagName {
				ah.nodeStates.FreeFinished()
			}
		} else {
			logger.Warn("should not in this switch")
		}
	}
	ah.scheduleTasks(hbFromAgent, hb)
	ah.addUploadFilesToHB(hb)
	return hb
}

func (ah *agentHolder) scheduleTasks(
	hbFromAgent *models.HeartBeatFromAgent,
	hb *models.HeartBeatFromServer,
) {
	bagName := ah.nodeStates.GetBagName()
	if bagName == emptyBagName {
		return
	}
	numOfRestResource := ah.maxRunningTasks - len(hbFromAgent.RunningTaskNames)
	for i := 0; i < numOfRestResource; i++ {
		taskName, err := ah.tasksClient.Deque(bagName)
		if err != nil {
			logger.Errorf("deque task from bag %s failed, err %v", bagName, err)
			break
		}
		if hb.ScheduledTaskNames == nil {
			hb.ScheduledTaskNames = make([]string, 0)
		}
		hb.ScheduledTaskNames = append(hb.ScheduledTaskNames, taskName)
	}
	if hb.ScheduledTaskNames != nil {
		logger.Infof("taskNames scheduled to %s is %v", ah.nodeId, hb.ScheduledTaskNames)
		ah.processScheduledTask(bagName, hb.ScheduledTaskNames)
	} else {
		logger.Infof("no task scheduled to %s", ah.nodeId)
	}
}

func (ah *agentHolder) addUploadFilesToHB(
	hb *models.HeartBeatFromServer,
) {
	ah.noUploadFilesMut.Lock()
	defer ah.noUploadFilesMut.Unlock()
	if len(ah.noUploadFiles) > 0 {
		hb.DownloadFiles = ah.noUploadFiles
		logger.Infof("add upload files to hb, %v", hb.DownloadFiles)
		ah.noUploadFiles = make([]models.FileDescription, 0)
	}
}

func (ah *agentHolder) processFinishedTask(bagName string, msg *models.HeartBeatFromAgent) (err error) {
	if bagName == emptyBagName {
		return nil
	}
	if len(msg.FinishedTaskNames) > 0 {
		go comm.
			GetAsyncWorksInstance().
			PersistFinishedTasks(bagName, msg.FinishedTaskNames)
	}
	return
}

func (ah *agentHolder) processScheduledTask(bagName string, scheduledTaskNames []string) (err error) {
	if bagName == emptyBagName {
		return nil
	}
	if len(scheduledTaskNames) > 0 {
		go comm.
			GetAsyncWorksInstance().
			PersistScheduledTasks(bagName, scheduledTaskNames, ah.nodeId)
	}
	return
}

func (ah *agentHolder) persistNodeInfo() (success bool) {
	nodeInfo := &models.NodeInfo{
		NodeId:          ah.nodeId,
		NodeName:        ah.nodeName,
		BagName:         ah.nodeStates.GetBagName(),
		MaxRunningTasks: ah.maxRunningTasks,
	}
	err := db.NewDBOperations().NodeInfos.Create(nodeInfo)
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}

func NewAgent(nodeId string, conn *websocket.Conn) (Agent, error) {
	ah := &agentHolder{
		conn:          conn,
		nodeId:        nodeId,
		hbAgent:       make(chan *models.HeartBeatFromAgent, 1),
		lastSeqId:     -1,
		tasksClient:   taskqueueclient.NewRedisTaskQueueClient(config.Instance().Redis),
		noUploadFiles: make([]models.FileDescription, 0),
	}
	hbStart := &models.HeartBeatStart{}
	err := hbconn.ReadMessage(ah.conn, hbStart)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Infof("new node, nodeId is %s", nodeId)
	ah.nodeStates = newNodeStates()
	ah.nodeName = hbStart.Node.NodeName
	ah.maxRunningTasks = max(hbStart.Node.MaxRunningTasks, 1)
	success := ah.persistNodeInfo()
	response := &models.HeartBeatStartResponse{}
	if success {
		response.Result = models.OK
	} else {
		response.Result = models.ADD_NODE_INFO_FAILED
	}
	if err = hbconn.WriteMessage(ah.conn, response); err != nil {
		return nil, err
	}
	ah.lastHBTime = time.Now()
	go ah.readHBLoop()
	return ah, nil
}
