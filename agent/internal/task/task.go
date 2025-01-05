package task

import (
	"Linda/agent/internal/data"
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

type TaskMetrics struct{}

type Task interface {
	GetName() string
	GetBag() string
	GetResource() int
	GetWorkingDir() string
	GetData() data.TaskData
	Start() error
	Stop() error
	Wait() error
	IsFinished() bool
}

type task struct {
	data.TaskData
	TaskMetrics
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

func (t *task) GetData() data.TaskData {
	return t.TaskData
}

func (t *task) Start() (err error) {
	func() {
		defer xerr.Recover(&err)

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
		defer xerr.Recover(&err)
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
	taskData data.TaskData,
) Task {
	t := &task{
		TaskData: taskData,
	}

	t.cmd = exec.Command(utils.GetDefaultShell(), t.PathToScript)
	t.cmd.Dir = taskData.WorkingDir
	return t
}
