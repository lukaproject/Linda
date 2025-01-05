package tasks

import (
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
)

var (
	bagsMgrInstance BagsMgr
)

func GetBagsMgrInstance() BagsMgr {
	return bagsMgrInstance
}

func InitBagsMgrInstance() {
	bagsMgrInstance = &bagsManager{
		tasksMgrs: make(map[string]TasksMgr),
		queCli:    taskqueueclient.NewRedisTaskQueueClient(config.Instance().Redis),
	}
	bagsMgrInstance.Init()
}
