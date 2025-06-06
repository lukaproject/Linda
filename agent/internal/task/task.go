package task

import (
	"Linda/agent/internal/data"
	"Linda/agent/internal/utils"
	"Linda/baselibs/abstractions/xctx"
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
	ExitCode() int
	IsFinished() bool
}

type task struct {
	data.TaskData
	TaskMetrics
	isFinished bool
	exitCode   int
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

func (t *task) Wait() (err error) {
	if t.cmd == nil {
		return ErrCommandIsNil
	}
	err = t.saveExitCode(t.cmd.Wait())
	xctx.Close(t.stdoutFile)
	xctx.Close(t.stderrFile)
	t.isFinished = true
	return err
}

func (t *task) Stop() (err error) {
	err = t.cmd.Process.Kill()
	xctx.Close(t.stdoutFile)
	xctx.Close(t.stderrFile)
	return
}

func (t *task) IsFinished() bool {
	return t.isFinished
}

func (t *task) ExitCode() int {
	if t.IsFinished() {
		return t.exitCode
	}
	return -1
}

// saveExitCode 如果是nil或者是exitError，那么就记录exitCode，
// 原封不动的返回error
func (t *task) saveExitCode(err error) error {
	if err == nil {
		t.exitCode = 0
		return nil
	}
	// 把抛出的exitcode记录一下
	if err1, ok := err.(*exec.ExitError); ok {
		logger.Errorf("exit code is not zero, exit code is %d", err1.ExitCode())
		t.exitCode = err1.ExitCode()
	} else {
		logger.Errorf("not a exit code error, err is %v", err)
	}
	return err
}

func NewTask(
	taskData data.TaskData,
) Task {
	t := &task{
		TaskData: taskData,
	}

	cmds := t.GetCommands(utils.GetDefaultShell())
	t.cmd = exec.Command(cmds[0], cmds[1:]...)
	t.cmd.Dir = taskData.WorkingDir
	return t
}
