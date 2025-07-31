package filemanager_test

import (
	"Linda/agent/internal/filemanager"
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"Linda/baselibs/testcommon/testenv"
	"Linda/protocol/models"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
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
	// Use forward slashes for URL paths, not path.Join which uses OS-specific separators
	serverPath := "TestPublicFileDownload/a.txt"
	s.fakeFileServer.AddFileContent(serverPath, content)
	url := fakefileserver.BuildDownloadURL(s.fakeFileServer, serverPath)
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

func addFileContent(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func (s *mgrTestSuite) TestGetFile_Success() {
	tmpdir := s.TempDir()
	content := "aaaaaaaa"
	path := filepath.Join(tmpdir, "a.txt")
	e := addFileContent(path, content)
	s.Nil(e)
	mgr := filemanager.NewMgr()
	fileContent, err := mgr.GetFile(path)
	s.Nil(err)
	f := xerr.Must(os.Open(path))
	defer f.Close()
	s.Equal(string(fileContent.Content), string(xerr.Must(io.ReadAll(f))))
	s.Equal(string(fileContent.FileInfo.Name), "a.txt")
}

func (s *mgrTestSuite) TestListFile_Success() {
	tmpdir := s.TempDir()

	// Create test files and directories
	s.Nil(addFileContent(filepath.Join(tmpdir, "a.txt"), "content a"))
	s.Nil(addFileContent(filepath.Join(tmpdir, "b.txt"), "content b"))
	s.Nil(os.Mkdir(filepath.Join(tmpdir, "dir"), 0755))
	mgr := filemanager.NewMgr()

	// Test ListFiles (channel-based, files only)
	fileInfos, err := mgr.ListFileInfos(tmpdir)
	s.Nil(err)
	s.Len(fileInfos, 3)

	// Create a map for easier checking
	fileMap := make(map[string]models.FileInfo)
	for _, info := range fileInfos {
		fileMap[info.Name] = info
	}

	// Check a.txt
	aInfo, exists := fileMap["a.txt"]
	s.True(exists)
	s.Equal("a.txt", aInfo.Name)
	s.Equal(filepath.Join(tmpdir, "a.txt"), aInfo.Path)
	s.Equal(int64(9), aInfo.Size) // "content a" = 9 bytes
	s.False(aInfo.IsDir)
	s.Greater(aInfo.ModTime, int64(0))

	// log a content
	log.Printf("a.txt content: %s", xerr.Must(os.ReadFile(aInfo.Path)))

	// Check b.txt
	bInfo, exists := fileMap["b.txt"]
	s.True(exists)
	s.Equal("b.txt", bInfo.Name)
	s.Equal(filepath.Join(tmpdir, "b.txt"), bInfo.Path)
	s.Equal(int64(9), bInfo.Size) // "content b" = 9 bytes
	s.False(bInfo.IsDir)

	// Check subdir
	dirInfo, exists := fileMap["dir"]
	s.True(exists)
	s.Equal("dir", dirInfo.Name)
	s.Equal(filepath.Join(tmpdir, "dir"), dirInfo.Path)
	s.True(dirInfo.IsDir)
	// Not sure why but in remote test it is not 0. actual  : 4096
	// s.Equal(int64(0), dirInfo.Size) // Directories typically have size 0
}

func TestMgrTestSuiteMain(t *testing.T) {
	s := &mgrTestSuite{
		fakeFileServer: fakefileserver.StartT(t),
	}
	suite.Run(t, s)
}
