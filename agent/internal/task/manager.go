package task

import (
	"Linda/agent/internal/config"
	"Linda/protocol/models"
	"fmt"
	"net/http"

	"github.com/lukaproject/xerr"
)

type Mgr struct {
	taskRunner *runner
}

func (m *Mgr) AddTask(taskName string) {
	xerr.Must0(m.taskRunner.AddTask(NewTask(m.fetchTaskDataByTaskName(taskName))))
}

func (m *Mgr) PopFinishedTasks() (finishedTaskNames []string) {
	count := len(m.taskRunner.FinishedTaskChan)
	finishedTaskNames = make([]string, 0, count)
	for ; count > 0; count-- {
		t, ok := <-m.taskRunner.FinishedTaskChan
		if ok {
			finishedTaskNames = append(finishedTaskNames, t)
		} else {
			break
		}
	}
	return
}

func (m *Mgr) fetchTaskDataByTaskName(taskName string) (data TaskData) {
	resp := xerr.Must(http.Get(
		fmt.Sprintf(
			"http://%s/bags/%s/tasks/%s",
			config.Instance().AgentCentralUrlPrefix, config.Instance().BagName, taskName)))
	t := &models.Task{}
	models.ReadJSON(resp.Body, t)
	data.FromTaskModel(t)
	return
}

func NewMgr() *Mgr {
	runnerIniter := RunnerIniter{
		MaxResourceCount: 1,
	}
	mgr := &Mgr{
		taskRunner: NewRunner(runnerIniter),
	}
	mgr.taskRunner.initial()
	return mgr
}
