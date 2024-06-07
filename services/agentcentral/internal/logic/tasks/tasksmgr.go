package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"

	"github.com/lukaproject/xerr"
)

type TasksMgr interface {
	AddTask(task *models.Task)
	GetTask(taskName string) *models.Task
}

type tasksManager struct {
	BagName string
}

func (m *tasksManager) AddTask(task *models.Task) {
	dbi := db.Instance()
	xerr.Must0(dbi.Save(task).Error)
}

func (m *tasksManager) GetTask(taskName string) (task *models.Task) {
	dbi := db.Instance()
	task = &models.Task{
		TaskName: taskName,
	}
	xerr.Must0(dbi.First(task).Error)
	return
}
