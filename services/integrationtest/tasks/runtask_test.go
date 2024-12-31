package tasks_test

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/integrationtest/stage"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type runTaskTestSuite struct {
	testenv.TestBase
}

// TestRunTask
// about 20s for whole test case.
func (s *runTaskTestSuite) TestRunTask() {
	currentStage := stage.NewStageT(s.T())
	bagName := currentStage.CreateBag("testbag")
	defer currentStage.DeleteBag(bagName)
	s.T().Logf("bag name %s", bagName)
	test_node_id := currentStage.ListNodeIds()[0]
	// join bag
	currentStage.NodeOperations.JoinBag(bagName, test_node_id)
	<-currentStage.WaitForNodeJoinFinished(test_node_id, bagName)
	defer func() {
		// free node
		s.T().Logf("free node %s", test_node_id)
		currentStage.NodeOperations.FreeNode(test_node_id)
		s.T().Log("free node request sent")
		<-currentStage.WaitForNodeFree(test_node_id)
	}()
	// upload file to node
	blockName := "block1"
	fileName := "test.sh"
	s.Nil(currentStage.FileOperations.UploadFileContent("echo test > test.txt", "block1", "test.sh"))
	_ = xerr.Must(currentStage.Cli.AgentsApi.AgentsUploadfilesPost(context.Background(), swagger.ApisUploadFilesReq{
		Nodes: []string{test_node_id},
		Files: []swagger.ApisUploadFilesReqFiles{
			{
				Uri:          fmt.Sprintf("http://172.17.0.1:5883/api/files/download/%s/%s", blockName, fileName),
				LocationPath: "/bin/test.sh",
			},
		},
	}))
	// add task
	<-time.After(3 * time.Second)
	taskName := currentStage.TasksOperations.Add(bagName, "test-task", "/bin/test.sh", "/")
	// verify task finished
	for {
		resp := currentStage.TasksOperations.Get(bagName, taskName)
		s.T().Logf(
			"task %s, finished time %d, create time %d, schedule time %d",
			resp.TaskName, resp.FinishTimeMs, resp.CreateTimeMs, resp.ScheduledTimeMs)
		if resp.FinishTimeMs != 0 {
			break
		}
		<-time.After(2 * time.Second)
	}
}

func TestRunTask(t *testing.T) {
	if !testenv.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return
	}
	suite.Run(t, new(runTaskTestSuite))
}
