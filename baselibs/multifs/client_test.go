package multifs_test

import (
	"Linda/baselibs/multifs"
	"Linda/baselibs/testcommon/gen"
	"Linda/baselibs/testcommon/testenv"
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type clientTestSuite struct {
	testenv.TestBase

	lfs *multifs.FileService
}

func (s *clientTestSuite) TestUpload() {
	c := multifs.NewClient("http://localhost:5555")
	testContent := "test"
	reader := bytes.NewBuffer([]byte(testContent))
	err := c.Upload("test/test1/test.txt", "test.txt", reader)
	s.Nil(err)

	getFile := bytes.NewBuffer([]byte{})
	err = c.DownloadToWriter("test/test1/test.txt", getFile)
	s.Nil(err)
	fileContent := string(xerr.Must(io.ReadAll(getFile)))
	s.Equal(testContent, fileContent)
	s.T().Log(fileContent)
}

func (s *clientTestSuite) TestUploadLargeFile() {
	c := multifs.NewClient("http://localhost:5555")
	charset := gen.CharsetDigit + gen.CharsetLowerCase + gen.CharsetUpperCase
	testContent, err := gen.StrGenerate(charset, 1<<15, 1<<20)
	s.Nil(err)
	reader := bytes.NewBuffer([]byte(testContent))
	err = c.Upload("test/test1/test.txt", "test.txt", reader)
	s.Nil(err)

	getFile := bytes.NewBuffer([]byte{})
	err = c.DownloadToWriter("test/test1/test.txt", getFile)
	s.Nil(err)
	fileContent := string(xerr.Must(io.ReadAll(getFile)))
	s.Equal(testContent, fileContent)
	s.T().Log(fileContent[:50] + "...")
}

func (s *clientTestSuite) TestDownloadNotExistFile() {
	c := multifs.NewClient("http://localhost:5555")
	getFile := bytes.NewBuffer([]byte{})
	err := c.DownloadToWriter("test/test1/notexists.txt", getFile)
	s.NotNil(err)
	s.Equal("download file failed, status code is 404", err.Error())
}

func (s *clientTestSuite) SetupTest() {
	s.lfs = multifs.New(
		multifs.NewFileServiceInput{
			Port:    5555,
			BaseDir: s.TempDir(),
			Type:    multifs.FileServiceType_Local,
		})
	s.lfs.Start()
}

func (s *clientTestSuite) TearDownTest() {
	s.lfs.Shutdown(context.Background())
}

func TestClientTestSuiteMain(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}
