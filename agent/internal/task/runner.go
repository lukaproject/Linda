package task

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

type runner struct {
	FinishedTasksCount atomic.Int32
	ResourceCount      atomic.Int32
	MaxResourceCount   int

	runningTasksMapMut sync.RWMutex
	// 任务名字与任务的映射
	runningTasksMap map[string]Task
}

func NewRunner(initer RunnerIniter) *runner {
	runner := &runner{
		runningTasksMap:  make(map[string]Task),
		MaxResourceCount: initer.MaxResourceCount,
	}
	runner.initial()
	return runner
}

func (r *runner) CountRunningTasks() int {
	r.runningTasksMapMut.RLock()
	defer r.runningTasksMapMut.RUnlock()
	return len(r.runningTasksMap)
}

func (r *runner) AddTask(t Task) (err error) {
	if int(r.ResourceCount.Load())+t.GetResource() > r.MaxResourceCount {
		return errors.New("have no enough resource, retry later")
	}
	r.runningTasksMapMut.Lock()
	defer r.runningTasksMapMut.Unlock()
	if _, ok := r.runningTasksMap[t.GetName()]; ok {
		return errors.New("task exist now")
	}
	r.runningTasksMap[t.GetName()] = t
	if err = t.Start(); err != nil {
		return err
	}
	r.ResourceCount.Add(int32(t.GetResource()))
	go r.taskRunningCallback(t)
	return
}

func (r *runner) taskRunningCallback(t Task) {
	if err := t.Wait(); err != nil {
		logrus.Error(err)
	}
	func() {
		r.runningTasksMapMut.Lock()
		defer r.runningTasksMapMut.Unlock()
		delete(r.runningTasksMap, t.GetName())
	}()
	r.ResourceCount.Add(int32(-t.GetResource()))
	r.FinishedTasksCount.Add(1)
}

func (r *runner) initial() {
}

type RunnerIniter struct {
	MaxResourceCount int
}

func DefaultRunnerIniter() RunnerIniter {
	initer := RunnerIniter{
		MaxResourceCount: runtime.NumCPU(),
	}
	return initer
}
