package handler

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/task"
	"Linda/protocol/models"
	"fmt"
	"time"

	"github.com/lukaproject/xerr"
)

type Handler struct {
	cli     client.IClient
	seqId   int64
	bagName string
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
				BagName:         h.bagName,
			},
		}))
	if resp.Result != models.OK {
		panic(fmt.Sprintf("start HB failed, %s", resp.Result))
	}
}

func (h *Handler) keepAlive() {
	for {
		resp := xerr.Must(h.cli.HeartBeat(h.packReq()))
		h.unPackResp(resp)
		h.seqId++
		<-time.After(2 * time.Second)
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
		SeqId:             h.seqId,
		FinishedTaskNames: h.taskMgr.PopFinishedTasks(),
	}
	return
}

func NewHandler(c *config.Config) *Handler {
	h := &Handler{
		seqId:   0,
		bagName: c.BagName,
		taskMgr: task.NewMgr(),
	}
	h.cli = xerr.Must(client.New(c.AgentCentralUrl()))
	return h
}
