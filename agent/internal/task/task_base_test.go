package task

import (
	"Linda/agent/internal/utils"
	"os"
	"path"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type testBase struct {
	suite.Suite
	utils.TimerUtils
	TempDir string
}

func (s *testBase) SetupSuite() {
	s.TempDir = s.T().TempDir()
	s.T().Logf("setup finished, temp dir is %s", s.TempDir)
}

func (s *testBase) SetupTest() {
	s.createTempTaskDir()
	s.T().Logf("setup test finished, temp dir is %s", s.tempTestDir())
}

func (s *testBase) tempShellPath() string {
	return path.Join(s.TempDir, s.T().Name(), "test.sh")
}

func (s *testBase) tempTestDir() string {
	return path.Join(s.TempDir, s.T().Name())
}

func (s *testBase) createTempTaskDir() {
	xerr.Must0(os.MkdirAll(s.tempTestDir(), os.ModePerm))
}

func (s *testBase) writeStrToTempShellFile(content string) {
	s.T().Log(s.tempShellPath())
	s.writeStrToNewFile(s.tempShellPath(), content)
}

func (s *testBase) taskRequireDir(t Task) string {
	return path.Join(s.tempTestDir(), t.GetBag(), t.GetName())
}

func (s *testBase) createTaskRequireDir(t Task) {
	utils.MkdirAll(s.taskRequireDir(t), os.ModePerm)
}

func (s *testBase) createNewTestTask(name, bag, script string, resource int) Task {
	workingDir := path.Join(s.tempTestDir(), bag, name)
	pathToScript := path.Join(workingDir, "test.sh")
	t := NewTask(TaskData{
		Name:         name,
		Bag:          bag,
		Resource:     resource,
		PathToScript: pathToScript,
		WorkingDir:   workingDir,
		TaskDir:      workingDir,
	})
	s.createTaskRequireDir(t)
	s.writeStrToNewFile(pathToScript, script)
	return t
}

func (s *testBase) getStrFromFile(filePath string) string {
	return string(xerr.Must(os.ReadFile(filePath)))
}

func (s *testBase) writeStrToNewFile(filePath, content string) {
	f := xerr.Must(os.Create(filePath))
	defer f.Close()
	xerr.Must(f.Write([]byte(content)))
}
