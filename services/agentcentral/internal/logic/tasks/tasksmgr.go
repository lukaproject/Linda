package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
)

type TasksMgr interface {
	// 新增一个task
	AddTask(task *models.Task)

	// 读取task内容
	GetTask(taskName string) *models.Task
}

type tasksManager struct {
	BagName string
	queCli  taskqueueclient.QuesManageClient
}

func (m *tasksManager) AddTask(task *models.Task) {
	db.NewDBOperations().Tasks.Create(task)
	go comm.GetAsyncWorksInstance().TaskEnque(task.TaskName, task.BagName, uint16(task.Priority))
}

func (m *tasksManager) GetTask(taskName string) (task *models.Task) {
	task = db.NewDBOperations().Tasks.Get(m.BagName, taskName)
	return
}

func NewTasksMgr(bagName string, queCli taskqueueclient.QuesManageClient) TasksMgr {
	return &tasksManager{
		BagName: bagName,
		queCli:  queCli,
	}
}
