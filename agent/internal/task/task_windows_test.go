package task

import (
	"Linda/agent/internal/data"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type taskTestWindowsSuite struct {
	testBase
}

func (s *taskTestWindowsSuite) writeStrToTempPowerShellFile(content string) {
	s.T().Log(s.tempPowerShellPath())
	s.writeStrToNewFile(s.tempPowerShellPath(), content)
}

func (s *taskTestWindowsSuite) tempPowerShellPath() string {
	return path.Join(s.TempDir(), "test.ps1")
}

func (s *taskTestWindowsSuite) TestRunNormalTask() {
	s.writeStrToTempPowerShellFile(
		`
    Write-Host 1
	Write-Host 2
	Write-Host 3
		`)
	td := data.TaskData{
		Name:         "testtask",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempPowerShellPath(),
		WorkingDir:   s.TempDir(),
		TaskDir:      s.TempDir(),
	}
	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	xerr.Must0(nowtask.Wait())

	expectOutput := "1\r\n2\r\n3\r\n"

	s.T().Log(s.getStrFromFile(s.tempPowerShellPath()))
	s.T().Log(s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.Equal("", s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestWindowsSuite) TestRunTaskAndKill() {
	s.writeStrToTempPowerShellFile(
		`
for ($i = 1; $i -le 10; $i++) {
    Write-Host $i
    Start-Sleep -Seconds 1
}
		`)
	td := data.TaskData{
		Name:         "testtask2",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempPowerShellPath(),
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
	s.Equal("exit status 1", err.Error())

	expectOutput := "1\r\n2\r\n3\r\n"
	s.Equal(
		expectOutput,
		s.getStrFromFile(path.Join(td.TaskDir, ConstStdOutFile)))
	s.T().Log(s.getStrFromFile(path.Join(td.TaskDir, ConstStdErrFile)))
}

func (s *taskTestWindowsSuite) TestRunTask_ExitCodeNonZero() {
	s.writeStrToTempPowerShellFile(
		`
Exit 5
		`)
	td := data.TaskData{
		Name:         "testtask3",
		Bag:          "testbag",
		Resource:     1,
		PathToScript: s.tempPowerShellPath(),
		WorkingDir:   s.TempDir(),
		TaskDir:      s.TempDir(),
	}
	nowtask := NewTask(td)
	xerr.Must0(nowtask.Start())
	err := nowtask.Wait()
	s.NotNil(err)
	s.Equal(5, nowtask.ExitCode())
}

func TestTaskTestWindowsSuiteMain(t *testing.T) {
	if runtime.GOOS == "windows" {
		suite.Run(t, new(taskTestWindowsSuite))
	} else {
		t.Skip("only running in windows")
	}
}
