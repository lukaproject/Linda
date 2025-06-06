package suboperations_test

import (
	"Linda/baselibs/abstractions"
	"Linda/baselibs/testcommon/gen"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/db/dbtestcommon"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ecodeclub/ekit/slice"
	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type tasksTestSuite struct {
	dbtestcommon.CommonTestSuite
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

func (s *tasksTestSuite) TestCreateTask_Script() {
	dbo := db.NewDBOperations()
	n := 10
	for i := range n {
		t := &models.Task{
			TaskDisplayName: fmt.Sprintf("testtask-%d", i),
			BagName:         "test-addtask-bag",
			Script:          "echo 1",
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
	tasks := make([]models.FinishedTaskResult, n)
	mapTasksExitCode := make(map[string]int32)
	accessKeys := make([]string, n)
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
		tasks[i].Name = task.TaskName
		tasks[i].ExitCode = int32(rand.Int() % 10)
		mapTasksExitCode[tasks[i].Name] = tasks[i].ExitCode
		var err error
		accessKeys[i], err = gen.StrGenerate(gen.CharsetLowerCase, 5, 10)
		s.Nil(err)
	}
	for i := range n {
		dbo.Tasks.UpdateOrderId(
			bagName,
			tasks[i].Name,
			dbo.GetBagEnqueuedTaskNumber(bagName)+1,
		)
	}
	nodeId := "test-bag-nodeid"
	scheduledTime := time.Now().UnixMilli()
	finishTime := time.Now().UnixMilli()
	dbo.Tasks.UpdateScheduledTime(
		bagName,
		slice.Map(tasks,
			func(_ int, taskResult models.FinishedTaskResult) string {
				return taskResult.Name
			}),
		accessKeys,
		nodeId,
		scheduledTime)
	dbo.Tasks.PersistFinishedTasks(bagName, tasks, finishTime)
	taskResults := dbo.Tasks.ListByMultiFields(map[string]any{
		"bag_name": bagName,
		"node_id":  nodeId,
	})
	s.Len(taskResults, n)
	for _, task := range taskResults {
		s.Equal(finishTime, task.FinishTimeMs)
		s.Equal(scheduledTime, task.ScheduledTimeMs)
		s.Equal(mapTasksExitCode[task.TaskName], task.ExitCode)
	}
}

func (s *tasksTestSuite) TestTasks_GetByAccessKey() {
	dbo := db.NewDBOperations()
	testBagName := "test-bag-name"
	testAccessKey := "test-access-key"
	for i := range 10 {
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "test_access_key_" + fmt.Sprintf("%05d", i),
			BagName:  testBagName,
			TaskBusiness: models.TaskBusiness{
				AccessKey: testAccessKey + "_" + strconv.Itoa(i),
			},
		}))
	}
	task := dbo.Tasks.GetByAccessKey(testBagName, "test_access_key_00008", testAccessKey+"_8")
	s.Equal(testAccessKey+"_8", task.AccessKey)
	s.Equal("test_access_key_00008", task.TaskName)
	s.Equal(testBagName, task.BagName)
}

func (s *tasksTestSuite) TestListTasks_Prefix_Limit() {
	dbo := db.NewDBOperations()
	testBagName := "test-bag-name"
	anotherTestBagName := "another-test-bag-name"
	for i := range 10 {
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix1_" + fmt.Sprintf("%05d", i),
			BagName:  testBagName,
		}))
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix1_" + fmt.Sprintf("%05d", i+100),
			BagName:  anotherTestBagName,
		}))
	}
	for i := range 10 {
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix2_" + fmt.Sprintf("%05d", i),
			BagName:  testBagName,
		}))
		s.Nil(dbo.Tasks.Create(&models.Task{
			TaskName: "prefix2_" + fmt.Sprintf("%05d", i+100),
			BagName:  anotherTestBagName,
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
	s.DropTables()
	db.ReInitialWithDSN(s.DSN)
}

func TestTaskstestSuite(t *testing.T) {
	s := &tasksTestSuite{}
	suite.Run(t, s)
}
