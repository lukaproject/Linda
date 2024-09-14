package handler_test

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/protocol/models"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type testHandlerUnitSuite struct {
	suite.Suite
}

func (s *testHandlerUnitSuite) TestJoinBag() {
	LocaldbDir := path.Join(s.T().TempDir(), "localdb")
	conf := &config.Config{
		NodeId:            "test-node-id-1",
		LocalDBDir:        LocaldbDir,
		HeartbeatPeriodMs: 50,
	}
	config.SetInstance(conf)
	localdb.Initial()
	mockCli := &client.MockClient{
		HBFromServersList: []*models.HeartBeatFromServer{
			{
				SeqId: 0,
				JoinBag: &models.JoinBag{
					BagName: "testBagName",
				},
			},
			{
				SeqId: 1,
				JoinBag: &models.JoinBag{
					BagName: "testBagName",
				},
			},
			{
				SeqId: 2,
				JoinBag: &models.JoinBag{
					BagName: "testBagName",
				},
			},
			{
				SeqId: 3,
				JoinBag: &models.JoinBag{
					BagName: "testBagName",
				},
			},
		},
	}
	h := handler.NewHandlerWithCliAndTaskMgr(mockCli, nil)
	h.Start()
	<-time.After(2 * time.Second)
	for id, v := range mockCli.HBFromAgentsList {
		if id < 20 {
			s.T().Log(v)
		}
	}
}

func TestHandlerUnit(t *testing.T) {
	s := &testHandlerUnitSuite{}
	suite.Run(t, s)
}
