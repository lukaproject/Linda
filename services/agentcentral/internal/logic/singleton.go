package logic

import "Linda/services/agentcentral/internal/logic/agents"

func AgentsMgr() agents.Mgr {
	return agents.GetMgrInstance()
}

func InitAgentsMgr() {
	agents.InitMgrInstance()
}
