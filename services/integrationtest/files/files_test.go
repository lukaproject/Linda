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
	targetFilePath := "testBlock/block1/test.txt.bak"
	s.Nil(currentStage.FileOperations.Upload("test file fileservice", targetFilePath))
	content := string(
		currentStage.DownloadFromURL(
			fmt.Sprintf("http://localhost:%d/files/%s",
				stage.FileServiceFEPort, targetFilePath)))
	s.Equal("test file fileservice", content)
}

func TestFilesTestSuiteMain(t *testing.T) {
	if !stage.HealthCheck(t, stage.FileServiceFEPort) {
		return
	}
	s := &filesTestSuite{}
	suite.Run(t, s)
}
