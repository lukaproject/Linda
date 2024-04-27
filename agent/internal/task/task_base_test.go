package task_test

import (
	"os"
	"path"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type taskBaseSuite struct {
	suite.Suite
	TempDir string
}

func (s *taskBaseSuite) SetupSuite() {
	s.TempDir = s.T().TempDir()
	s.T().Logf("setup finished, temp dir is %s", s.TempDir)
}

func (s *taskBaseSuite) SetupTest() {
	s.createTempTaskDir()
	s.T().Logf("setup test finished, temp dir is %s", s.tempTestDir())
}

func (s *taskBaseSuite) tempShellPath() string {
	return path.Join(s.TempDir, s.T().Name(), "test.sh")
}

func (s *taskBaseSuite) tempTestDir() string {
	return path.Join(s.TempDir, s.T().Name())
}

func (s *taskBaseSuite) newTempShellFile() (*os.File, error) {
	return os.Create(s.tempShellPath())
}

func (s *taskBaseSuite) createTempTaskDir() {
	xerr.Must0(os.MkdirAll(s.tempTestDir(), os.ModePerm))
}

func (s *taskBaseSuite) writeStrToTempShellFile(content string) {
	f := xerr.Must(s.newTempShellFile())
	s.T().Log(s.tempShellPath())
	defer f.Close()
	xerr.Must(f.Write([]byte(content)))
}

func (s *taskBaseSuite) getStrFromFile(filePath string) string {
	return string(xerr.Must(os.ReadFile(filePath)))
}
