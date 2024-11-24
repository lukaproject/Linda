package handler

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/task"
	"Linda/protocol/models"
	"fmt"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	cli     client.IClient
	seqId   int64
	taskMgr task.IMgr
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
	if len(resp.ScheduledTaskNames) != 0 {
		go func() {
			for _, taskName := range resp.ScheduledTaskNames {
				h.taskMgr.AddTask(taskName)
			}
		}()
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
		req.FinishedTaskNames = h.taskMgr.PopFinishedTasks()
	}
	return
}

func (h *Handler) joinBag(joinBag *models.JoinBag) {
	if joinBag == nil {
		return
	}
	nowBagName := data.GetData(data.Instance().NodeData, true).BagName
	if nowBagName != joinBag.BagName {
		logrus.Warnf(
			"join bag failed, current bag %s not equal to comming bag %s",
			nowBagName, joinBag.BagName)
	}

	data.Instance().UpdateNodeData(
		func(nd *data.NodeData) *data.NodeData {
			nd.BagName = joinBag.BagName
			return nd
		})
	logrus.Infof("join bag %s", joinBag.BagName)
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
	logrus.Info("free node")
}

func NewHandler(c *config.Config) *Handler {
	if c == nil {
		c = config.Instance()
	}
	return NewHandlerWithCliAndTaskMgr(
		xerr.Must(client.New(c.AgentHeartBeatUrl())),
		task.NewMgr())
}

func NewHandlerWithCliAndTaskMgr(
	cli client.IClient,
	taskMgr task.IMgr,
) *Handler {
	return &Handler{
		seqId:   0,
		taskMgr: taskMgr,
		cli:     cli,
	}
}
