package client_test

import (
	"Linda/agent/client"
	"Linda/protocol/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type agentCentralClientTestSuite struct {
	suite.Suite

	localAgentURLFormat string
}

func (s *agentCentralClientTestSuite) TestNormal() {
	cli, err := client.New(fmt.Sprintf(s.localAgentURLFormat, "12121212"))
	s.Nil(err)
	resp, err := cli.HeartBeatStart(&models.HeartBeatStart{
		Node: models.NodeInfo{
			BagId: "testbagid",
		},
	})
	s.Nil(err)
	s.Equal(models.OK, resp.Result)

	serverHB, err := cli.HeartBeat(&models.HeartBeatFromAgent{
		SeqId: 0,
	})

	s.Nil(err)
	s.Equal(int64(0), serverHB.SeqId)
	cli.Close()
}

func TestAgentCentralClient(t *testing.T) {

	if !client.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		return
	}

	s := &agentCentralClientTestSuite{
		localAgentURLFormat: "ws://localhost:5883/api/agent/heartbeat/%s",
	}
	suite.Run(t, s)
}
