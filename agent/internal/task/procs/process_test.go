package procs

import (
	"Linda/baselibs/testcommon/testenv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProcessTestSuite struct {
	testenv.TestBase
}

func (s *ProcessTestSuite) Test_ProcessSuccess() {
}

func TestProcessTestSuiteMain(t *testing.T) {
	suite.Run(t, new(ProcessTestSuite))
}
