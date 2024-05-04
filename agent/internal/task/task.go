package task

import (
	"Linda/agent/internal/utils"
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/lukaproject/xerr"
)

const (
	ConstStdOutFile = "stdout.txt"
	ConstStdErrFile = "stderr.txt"
)

type TaskData struct {
	Name         string
	Bag          string
	Resource     int
	PathToScript string

	// script running dir
	WorkingDir string
	// task data located dir, such as stdout / stderr or others.
	TaskDir string
}

type Task interface {
	GetName() string
	GetBag() string
	GetResource() int
	GetWorkingDir() string
	GetData() TaskData
	Start() error
	Stop() error
	Wait() error
	IsFinished() bool
}

type task struct {
	TaskData
	isFinished bool
	cmd        *exec.Cmd

	stdoutFile *os.File
	stderrFile *os.File
}

func (t *task) GetName() string {
	return t.Name
}

func (t *task) GetBag() string {
	return t.Bag
}

func (t *task) GetResource() int {
	return t.Resource
}

func (t *task) GetWorkingDir() string {
	return t.WorkingDir
}

func (t *task) GetData() TaskData {
	return t.TaskData
}

func (t *task) Start() (err error) {
	func() {
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		t.stdoutFile = xerr.Must(os.Create(path.Join(t.TaskDir, ConstStdOutFile)))
		t.cmd.Stdout = t.stdoutFile

		t.stderrFile = xerr.Must(os.Create(path.Join(t.TaskDir, ConstStdErrFile)))
		t.cmd.Stderr = t.stderrFile

		xerr.Must0(t.cmd.Start())
	}()
	return
}

func (t *task) Wait() error {
	if t.cmd == nil {
		return errors.New("cmd is nil")
	}
	return t.cmd.Wait()
}

func (t *task) Stop() (err error) {
	func() {
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		xerr.Must0(t.cmd.Process.Kill())
		t.isFinished = true
		xerr.Must0(t.stdoutFile.Close())
		xerr.Must0(t.stderrFile.Close())
	}()
	return
}

func (t *task) IsFinished() bool {
	return t.isFinished
}

func NewTask(
	taskData TaskData,
) Task {
	t := &task{
		TaskData: taskData,
	}

	t.cmd = exec.Command(utils.GetDefaultShell(), t.PathToScript)
	t.cmd.Dir = taskData.WorkingDir
	return t
}
