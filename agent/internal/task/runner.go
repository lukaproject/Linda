package task

import (
	"Linda/baselibs/abstractions/xos"
	"io/fs"
	"runtime"
	"sync"
	"sync/atomic"
)

type runner struct {
	FinishedTasksCount atomic.Int32
	// 完成的任务会被塞入这个channel中
	FinishedTaskChan chan string
	ResourceCount    atomic.Int32
	MaxResourceCount int

	runningTasksMapMut sync.RWMutex
	// 任务名字与任务的映射
	runningTasksMap map[string]Task
}

func NewRunner(initer RunnerIniter) *runner {
	runner := &runner{
		runningTasksMap:  make(map[string]Task),
		MaxResourceCount: initer.MaxResourceCount,
		FinishedTaskChan: make(chan string, 32),
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
	r.createWorkingDir(t)
	if int(r.ResourceCount.Load())+t.GetResource() > r.MaxResourceCount {
		return ErrNoEnoughResource
	}
	r.runningTasksMapMut.Lock()
	defer r.runningTasksMapMut.Unlock()
	if _, ok := r.runningTasksMap[t.GetName()]; ok {
		return ErrTaskExist
	}
	r.runningTasksMap[t.GetName()] = t
	if err = t.Start(); err != nil {
		logger.Errorf("error=%v, taskName=%s", err, t.GetName())
		return err
	}
	r.ResourceCount.Add(int32(t.GetResource()))
	go r.taskRunningCallback(t)
	return
}

func (r *runner) GetTask(taskName string) (t Task, err error) {
	r.runningTasksMapMut.RLock()
	defer r.runningTasksMapMut.RUnlock()
	var ok bool
	t, ok = r.runningTasksMap[taskName]
	if !ok {
		return nil, ErrTaskNotExist
	}
	return t, nil
}

func (r *runner) taskRunningCallback(t Task) {
	if err := t.Wait(); err != nil {
		logger.Error(err)
	}
	func() {
		r.runningTasksMapMut.Lock()
		defer r.runningTasksMapMut.Unlock()
		delete(r.runningTasksMap, t.GetName())
	}()
	r.ResourceCount.Add(int32(-t.GetResource()))
	r.FinishedTasksCount.Add(1)
}

func (r *runner) createWorkingDir(t Task) (err error) {
	xos.MkdirAll(t.GetWorkingDir(), fs.ModePerm)
	return
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
