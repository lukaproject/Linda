package nodecontrol

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/integrationtest/stage"
	"context"
	"net/http"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

const (
	test_node_id = "test-node-id-1"
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
	testBagName := currentStage.CreateBag("test-current-bag")
	joinBagResp, resp := xerr.Must2(cli.DefaultApi.AgentsJoinNodeIdPost(
		context.Background(), swagger.ApisNodeJoinReq{
			BagName: testBagName,
		}, test_node_id))
	s.Equal(http.StatusOK, resp.StatusCode)
	s.T().Log(joinBagResp)
	<-currentStage.WaitForNodeJoinFinished(test_node_id, testBagName)
	s.T().Log("join bag finished")

	freeNodeResp, resp := xerr.Must2(cli.DefaultApi.AgentsFreeNodeIdPost(
		context.Background(), swagger.ApisNodeFreeReq{}, test_node_id))
	s.Equal(http.StatusOK, resp.StatusCode)
	s.T().Log(freeNodeResp)

	<-currentStage.WaitForNodeFree(test_node_id)
	currentStage.DeleteBag(testBagName)
}

func TestNodeControlTestSuiteMain(t *testing.T) {
	if !testenv.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return
	}

	s := &nodeControlTestSuite{}
	suite.Run(t, s)
}
