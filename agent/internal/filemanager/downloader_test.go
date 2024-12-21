package filemanager_test

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/abstractions/xos"
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"Linda/baselibs/testcommon/testenv"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type fileDownloaderTestSuite struct {
	testenv.TestBase

	fakeFileServer fakefileserver.FileServer
}

func (ts *fileDownloaderTestSuite) TestDownloadFromPublic() {
	path := filepath.Join("test1", "p1", "public.txt")
	targetPath := filepath.Join(ts.TempDir(), "public.txt")
	ts.fakeFileServer.AddFileContent(
		filepath.Join("test1", "p1", "public.txt"),
		"test content")
	fd := &filemanager.Downloader{}
	ts.Nil(fd.DownloadFromPublicURL(
		fakefileserver.BuildDownloadURL(ts.fakeFileServer, path),
		targetPath))

	ts.Equal("test content", xos.ReadStringFromFile(targetPath))
}

func TestFileDownloadTestSuite(t *testing.T) {
	s := &fileDownloaderTestSuite{
		fakeFileServer: fakefileserver.StartT(t),
	}
	suite.Run(t, s)
}
