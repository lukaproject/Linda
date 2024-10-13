package files

import (
	"Linda/baselibs/apiscall/swagger"
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/integrationtest/stage"
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type filesTestSuite struct {
	suite.Suite
}

func (s *filesTestSuite) TestFileUpload() {
	conf := swagger.NewConfiguration()
	conf.BasePath = "http://localhost:5883/api"
	cli := swagger.NewAPIClient(conf)

	currentStage := &stage.Stage{}

	tmpdir := s.T().TempDir()
	f := xerr.Must(os.Create(path.Join(tmpdir, "test.txt")))
	s.T().Log(tmpdir)
	f.Write([]byte("test file"))
	s.Nil(f.Close())
	blockName := "testBlock"
	_, _, err := cli.FilesApi.FilesUploadPost(
		context.Background(),
		"test.txt.bak",
		blockName,
		xerr.Must(os.Open(path.Join(tmpdir, "test.txt"))))
	s.Nil(err)
	content := string(
		currentStage.DownloadFromURL(
			fmt.Sprintf("http://localhost:5883/api/files/download/%s/%s", blockName, "test.txt.bak")))
	s.Equal("test file", content)
}

func TestFilesTestSuiteMain(t *testing.T) {
	if !testenv.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return
	}

	s := &filesTestSuite{}
	suite.Run(t, s)
}
