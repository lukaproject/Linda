package task

import (
	"errors"
	"os"
)

func SignalMsg(s os.Signal) string {
	return "signal: " + s.String()
}

var (
	ErrNoEnoughResource = errors.New("have no enough resource, retry later")
	ErrTaskExist        = errors.New("task exist now")
	ErrTaskNotExist     = errors.New("no such task in runner")
	ErrCommandIsNil     = errors.New("cmd is nil")
)
