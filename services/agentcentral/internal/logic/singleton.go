package logic

import "Linda/services/agentcentral/internal/logic/agents"

var (
	agentsMgr agents.Mgr
)

func AgentsMgr() agents.Mgr {
	return agentsMgr
}

func InitAgentsMgr() {
	agentsMgr = agents.NewMgr()
}
