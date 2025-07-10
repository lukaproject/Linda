package handler

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/filemanager"
	"Linda/agent/internal/task"
	"Linda/baselibs/abstractions/xlog"
	"Linda/protocol/models"
	"fmt"
	"sync"
	"time"

	"github.com/lukaproject/xerr"
)

type Handler struct {
	cli     client.IClient
	seqId   int64
	taskMgr task.IMgr
	fileMgr filemanager.Mgr

	logger xlog.Logger

	// Add response queues
	fileListResponses []models.FileListResponse
	fileGetResponses  []models.FileGetResponse
	responsesMutex    sync.Mutex
}

func (h *Handler) Start() {
	h.startHeartBeat()
	go h.keepAlive()
}

func (h *Handler) Run() {
	h.startHeartBeat()
	h.keepAlive()
}

func (h *Handler) startHeartBeat() {
	resp := xerr.Must(
		h.cli.HeartBeatStart(&models.HeartBeatStart{
			Node: models.NodeInfo{
				MaxRunningTasks: 1,
				NodeName:        data.Instance().NodeData.NodeName,
			},
		}))
	if resp.Result != models.OK {
		panic(fmt.Sprintf("start HB failed, %s", resp.Result))
	}
}

func (h *Handler) keepAlive() {
	for {
		resp := xerr.Must(h.cli.HeartBeat(h.packReq()))
		h.joinBag(resp.JoinBag)
		h.freeNode(resp.FreeNode)
		h.unPackResp(resp)
		h.seqId++
		<-time.After(config.Instance().HeartbeatPeriod())
	}
}

// handler the heartbeat from server
// it will handle the scheduled tasks, download files, file list requests, and file get requests
func (h *Handler) unPackResp(resp *models.HeartBeatFromServer) {
	if len(resp.ScheduledTasks) != 0 {
		go func() {
			for _, taskInfo := range resp.ScheduledTasks {
				h.taskMgr.AddTask(task.AddTaskInput{
					Name:      taskInfo.Name,
					AccessKey: taskInfo.AccessKey,
				})
			}
		}()
	}
	if len(resp.DownloadFiles) != 0 {
		go downloadFiles(h.logger, h.fileMgr, resp.DownloadFiles)
	}

	// Handle file list requests
	if len(resp.FileListRequests) != 0 {
		go listFileInfos(h.logger, h.fileMgr, resp.FileListRequests)
	}

	// Handle file get requests
	if len(resp.FileGetRequests) != 0 {
		go getFiles(h.logger, h.fileMgr, resp.FileGetRequests)
	}
}

// Pack the heartbeat request to send to the server
// It includes the sequence ID, node information, finished tasks, and file operation responses.
func (h *Handler) packReq() (req *models.HeartBeatFromAgent) {
	req = &models.HeartBeatFromAgent{
		SeqId: h.seqId,
	}
	nodeData := data.GetData(data.Instance().NodeData, true)
	req.Node = models.NodeInfo{
		BagName: nodeData.BagName,
	}
	if h.taskMgr != nil {
		finishedTaskResult := h.taskMgr.PopFinishedTasks()
		req.FinishedTasks = make([]models.FinishedTaskResult, 0, len(finishedTaskResult))
		for _, result := range finishedTaskResult {
			req.FinishedTasks = append(req.FinishedTasks,
				models.FinishedTaskResult{
					Name:     result.Name,
					ExitCode: result.ExitCode,
				})
		}
	}

	// Add file operation responses
	h.responsesMutex.Lock()
	req.FileListResponses = make([]models.FileListResponse, len(h.fileListResponses))
	copy(req.FileListResponses, h.fileListResponses)
	h.fileListResponses = h.fileListResponses[:0]

	req.FileGetResponses = make([]models.FileGetResponse, len(h.fileGetResponses))
	copy(req.FileGetResponses, h.fileGetResponses)
	h.fileGetResponses = h.fileGetResponses[:0]
	h.responsesMutex.Unlock()

	return
}

func (h *Handler) joinBag(joinBag *models.JoinBag) {
	if joinBag == nil {
		return
	}
	nowBagName := data.GetData(data.Instance().NodeData, true).BagName
	if nowBagName != joinBag.BagName && nowBagName != "" {
		h.logger.Warnf(
			"join bag failed, current bag %s not equal to comming bag %s",
			nowBagName, joinBag.BagName)
	}

	data.Instance().UpdateNodeData(
		func(nd *data.NodeData) *data.NodeData {
			nd.BagName = joinBag.BagName
			return nd
		})
	h.logger.Infof("join bag %s", joinBag.BagName)
}

func (h *Handler) freeNode(freeNode *models.FreeNode) {
	if freeNode == nil {
		return
	}
	data.Instance().UpdateNodeData(
		func(nd *data.NodeData) *data.NodeData {
			nd.BagName = ""
			return nd
		})
	h.logger.Info("free node")
}

func NewHandler(c *config.Config) *Handler {
	if c == nil {
		c = config.Instance()
	}
	return NewHandlerWithParameters(
		xerr.Must(client.New(c.AgentHeartBeatUrl())),
		task.NewMgr(),
		filemanager.NewMgr())
}

func NewHandlerWithParameters(
	cli client.IClient,
	taskMgr task.IMgr,
	fileMgr filemanager.Mgr,
) *Handler {
	return &Handler{
		seqId:   0,
		taskMgr: taskMgr,
		fileMgr: fileMgr,
		cli:     cli,
		logger:  xlog.NewForPackage(),
	}
}
