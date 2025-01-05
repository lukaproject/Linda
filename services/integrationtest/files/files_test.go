package files

import (
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/integrationtest/stage"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type filesTestSuite struct {
	testenv.TestBase
}

func (s *filesTestSuite) TestFileUpload() {
	currentStage := stage.NewStageT(s.T())
	blockName := "testBlock"
	targetFileName := "test.txt.bak"
	s.Nil(currentStage.FileOperations.UploadFileContent("test file", blockName, targetFileName))
	content := string(
		currentStage.DownloadFromURL(
			fmt.Sprintf("http://localhost:5883/api/files/download/%s/%s", blockName, targetFileName)))
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
