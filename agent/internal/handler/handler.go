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
	"time"

	"github.com/lukaproject/xerr"
)

type Handler struct {
	cli     client.IClient
	seqId   int64
	taskMgr task.IMgr
	fileMgr filemanager.Mgr

	logger xlog.Logger
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
}

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
