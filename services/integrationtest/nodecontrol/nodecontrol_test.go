package nodecontrol

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/services/integrationtest/stage"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

// 这里用来测试node join bag / free node / get node info
// 相关的API
type nodeControlTestSuite struct {
	suite.Suite
}

func (s *nodeControlTestSuite) TestNormalScenario() {
	conf := swagger.NewConfiguration()
	conf.BasePath = fmt.Sprintf("http://localhost:%d/api", stage.AgentCentralPort)
	cli := swagger.NewAPIClient(conf)

	currentStage := &stage.Stage{}
	currentStage.SetUp(s.T(), cli)
	testBagName := currentStage.CreateBag("test-current-bag")
	testNodeId := currentStage.SelectOneNodeJoinToBag(testBagName)
	s.T().Log("join bag finished")
	currentStage.NodeOperations.FreeNode(testNodeId)
	<-currentStage.WaitForNodeFree(testNodeId)
	currentStage.DeleteBag(testBagName)
}

func (s *nodeControlTestSuite) TestReJoinDifferentPoolFailed() {
	conf := swagger.NewConfiguration()
	conf.BasePath = fmt.Sprintf("http://localhost:%d/api", stage.AgentCentralPort)
	cli := swagger.NewAPIClient(conf)

	currentStage := &stage.Stage{}
	currentStage.SetUp(s.T(), cli)
	testBagName := currentStage.CreateBag("test-current-bag")
	testBagName2 := currentStage.CreateBag("test-current-bag-2")

	defer func() {
		currentStage.DeleteBag(testBagName)
		currentStage.DeleteBag(testBagName2)
	}()

	testNodeId := currentStage.SelectOneNodeJoinToBag(testBagName)
	s.T().Log("join bag finished")
	s.Equal(http.StatusConflict, currentStage.NodeOperations.JoinBagWithStatusCode(testBagName2, testNodeId))
	currentStage.NodeOperations.FreeNode(testNodeId)
	<-currentStage.WaitForNodeFree(testNodeId)
}

func TestNodeControlTestSuiteMain(t *testing.T) {
	if !stage.HealthCheck(t, stage.AgentCentralPort) {
		return
	}
	s := &nodeControlTestSuite{}
	suite.Run(t, s)
}
