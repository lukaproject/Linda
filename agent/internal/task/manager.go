package task

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/baselibs/abstractions/xos"
	"Linda/protocol/models"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/lukaproject/xerr"
)

type AddTaskInput struct {
	Name      string
	AccessKey string
}

type IMgr interface {
	AddTask(addTaskInput AddTaskInput)
	PopFinishedTasks() (finishedTaskNames []string)
}

type Mgr struct {
	taskRunner *runner
}

func (m *Mgr) AddTask(addTaskInput AddTaskInput) {
	taskData, err := m.fetchTaskDataByTaskName(addTaskInput.Name, addTaskInput.AccessKey)
	if err != nil {
		logger.Error(err)
		return
	}
	// need to pack task directory.
	taskData.TaskDir = m.newTaskDir(&taskData)
	err = m.taskRunner.AddTask(NewTask(taskData))
	if err != nil {
		logger.Error(err)
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

func (m *Mgr) fetchTaskDataByTaskName(taskName string, accessKey string) (taskData data.TaskData, err error) {
	func() {
		defer xerr.Recover(&err)
		bagName := data.Instance().NodeData.BagName
		taskUrl := m.getTaskUrl(bagName, taskName)
		// agent访问 task 必须 accessKey
		resp := xerr.Must(http.PostForm(
			taskUrl,
			url.Values{
				"accessKey": []string{accessKey},
			}))
		if resp.StatusCode != http.StatusOK {
			logger.Errorf(
				"can not to fetch task body, task name %s, bag name %s, status %s",
				taskName, bagName, resp.Status)
			return
		}
		t := &models.Task{}
		models.ReadJSON(resp.Body, t)
		taskData.FromTaskModel(t)
	}()
	return
}

func (m *Mgr) getTaskUrl(bagName, taskName string) string {
	return strings.Join(
		[]string{
			config.Instance().AgentAPIUrl("http"),
			"agent/innercall/bags",
			bagName,
			"tasks",
			taskName,
		}, "/")
}

func (m *Mgr) newTaskDir(taskData *data.TaskData) string {
	tasksBaseDir := config.Instance().TasksDir
	taskDir := path.Join(tasksBaseDir, taskData.Bag, taskData.Name)
	if !xos.PathExists(taskDir) {
		xos.MkdirAll(taskDir, os.ModePerm)
	}
	return taskDir
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
