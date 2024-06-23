package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm"
)

type TasksMgr interface {
	AddTask(task *models.Task)
	GetTask(taskName string) *models.Task
}

type tasksManager struct {
	BagName string
}

func (m *tasksManager) AddTask(task *models.Task) {
	db.NewDBOperations().AddTask(task)
	go comm.GetAsyncWorksInstance().TaskEnque(task.TaskName, task.BagName)
}

func (m *tasksManager) GetTask(taskName string) (task *models.Task) {
	task = db.NewDBOperations().GetTaskByTaskNameAndBagName(taskName, m.BagName)
	return
}
