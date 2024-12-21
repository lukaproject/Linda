package comm_test

import (
	"Linda/baselibs/testcommon/testenv"
	"Linda/services/agentcentral/internal/logic/comm"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testFileSaverSuite struct {
	testenv.TestBase
}

func (s *testFileSaverSuite) tmpFileConent(size int) string {
	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		sb.WriteString("0")
	}
	return sb.String()
}

func (s *testFileSaverSuite) TestLocalFileSaver() {
	tmpdir := s.TempDir()
	saver := comm.NewLocalFileSaver()

	originPath := path.Join(tmpdir, "testlocalfilesaver.txt")
	f, err := os.Create(originPath)
	expectSize := 1 << 15
	s.Nil(err)
	content := s.tmpFileConent(expectSize)
	s.Equal(expectSize, len(content))
	n, err := f.WriteString(content)
	s.Nil(err)
	s.Equal(expectSize, n)
	s.Nil(f.Close())

	targetPath := path.Join(tmpdir, "testlocalfilesaver.txt.bak")

	f, err = os.Open(originPath)
	s.Nil(err)
	saver.WriteWithReader(targetPath, f)
	s.Nil(f.Close())

	f, err = os.Open(originPath)
	s.Nil(err)
	copyContent, err := io.ReadAll(f)
	s.Nil(err)
	s.Equal(content, string(copyContent))
	s.Nil(f.Close())
}

func TestFileSaverMain(t *testing.T) {
	suite.Run(t, new(testFileSaverSuite))
}
