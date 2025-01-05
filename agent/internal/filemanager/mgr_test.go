package filemanager_test

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"Linda/baselibs/testcommon/testenv"
	"io"
	"os"
	"path"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type mgrTestSuite struct {
	testenv.TestBase

	fakeFileServer fakefileserver.FileServer
}

func (s *mgrTestSuite) TestPublicFileDownload_Success() {
	tmpdir := s.TempDir()
	content := "aaaaaaaa"
	s.fakeFileServer.AddFileContent(path.Join("TestPublicFileDownload", "a.txt"), content)
	url := fakefileserver.BuildDownloadURL(s.fakeFileServer, path.Join("TestPublicFileDownload", "a.txt"))
	mgr := filemanager.NewMgr()
	targetPath := path.Join(tmpdir, "a.txt")
	s.Nil(mgr.Download(filemanager.DownloadInput{
		Type:       filemanager.DownloadSourceType_public,
		SourceURL:  url,
		TargetPath: targetPath,
	}))
	f := xerr.Must(os.Open(targetPath))
	defer f.Close()
	s.Equal(content, string(xerr.Must(io.ReadAll(f))))
}

func TestMgrTestSuiteMain(t *testing.T) {
	s := &mgrTestSuite{
		fakeFileServer: fakefileserver.StartT(t),
	}
	suite.Run(t, s)
}
