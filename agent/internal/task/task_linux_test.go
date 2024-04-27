package task_test

import (
	"Linda/agent/internal/task"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type taskTestLinuxSuite struct {
	taskBaseSuite
}

func (s *taskTestLinuxSuite) TestRunNormalTask() {
	s.writeStrToTempShellFile(
		`
echo 1
echo 2
echo 3
		`)
	nowtask := task.NewTask(s.tempShellPath(), s.tempTestDir(), 1)
	xerr.Must0(nowtask.Start())
	xerr.Must0(nowtask.Wait())

	expectOutput := "1\n2\n3\n"

	s.T().Log(s.getStrFromFile(path.Join(s.tempTestDir(), "test.sh")))
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(s.tempTestDir(), task.ConstStdOutFile)))
	s.Equal("", s.getStrFromFile(path.Join(s.tempTestDir(), task.ConstStdErrFile)))
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
	nowtask := task.NewTask(s.tempShellPath(), s.tempTestDir(), 1)
	xerr.Must0(nowtask.Start())
	go func() {
		<-time.After(2800 * time.Millisecond)
		xerr.Must0(nowtask.Stop())
	}()
	err := nowtask.Wait()
	s.Equal(task.SignalMsg(os.Kill), err.Error())

	expectOutput := "1\n2\n3\n"
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(s.tempTestDir(), task.ConstStdOutFile)))
	s.T().Log(s.getStrFromFile(path.Join(s.tempTestDir(), task.ConstStdErrFile)))
}

func TestTaskTestLinuxSuiteMain(t *testing.T) {
	if runtime.GOOS == "linux" {
		suite.Run(t, new(taskTestLinuxSuite))
	}
}
