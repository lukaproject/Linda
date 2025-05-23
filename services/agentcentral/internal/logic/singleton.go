package logic

import (
	"Linda/services/agentcentral/internal/logic/agents"
	"Linda/services/agentcentral/internal/logic/comm"
	"Linda/services/agentcentral/internal/logic/tasks"
)

func AgentsMgr() agents.Mgr {
	return agents.GetMgrInstance()
}

func InitAgentsMgr() {
	agents.InitMgrInstance()
}

func InitTasksMgr() {
	tasks.InitBagsMgrInstance()
}

func InitAsyncWorks() {
	comm.InitAsyncWorksInstance()
}
