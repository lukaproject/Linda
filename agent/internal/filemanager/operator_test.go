package filemanager_test

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/abstractions/xos"
	"Linda/baselibs/testcommon/gen"
	"io/fs"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testOperatorTestSuite struct {
	suite.Suite
}

func (s *testOperatorTestSuite) TestListFilesNoBlock() {
	tmpdir := path.Join(s.T().TempDir(), "TestListFilesNoBlock")
	xos.MkdirAll(tmpdir, fs.ModePerm)

	op := filemanager.Operator{}
	roles := gen.FileGenerateRoles{
		RootDir:     tmpdir,
		MaxNameLen:  5,
		MaxDirDepth: 3,
		MaxCount:    15,
	}
	s.Nil(gen.FileGenerate(roles))

	filesChan := make(chan string, 15)

	op.ListFileNames(tmpdir, filesChan)
	cnt := 0
	for {
		path, ok := <-filesChan
		if !ok {
			break
		}
		cnt++
		s.T().Logf("path=%s", path)
	}
	s.T().Log(cnt)
	s.LessOrEqual(cnt, roles.MaxCount)
}

func TestOperatorMain(t *testing.T) {
	suite.Run(t, new(testOperatorTestSuite))
}
