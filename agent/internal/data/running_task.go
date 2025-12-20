package data

import (
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xctx"
	"sync"

	"github.com/lukaproject/xerr"
)

const BucketRunningTaskData = "running_tasks"

// 这里是存所有的正在执行的任务id，key为string
// 存储的bucket name为running_tasks

var rtc *RunningTasksContainer = nil

func InitialRunningTasksContainer() {
	rtc = NewRunningTasksContainer()
}

func GetRunningTasksContainerInstance() *RunningTasksContainer {
	return rtc
}

type RunningTasksContainer struct {
	mut       sync.RWMutex
	taskIds   map[string]any
	persistor *localdb.Persistor[*StringType, *StringType]
}

func NewRunningTasksContainer() *RunningTasksContainer {
	return &RunningTasksContainer{
		taskIds:   make(map[string]any),
		persistor: xerr.Must(localdb.GetPersistor[*StringType, *StringType](BucketRunningTaskData)),
	}
}

func (c *RunningTasksContainer) Init() {
	xctx.NewLocker(&c.mut).Run(
		func() {
			keys := c.persistor.GetKeys()
			for _, k := range keys {
				c.taskIds[k.Key] = k.Key
			}
		})
}

func (c *RunningTasksContainer) ListAll() []string {
	c.mut.RLock()
	defer c.mut.RUnlock()
	tasksList := make([]string, 0, len(c.taskIds))
	for k := range c.taskIds {
		tasksList = append(tasksList, k)
	}
	return tasksList
}

func (c *RunningTasksContainer) ToRunning(taskId string) error {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.taskIds[taskId] = new(any)
	logger.Infof("task to running %s", taskId)
	return c.persistor.Set(NewString(taskId), NewString(taskId))
}

func (c *RunningTasksContainer) ToFinished(taskId string) error {
	c.mut.Lock()
	defer c.mut.Unlock()
	delete(c.taskIds, taskId)
	logger.Infof("task to finish %s", taskId)
	return c.persistor.Delete(NewString(taskId))
}
