package nodecontrol

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/services/integrationtest/stage"
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
	conf.BasePath = "http://localhost:5883/api"
	cli := swagger.NewAPIClient(conf)

	currentStage := &stage.Stage{}
	currentStage.SetUp(s.T(), cli)
	testNodeId := currentStage.ListNodeIds()[0]
	testBagName := currentStage.CreateBag("test-current-bag")
	currentStage.NodeOperations.JoinBag(testBagName, testNodeId)
	<-currentStage.WaitForNodeJoinFinished(testNodeId, testBagName)
	s.T().Log("join bag finished")
	currentStage.NodeOperations.FreeNode(testNodeId)
	<-currentStage.WaitForNodeFree(testNodeId)
	currentStage.DeleteBag(testBagName)
}

func TestNodeControlTestSuiteMain(t *testing.T) {
	if !stage.HealthCheck(t, stage.AgentCentralPort) {
		return
	}
	s := &nodeControlTestSuite{}
	suite.Run(t, s)
}
