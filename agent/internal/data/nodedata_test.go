package data_test

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/localdb"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testNodeDataSuite struct {
	suite.Suite
}

func (s *testNodeDataSuite) TestLoadAndStore() {
	LocaldbDir := path.Join(s.T().TempDir(), "localdb")
	conf := &config.Config{
		NodeId:            "test-node-id-1",
		LocalDBDir:        LocaldbDir,
		HeartbeatPeriodMs: 50,
	}
	config.SetInstance(conf)
	localdb.Initial()
	nd := &data.NodeData{
		BagName: "test-bag-name",
	}
	nd.Store()
	newNd := &data.NodeData{}
	newNd.Load()
	s.Equal(nd.BagName, newNd.BagName)
}

func TestNodeData(t *testing.T) {
	suite.Run(t, new(testNodeDataSuite))
}
