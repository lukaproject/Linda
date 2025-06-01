package task

import (
	"Linda/agent/internal/data"
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type taskTestLinuxSuite struct {
	testBase
}

func (s *taskTestLinuxSuite) TestRunNormalTask_ScriptPath() {
	s.writeStrToTempShellFile(
		`
echo 1
echo 2
echo 3
		`)
	td := data.TaskData{
		Name:         "testtask",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempShellPath(),
		WorkingDir:   s.TempDir(),
		TaskDir:      s.TempDir(),
	}
	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	xerr.Must0(nowtask.Wait())

	expectOutput := "1\n2\n3\n"

	s.T().Log(s.getStrFromFile(path.Join(s.TempDir(), "test.sh")))
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.Equal("", s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestLinuxSuite) TestRunNormalTask_Script() {
	currDir := s.TempDir()
	td := data.TaskData{
		Name:       "testtask",
		Bag:        "testbag",
		Resource:   1,
		Script:     "echo 1",
		WorkingDir: currDir,
		TaskDir:    currDir,
	}

	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	xerr.Must0(nowtask.Wait())

	expectOutput := "1\n"
	s.Equal(expectOutput, s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.Equal("", s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestLinuxSuite) TestRunNormalTask_ScriptOnly1Command() {
	currDir := s.TempDir()
	td := data.TaskData{
		Name:       "testtask",
		Bag:        "testbag",
		Resource:   1,
		Script:     "pwd",
		WorkingDir: currDir,
		TaskDir:    currDir,
	}

	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	xerr.Must0(nowtask.Wait())

	expectOutput := fmt.Sprintf("%s\n", td.WorkingDir)
	s.Equal(expectOutput, s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.Equal("", s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestLinuxSuite) TestRunTaskAndKill() {
	s.writeStrToTempShellFile(
		`
for i in {1..10}
do
	echo $i
	sleep 1
done
		`)
	td := data.TaskData{
		Name:         "testtask2",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempShellPath(),
		WorkingDir:   s.TempDir(),
		TaskDir:      s.TempDir(),
	}
	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	go func() {
		<-time.After(2800 * time.Millisecond)
		xerr.Must0(nowtask.Stop())
	}()
	err := nowtask.Wait()
	s.Equal(SignalMsg(os.Kill), err.Error())

	expectOutput := "1\n2\n3\n"
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.T().Log(s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestLinuxSuite) TestRunTask_ExitCodeNonZero() {
	s.writeStrToTempShellFile(
		`
exit 9
		`)
	td := data.TaskData{
		Name:         "testtask3",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempShellPath(),
		WorkingDir:   s.TempDir(),
		TaskDir:      s.TempDir(),
	}
	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	err := nowtask.Wait()
	s.NotNil(err)
	s.Equal(9, nowtask.ExitCode())
}

func TestTaskTestLinuxSuiteMain(t *testing.T) {
	if runtime.GOOS == "linux" {
		suite.Run(t, new(taskTestLinuxSuite))
	} else {
		t.Skip("only running in linux")
	}
}
