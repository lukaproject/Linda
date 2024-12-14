package client

import (
	"Linda/protocol/models"
)

type MockClient struct {
	HBFromServersList []*models.HeartBeatFromServer
	HBFromAgentsList  []*models.HeartBeatFromAgent

	HBHandleFunc func(*models.HeartBeatFromAgent) *models.HeartBeatFromServer
}

func (mc *MockClient) HeartBeat(hb *models.HeartBeatFromAgent) (hbFromServer *models.HeartBeatFromServer, err error) {
	mc.HBFromAgentsList = append(mc.HBFromAgentsList, hb)
	if len(mc.HBFromServersList) == 0 {
		hbFromServer := mc.handleHBFromAgent(hb)
		hbFromServer.SeqId = hb.SeqId
		return hbFromServer, nil
	} else {
		hbFromServer = mc.HBFromServersList[0]
		mc.HBFromServersList = mc.HBFromServersList[1:]
	}
	return hbFromServer, nil
}

func (mc *MockClient) HeartBeatStart(*models.HeartBeatStart) (resp *models.HeartBeatStartResponse, err error) {
	resp = &models.HeartBeatStartResponse{
		Result: models.OK,
	}
	return
}

func (mc *MockClient) handleHBFromAgent(hbFromAgent *models.HeartBeatFromAgent) (hbFromServer *models.HeartBeatFromServer) {
	if mc.HBHandleFunc == nil {
		return &models.HeartBeatFromServer{}
	}
	return mc.HBHandleFunc(hbFromAgent)
}
