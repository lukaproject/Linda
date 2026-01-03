package data_test

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/testcommon/testenv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type runningTasksTestSuite struct {
	testenv.TestBase
}

func (s *runningTasksTestSuite) Test_NormalRunningTasksTransactions() {
	config.SetInstance(&config.Config{
		LocalDBDir: s.TempDir(),
	})
	localdb.Initial()

	rtc := data.NewRunningTasksContainer()

	rtc.Init()
	taskIds := []string{"xx1", "xx2", "xx3"}
	for _, taskId := range taskIds {
		s.NoError(rtc.ToRunning(taskId))
	}
	actualTaskIds := rtc.ListAll()
	s.ElementsMatch(taskIds, actualTaskIds)

	s.NoError(rtc.ToFinished("xx2"))
	actualTaskIds = rtc.ListAll()
	s.ElementsMatch([]string{"xx3", "xx1"}, actualTaskIds)
}

func TestRunningTasksTestSuiteMain(t *testing.T) {
	suite.Run(t, new(runningTasksTestSuite))
}
