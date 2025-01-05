package task

import (
	"Linda/agent/internal/data"
	"Linda/agent/internal/utils"
	"Linda/baselibs/abstractions/xos"
	"Linda/baselibs/testcommon/testenv"
	"os"
	"path"

	"github.com/lukaproject/xerr"
)

type testBase struct {
	testenv.TestBase
	utils.TimerUtils
}

func (s *testBase) SetupTest() {
	s.TempDir()
	s.T().Logf("setup test finished, temp dir is %s", s.TempDir())
}

func (s *testBase) tempShellPath() string {
	return path.Join(s.TempDir(), "test.sh")
}

func (s *testBase) writeStrToTempShellFile(content string) {
	s.T().Log(s.tempShellPath())
	s.writeStrToNewFile(s.tempShellPath(), content)
}

func (s *testBase) taskRequireDir(t Task) string {
	return path.Join(s.TempDir(), t.GetBag(), t.GetName())
}

func (s *testBase) createTaskRequireDir(t Task) {
	xos.MkdirAll(s.taskRequireDir(t), os.ModePerm)
}

func (s *testBase) createNewTestTask(name, bag, script string, resource int) Task {
	workingDir := path.Join(s.TempDir(), bag, name)
	pathToScript := path.Join(workingDir, "test.sh")
	t := NewTask(data.TaskData{
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
