package data

import "sync"

const BucketRunningTaskData = "running_tasks"

type RunningTasksContainer struct {
	mut   sync.RWMutex
	tasks map[string]any
}

func NewRunningTasksContainer() *RunningTasksContainer {
	return &RunningTasksContainer{
		tasks: make(map[string]any),
	}
}

func (c *RunningTasksContainer) Iter() []string {
	c.mut.RLock()
	defer c.mut.RUnlock()
	tasksList := make([]string, len(c.tasks))
	return tasksList
}

func (c *RunningTasksContainer) ToRunning(taskId string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.tasks[taskId] = new(any)
}

func (c *RunningTasksContainer) ToFinished(taskId string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	delete(c.tasks, taskId)
}
