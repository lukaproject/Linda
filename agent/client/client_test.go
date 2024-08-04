package client_test

import (
	"Linda/agent/client"
	"Linda/protocol/models"
	"fmt"
	"testing"
	"time"

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
			BagName: "testbagid",
		},
	})
	s.Nil(err)
	s.Equal(models.OK, resp.Result)

	for i := 0; i < 10; i++ {
		serverHB, err := cli.HeartBeat(&models.HeartBeatFromAgent{
			SeqId: int64(i),
		})

		s.Nil(err)
		s.Equal(int64(i), serverHB.SeqId)
		<-time.After(3 * time.Second)
	}

	cli.Close()
}

func TestAgentCentralClient(t *testing.T) {

	if !client.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return
	}

	s := &agentCentralClientTestSuite{
		localAgentURLFormat: "ws://localhost:5883/api/agent/heartbeat/%s",
	}
	suite.Run(t, s)
}
