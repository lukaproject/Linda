package suboperations_test

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type tasksTestSuite struct {
	CommonTestSuite
}

func (s *tasksTestSuite) TestCreateTask() {
	dbo := db.NewDBOperations()
	n := 10
	for i := range n {
		t := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         "test-addtask-bag",
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.Tasks.Create(t)
		s.NotNil(t.TaskName)
		s.True(t.CreateTimeMs != 0)
	}
}

// 这个测试里面测了update和get task的操作
func (s *tasksTestSuite) TestGetBagEnqueuedTaskNumber() {
	dbo := db.NewDBOperations()
	n := 10
	taskNames := make([]string, n)
	bagName := "test-getbag-enqueued-task-number-bag"
	for i := range n {
		task := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         bagName,
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.Tasks.Create(task)
		s.NotNil(task.TaskName)
		s.True(task.CreateTimeMs != 0)
		taskNames[i] = task.TaskName
	}

	for i := range n {
		dbo.Tasks.UpdateOrderId(
			bagName,
			taskNames[i],
			dbo.GetBagEnqueuedTaskNumber(bagName)+1,
		)
	}

	for i := range n {
		task := dbo.Tasks.Get(bagName, taskNames[i])
		s.Equal(uint32(i)+1, task.OrderId)
	}
}

func (s *tasksTestSuite) TestTaskScheduledAndFinishScenario() {
	dbo := db.NewDBOperations()
	n := 10
	taskNames := make([]string, n)
	bagName := "test-task-scheduled-finished-scenario"
	for i := range n {
		task := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         bagName,
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.Tasks.Create(task)
		s.NotNil(task.TaskName)
		s.True(task.CreateTimeMs != 0)
		taskNames[i] = task.TaskName
	}
	for i := range n {
		dbo.Tasks.UpdateOrderId(
			bagName,
			taskNames[i],
			dbo.GetBagEnqueuedTaskNumber(bagName)+1,
		)
	}
	nodeId := "test-bag-nodeid"
	scheduledTime := time.Now().UnixMilli()
	finishTime := time.Now().UnixMilli()
	dbo.Tasks.UpdateScheduledTime(bagName, taskNames, nodeId, scheduledTime)
	dbo.Tasks.UpdateFinishedTime(bagName, taskNames, finishTime)
	taskResults := dbo.Tasks.ListByMultiFields(map[string]any{
		"bag_name": bagName,
		"node_id":  nodeId,
	})
	s.Len(taskResults, n)
	for _, task := range taskResults {
		s.Equal(finishTime, task.FinishTimeMs)
		s.Equal(scheduledTime, task.ScheduledTimeMs)
	}
}

func (s *tasksTestSuite) TestListTasks_Prefix_Limit() {
	dbo := db.NewDBOperations()
	testBagName := "test-bag-name"
	for i := range 10 {
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix1_" + fmt.Sprintf("%05d", i),
			BagName:  testBagName,
		}))
	}
	for i := range 10 {
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix2_" + fmt.Sprintf("%05d", i),
			BagName:  testBagName,
		}))
	}
	check := func(prefix string, expectCount, limit int) {
		query := url.Values{}
		query.Set("prefix", prefix)
		query.Set("limit", strconv.Itoa(limit))
		ch := dbo.Tasks.List(testBagName, xerr.Must(abstractions.NewListQueryPacker(query)))
		cnt := 0
		for task := range ch {
			cnt++
			s.True(strings.HasPrefix(task.TaskName, prefix))
			s.T().Log(task.TaskName)
		}
		s.Equal(min(expectCount, limit), cnt)
	}
	check("prefix1", 10, 5)
	check("prefix2", 10, 5)
	check("prefix1", 10, 15)
	check("prefix2", 10, 15)
}

func (s *tasksTestSuite) SetupSuite() {
	s.HealthCheckAndSetup()
	tables := []string{"tasks", "bags", "node_infos"}
	s.T().Logf("drop tables %v", tables)
	for _, table := range tables {
		xerr.Must0(db.Instance().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error)
	}
	db.ReInitialWithDSN(s.dsn)
}

func TestTaskstestSuite(t *testing.T) {
	s := &tasksTestSuite{}
	suite.Run(t, s)
}
