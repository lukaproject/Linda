package localdb_test

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/localdb"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type localdbTestSuite struct {
	suite.Suite
}

func (s *localdbTestSuite) TestBucket() {
	localDBDir := path.Join(s.T().TempDir(), "localdb")
	c := &config.Config{
		LocalDBDir: localDBDir,
	}
	config.SetInstance(c)
	localdb.Initial()

	s.False(localdb.Instance().ExistBucket("test1"))
	s.Nil(localdb.Instance().NewBucket("test2"))
	s.True(localdb.Instance().ExistBucket("test2"))
}

func TestLocalDBMain(t *testing.T) {
	suite.Run(t, new(localdbTestSuite))
}
