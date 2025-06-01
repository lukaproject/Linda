package tasks_test

import (
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/integrationtest/stage"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type runTaskTestSuite struct {
	testenv.TestBase
}

// about 20s for whole test case.
func (s *runTaskTestSuite) TestRunTask_ScriptPath() {
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
	currentStage.UploadFilesToNodes(
		[]string{testNodeId},
		[]struct{ Uri, LocationPath string }{
			{
				Uri:          fmt.Sprintf("http://172.17.0.1:%d/files/%s", stage.FileServiceFEPort, filePath),
				LocationPath: "/bin/test.sh",
			},
		})

	// add task
	<-time.After(3 * time.Second)
	taskName := currentStage.TasksOperations.Add(bagName, "test-task", "/bin/test.sh", "", "/")

	for {
		if currentStage.TasksOperations.VerifyTaskIsFinished(bagName, taskName, 0) {
			break
		}
		<-time.After(2 * time.Second)
	}
}

func (s *runTaskTestSuite) TestRunTask_Script() {
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
	// add task
	<-time.After(3 * time.Second)
	taskName := currentStage.TasksOperations.Add(bagName, "test-task", "", "echo 1", "/")

	for {
		if currentStage.TasksOperations.VerifyTaskIsFinished(bagName, taskName, 0) {
			break
		}
		<-time.After(2 * time.Second)
	}
}

func (s *runTaskTestSuite) TestRunTask_ScriptPath_ExitNonZero() {
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
	s.Nil(currentStage.FileOperations.Upload("exit 8", filePath))
	currentStage.UploadFilesToNodes(
		[]string{testNodeId},
		[]struct{ Uri, LocationPath string }{
			{
				Uri:          fmt.Sprintf("http://172.17.0.1:%d/files/%s", stage.FileServiceFEPort, filePath),
				LocationPath: "/bin/test.sh",
			},
		})

	// add task
	<-time.After(3 * time.Second)
	taskName := currentStage.TasksOperations.Add(bagName, "test-task", "/bin/test.sh", "", "/")

	for {
		if currentStage.TasksOperations.VerifyTaskIsFinished(bagName, taskName, 8) {
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
