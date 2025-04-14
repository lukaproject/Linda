package config

import (
	"Linda/baselibs/testcommon/testenv"
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type testConfigTestSuite struct {
	testenv.TestBase
}

func (s *testConfigTestSuite) TestLoadFromOSEnvAndConfigFile() {
	tmpdir := s.TempDir()
	*configFile = path.Join(tmpdir, "testconfig.json")
	s.T().Setenv("LINDA_AGENT_CENTRAL_ENDPOINT", "127.0.0.1:5883")
	s.T().Setenv("LINDA_NODE_ID", "test-node-id")
	s.T().Setenv("LINDA_NODE_NAME", "test-node-name-env")
	configMap := map[string]string{
		"NodeName": "testsuite-testname",
	}
	f, err := os.Create(*configFile)
	s.Nil(err)
	f.Write(xerr.Must(json.Marshal(configMap)))
	s.Nil(f.Close())

	Initial()

	s.Equal("testsuite-testname", c.NodeName)
	s.Equal("127.0.0.1:5883", c.AgentCentralEndPoint)
	s.Equal("test-node-id", c.NodeId)
	s.Equal("/linda/tasks", c.TasksDir)
	s.Equal(50, c.HeartbeatPeriodMs)
}

func TestConfigMain(t *testing.T) {
	suite.Run(t, new(testConfigTestSuite))
}
