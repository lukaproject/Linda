package logic

import "Linda/services/agentcentral/logic/agents"

var (
	agentsMgr agents.Mgr
)

func AgentsMgr() agents.Mgr {
	return agentsMgr
}

func InitAgentsMgr() {
}
