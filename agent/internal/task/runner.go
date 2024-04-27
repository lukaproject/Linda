package task

import (
	"sync"
	"sync/atomic"
)

type Runner interface {
}

type runner struct {
	FinishedTasksCount atomic.Int32
	RunningTasksCount  atomic.Int32
	ResourceCount      atomic.Int32

	runningTasksMapMut sync.RWMutex
	runningTasksMap    map[int]Task
}

func NewRunner() Runner {
	runner := &runner{
		runningTasksMap: make(map[int]Task),
	}
	return runner
}
