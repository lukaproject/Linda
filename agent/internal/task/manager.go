package task

import (
	"Linda/agent/internal/config"
	"Linda/protocol/models"
	"net/http"
	"strings"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type IMgr interface {
	AddTask(taskName string)
	PopFinishedTasks() (finishedTaskNames []string)
}

type Mgr struct {
	taskRunner *runner
}

func (m *Mgr) AddTask(taskName string) {
	taskData, err := m.fetchTaskDataByTaskName(taskName)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = m.taskRunner.AddTask(NewTask(taskData))
	if err != nil {
		logrus.Error(err)
		return
	}
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

func (m *Mgr) fetchTaskDataByTaskName(taskName string) (data TaskData, err error) {
	func() {
		defer xerr.Recover(&err)
		taskUrl := m.getTaskUrl(taskName)
		resp := xerr.Must(http.Get(taskUrl))
		if resp.StatusCode != http.StatusOK {
			logrus.Errorf(
				"can not to fetch task body, task name %s, bag name %s, status %s",
				taskName, config.Instance().BagName, resp.Status)
			return
		}
		t := &models.Task{}
		models.ReadJSON(resp.Body, t)
		data.FromTaskModel(t)
	}()
	return
}

func (m *Mgr) getTaskUrl(taskName string) string {
	return strings.Join(
		[]string{
			config.Instance().AgentAPIUrl("http"),
			"bags",
			config.Instance().BagName,
			"tasks",
			taskName,
		}, "/")
}

func NewMgr() IMgr {
	runnerIniter := RunnerIniter{
		MaxResourceCount: 1,
	}
	mgr := &Mgr{
		taskRunner: NewRunner(runnerIniter),
	}
	mgr.taskRunner.initial()
	return mgr
}
