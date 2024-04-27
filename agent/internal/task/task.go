package task

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/lukaproject/xerr"
)

const (
	ConstStdOutFile = "stdout.txt"
	ConstStdErrFile = "stderr.txt"
)

type Task interface {
	Start() error
	Stop() error
	Wait() error

	IsFinished() bool
}

type task struct {
	Resource int

	pathToScript string
	workingDir   string
	isFinished   bool
	cmd          *exec.Cmd

	stdoutFile *os.File
	stderrFile *os.File
}

func (t *task) Start() (err error) {
	func() {
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		t.stdoutFile = xerr.Must(os.Create(path.Join(t.workingDir, ConstStdOutFile)))
		t.cmd.Stdout = t.stdoutFile

		t.stderrFile = xerr.Must(os.Create(path.Join(t.workingDir, ConstStdErrFile)))
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
	pathToScript, workingDir string,
	resource int,
) Task {
	var shell string

	switch runtime.GOOS {
	case "windows":
		shell = "pwsh"
	case "linux":
		shell = "/bin/bash"
	default:
		panic(errors.New("unsupported OS"))
	}
	t := &task{
		Resource:     resource,
		pathToScript: pathToScript,
		workingDir:   workingDir,
	}

	t.cmd = exec.Command(shell, t.pathToScript)
	t.cmd.Dir = workingDir
	return t
}
