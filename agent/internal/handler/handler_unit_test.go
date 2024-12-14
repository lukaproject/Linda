package handler_test

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
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
	localdbDir := path.Join(s.T().TempDir(), "TestJoinBag", "localdb")
	conf := &config.Config{
		NodeId:            "test-node-id-1",
		LocalDBDir:        localdbDir,
		HeartbeatPeriodMs: 50,
	}
	config.SetInstance(conf)
	localdb.Initial()
	data.Initial()
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
	h := handler.NewHandlerWithParameters(mockCli, nil)
	h.Start()
	<-time.After(2 * time.Second)
	count := 0
	for _, v := range mockCli.HBFromAgentsList {
		if v.Node.BagName == "testBagName" {
			count++
		}
	}
	s.Greater(count, 5)
}

func (s *testHandlerUnitSuite) TestJoinBagAndFreeNode() {
	localdbDir := path.Join(s.T().TempDir(), "TestJoinBagAndFreeNode", "localdb")
	conf := &config.Config{
		NodeId:            "test-node-id-1",
		LocalDBDir:        localdbDir,
		HeartbeatPeriodMs: 50,
	}
	config.SetInstance(conf)
	localdb.Initial()
	data.Initial()

	joinSendOK := false
	joinReportNum := 0
	freeSendOK := false
	freeReportNum := 0
	testBagName := "testBagName"

	mockCli := &client.MockClient{
		HBHandleFunc: func(req *models.HeartBeatFromAgent) (resp *models.HeartBeatFromServer) {
			resp = &models.HeartBeatFromServer{
				SeqId: req.SeqId,
			}
			if req.Node.BagName == testBagName {
				joinReportNum++
			}
			if req.Node.BagName == "" && freeSendOK {
				freeReportNum++
			}
			if !joinSendOK {
				resp.JoinBag = &models.JoinBag{
					BagName: testBagName,
				}
				joinSendOK = true
				return
			} else if joinReportNum < 5 {
				return
			} else if !freeSendOK {
				freeSendOK = true
				resp.FreeNode = &models.FreeNode{}
				return
			}
			return
		},
	}
	h := handler.NewHandlerWithParameters(mockCli, nil)
	h.Start()
	<-time.After(3 * time.Second)
	s.True(joinSendOK)
	s.GreaterOrEqual(joinReportNum, 5)
	s.GreaterOrEqual(freeReportNum, 5)
	s.True(freeSendOK)
}

func TestHandlerUnit(t *testing.T) {
	s := &testHandlerUnitSuite{}
	suite.Run(t, s)
}
