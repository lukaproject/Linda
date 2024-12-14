package filemanager_test

import (
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mgrTestSuite struct {
	suite.Suite

	fakeFileServer fakefileserver.FileServer
}

func TestMgrTestSuiteMain(t *testing.T) {
	s := &mgrTestSuite{
		fakeFileServer: fakefileserver.StartT(t),
	}
	suite.Run(t, s)
}
