package task

import (
	"fmt"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type runnerTestSuite struct {
	testBase
}

func (s *runnerTestSuite) TestRunnerNormalTask() {
	initer := DefaultRunnerIniter()
	n := 10
	initer.MaxResourceCount = n

	runner := NewRunner(initer)

	tasks := make([]Task, 0)

	for i := 0; i < n; i++ {
		taskName := fmt.Sprintf("%d", i)
		taskBag := "testbag"
		t := s.createNewTestTask(taskName, taskBag, fmt.Sprintf("echo %d", i), 1)
		s.Nil(runner.AddTask(t))
		tasks = append(tasks, t)
	}
	s.WaitUntilWithTimeout(func() bool {
		return runner.CountRunningTasks() == 0
	}, time.Minute)

	for _, t := range tasks {
		outputFilePath := path.Join(s.taskRequireDir(t), ConstStdOutFile)
		output := s.getStrFromFile(outputFilePath)
		s.Equal(t.GetName()+"\n", output)
	}
}

func TestRunnerTestSuiteMain(t *testing.T) {
	if runtime.GOOS == "linux" {
		suite.Run(t, new(runnerTestSuite))
	} else {
		t.Skip("only running in linux")
	}
}
