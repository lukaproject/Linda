package agents_test

import (
	"Linda/services/agentcentral/internal/logic/agents"
	"testing"
)

func TestGenNodeId(t *testing.T) {
	t.Log(agents.GenNodeId())
}

func BenchmarkGenNodeId(b *testing.B) {
	agents.GenNodeId()
}
