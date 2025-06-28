package tasks_test

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/db/dbtestcommon"
	"Linda/services/agentcentral/internal/logic/comm"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"Linda/services/agentcentral/internal/logic/tasks"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/suite"
)

type tasksMgrTestSuite struct {
	dbtestcommon.CommonTestSuite

	rds *miniredis.Miniredis
}

func (s *tasksMgrTestSuite) TestTaskMgr_AddTask() {
	aw := comm.GetAsyncWorksInstance()
	testBagName := "testBag"
	testTaskName := "testTask"
	queCli := taskqueueclient.NewRedisQuesManageClient(config.Instance().Redis)
	aw.AddBag(testBagName)
	tasksMgr := tasks.NewTasksMgr(testBagName, queCli)

	tasksMgr.AddTask(&models.Task{
		TaskName: testTaskName,
		BagName:  testBagName,
		Script:   "echo 1",
		Priority: 100,
	})

	<-time.After(2 * time.Second)
	task := tasksMgr.GetTask(testTaskName)
	s.Equal(int16(100), task.Priority)
	s.Equal(uint32(1), task.OrderId)
}

func (s *tasksMgrTestSuite) SetupTest() {
	s.HealthCheckAndSetup()
	s.DropTables()
	s.rds = miniredis.RunT(s.T())
	config.Instance().Redis = &config.RedisConfig{
		Addrs: []string{s.rds.Addr()},
	}
	db.ReInitialWithDSN(s.DSN)
	comm.InitAsyncWorksInstance()
}

func TestTaskMgrMain(t *testing.T) {
	suite.Run(t, new(tasksMgrTestSuite))
}
