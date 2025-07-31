package generator_test

import (
	"Linda/baselibs/abstractions/serviceskit/generator"
	"testing"
)

func TestGenNodeId(t *testing.T) {
	generator.Initial()
	t.Log(generator.GetInstance().NodeId())
}

func BenchmarkGenNodeId(b *testing.B) {
	generator.GetInstance().NodeId()
}
