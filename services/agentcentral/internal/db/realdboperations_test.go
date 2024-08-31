package db_test

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type realDBOperationsTestSuite struct {
	suite.Suite

	dsn string
}

func (s *realDBOperationsTestSuite) TestBagCURD() {
	dbo := db.NewDBOperations()
	n := 10
	bags := make([]*models.Bag, n)
	for i := 0; i < n; i++ {
		bags[i] = &models.Bag{
			BagDisplayName: fmt.Sprintf("test-bag-%d", i),
		}
		dbo.AddBag(bags[i])
	}

	for i := 0; i < n; i++ {
		result := dbo.GetBagByBagName(bags[i].BagName)
		s.Equal(bags[i].BagDisplayName, result.BagDisplayName)
	}

	dbo.DeleteBagByBagName(bags[3].BagName)

	var err error = nil
	func() {
		defer xerr.Recover(&err)
		dbo.GetBagByBagName(bags[3].BagName)
	}()
	s.Equal(err, gorm.ErrRecordNotFound)
}

func (s *realDBOperationsTestSuite) TestAddTask() {
	dbo := db.NewDBOperations()
	n := 10
	for i := 0; i < n; i++ {
		t := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         "test-addtask-bag",
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.AddTask(t)
		s.NotNil(t.TaskName)
		s.True(t.CreateTimeMs != 0)
	}
}

// 这个测试里面测了update和get task的操作
func (s *realDBOperationsTestSuite) TestGetBagEnqueuedTaskNumber() {
	dbo := db.NewDBOperations()
	n := 10
	taskNames := make([]string, n)
	bagName := "test-getbag-enqueued-task-number-bag"
	for i := 0; i < n; i++ {
		task := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         bagName,
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.AddTask(task)
		s.NotNil(task.TaskName)
		s.True(task.CreateTimeMs != 0)
		taskNames[i] = task.TaskName
	}

	for i := 0; i < n; i++ {
		dbo.UpdateTaskOrderId(
			bagName,
			taskNames[i],
			dbo.GetBagEnqueuedTaskNumber(bagName)+1,
		)
	}

	for i := 0; i < n; i++ {
		task := dbo.GetTaskByTaskNameAndBagName(taskNames[i], bagName)
		s.Equal(uint32(i)+1, task.OrderId)
	}
}

func (s *realDBOperationsTestSuite) TestTaskScheduledAndFinishScenario() {
	dbo := db.NewDBOperations()
	n := 10
	taskNames := make([]string, n)
	bagName := "test-task-scheduled-finished-scenario"
	for i := 0; i < n; i++ {
		task := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         bagName,
			ScriptPath:      "/bin/task.sh",
			WorkingDir:      "bin/workingdir",
		}
		dbo.AddTask(task)
		s.NotNil(task.TaskName)
		s.True(task.CreateTimeMs != 0)
		taskNames[i] = task.TaskName
	}
	for i := 0; i < n; i++ {
		dbo.UpdateTaskOrderId(
			bagName,
			taskNames[i],
			dbo.GetBagEnqueuedTaskNumber(bagName)+1,
		)
	}
	nodeId := "test-bag-nodeid"
	scheduledTime := time.Now().UnixMilli()
	finishTime := time.Now().UnixMilli()
	dbo.UpdateTasksScheduledTime(bagName, taskNames, nodeId, scheduledTime)
	dbo.UpdateTasksFinishTime(bagName, taskNames, finishTime)
	taskResults := dbo.GetTaskByMultiFields(map[string]any{
		"bag_name": bagName,
		"node_id":  nodeId,
	})
	s.Len(taskResults, n)
	for _, task := range taskResults {
		s.Equal(finishTime, task.FinishTimeMs)
		s.Equal(scheduledTime, task.ScheduledTimeMs)
	}
}

func (s *realDBOperationsTestSuite) TestListBagNames() {
	dbo := db.NewDBOperations()
	n := 55
	bags := make([]*models.Bag, n)
	for i := 0; i < n; i++ {
		bags[i] = &models.Bag{
			BagDisplayName: fmt.Sprintf("test-bag-%d", i),
		}
		dbo.AddBag(bags[i])
	}

	result := dbo.ListBagNames()
	s.Len(bags, len(result))
	sort.Slice(bags, func(i, j int) bool {
		return bags[i].BagName < bags[j].BagName
	})
	for i := 0; i < n; i++ {
		s.Equal(result[i], bags[i].BagName)
	}
}

func (s *realDBOperationsTestSuite) SetupSuite() {
	tables := []string{"tasks", "bags"}
	s.T().Logf("drop tables %v", tables)
	for _, table := range tables {
		xerr.Must0(db.Instance().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error)
	}
	db.ReInitialWithDSN(s.dsn)
}

func TestRealDBOperationsTestSuite(t *testing.T) {
	var err error
	s := &realDBOperationsTestSuite{}
	func() {
		s.dsn = config.TestConfig().PGSQL_DSN
		defer xerr.Recover(&err)
		db.InitialWithDSN(s.dsn)
	}()
	if err != nil {
		t.Logf("failed to connect db, err is %v", err)
		return
	}
	t.Log("success init! begin to test real db-operations test suite.")
	suite.Run(t, s)
}
