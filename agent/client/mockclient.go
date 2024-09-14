package client

import (
	"Linda/protocol/models"
)

type MockClient struct {
	HBFromServersList []*models.HeartBeatFromServer
	HBFromAgentsList  []*models.HeartBeatFromAgent
}

func (mc *MockClient) HeartBeat(hb *models.HeartBeatFromAgent) (hbFromServer *models.HeartBeatFromServer, err error) {
	mc.HBFromAgentsList = append(mc.HBFromAgentsList, hb)
	if len(mc.HBFromServersList) == 0 {
		return &models.HeartBeatFromServer{
			SeqId: hb.SeqId,
		}, nil
	}
	hbFromServer = mc.HBFromServersList[0]
	mc.HBFromServersList = mc.HBFromServersList[1:]
	return hbFromServer, nil
}

func (mc *MockClient) HeartBeatStart(*models.HeartBeatStart) (resp *models.HeartBeatStartResponse, err error) {
	resp = &models.HeartBeatStartResponse{
		Result: models.OK,
	}
	return
}
