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
	testNodeId := currentStage.SelectOneNodeJoinToBag(bagName)
	defer func() {
		// free node
		s.T().Logf("free node %s", testNodeId)
		currentStage.NodeOperations.FreeNode(testNodeId)
		s.T().Log("free node request sent")
		<-currentStage.WaitForNodeFree(testNodeId)
	}()
	// upload file to node
	filePath := "block1/test.sh"
	s.Nil(currentStage.FileOperations.Upload("echo test > test.txt", filePath))
	_ = xerr.Must(currentStage.Cli.AgentsApi.AgentsUploadfilesPost(context.Background(), swagger.ApisUploadFilesReq{
		Nodes: []string{testNodeId},
		Files: []swagger.ApisUploadFilesReqFiles{
			{
				Uri:          fmt.Sprintf("http://172.17.0.1:%d/files/%s", stage.FileServiceFEPort, filePath),
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
	if !stage.HealthCheck(t, stage.AgentCentralPort) {
		return
	}
	if !stage.HealthCheck(t, stage.FileServiceFEPort) {
		return
	}
	suite.Run(t, new(runTaskTestSuite))
}
