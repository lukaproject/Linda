package handler_test

import (
	"Linda/agent/client"
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/filemanager"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/testcommon/fake/fakefileserver"
	"Linda/baselibs/testcommon/testenv"
	"Linda/protocol/models"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type testHandlerUnitSuite struct {
	testenv.TestBase

	fileserver fakefileserver.FileServer
}

func (s *testHandlerUnitSuite) setup(localdbDir string) {
	conf := &config.Config{
		NodeId:            "test-node-id-1",
		LocalDBDir:        localdbDir,
		HeartbeatPeriodMs: 50,
	}
	config.SetInstance(conf)
	localdb.Initial()
	data.Initial()
}

func (s *testHandlerUnitSuite) TestJoinBagAndFreeNode() {
	s.setup(path.Join(s.TempDir(), "localdb"))

	joinSendOK := false
	joinReportNum := 0
	freeSendOK := false
	freeReportNum := 0
	testBagName := "testBagName"

	mockCli := &client.MockClient{
		HBHandleFunc: func(req *models.HeartBeatFromAgent) (resp *models.HeartBeatFromServer) {
			resp = &models.HeartBeatFromServer{
				SeqId: req.SeqId,
			}
			if req.Node.BagName == testBagName {
				joinReportNum++
			}
			if req.Node.BagName == "" && freeSendOK {
				freeReportNum++
			}
			if !joinSendOK {
				resp.JoinBag = &models.JoinBag{
					BagName: testBagName,
				}
				joinSendOK = true
				return
			} else if joinReportNum < 5 {
				return
			} else if !freeSendOK {
				freeSendOK = true
				resp.FreeNode = &models.FreeNode{}
				return
			}
			return
		},
	}
	h := handler.NewHandlerWithParameters(mockCli, nil, nil)
	h.Start()
	<-time.After(3 * time.Second)
	s.True(joinSendOK)
	s.GreaterOrEqual(joinReportNum, 5)
	s.GreaterOrEqual(freeReportNum, 5)
	s.True(freeSendOK)
}

func (s *testHandlerUnitSuite) TestFileDownload_PublicDownload_MockMgr() {
	s.setup(path.Join(s.TempDir(), "localdb"))
	uriPath := filepath.Join("TestFileDownload", "test1", "ok", "test.txt")
	s.Nil(s.fileserver.AddFileContent(uriPath, "test context"))
	connected := false
	addFileSendOK := false
	fileDownloadedCh := make(chan struct{}, 1)

	mockCli := &client.MockClient{
		HBHandleFunc: func(req *models.HeartBeatFromAgent) (resp *models.HeartBeatFromServer) {
			resp = &models.HeartBeatFromServer{
				SeqId: req.SeqId,
			}
			if !connected {
				connected = true
				return
			} else if !addFileSendOK {
				resp.DownloadFiles = make([]models.FileDescription, 0)
				resp.DownloadFiles = append(resp.DownloadFiles, models.FileDescription{
					Uri:          fakefileserver.BuildDownloadURL(s.fileserver, uriPath),
					LocationPath: "/test.txt",
				})
				addFileSendOK = true
				return
			}
			return
		},
	}
	mockFileMgr := &filemanager.MockMgr{}
	mockFileMgr.MockFuncs.Download = func(input filemanager.DownloadInput) error {
		s.True(input.Type.IsPublic())
		s.Equal("/test.txt", input.TargetPath)
		s.T().Logf("download %s success!", input.SourceURL)
		close(fileDownloadedCh)
		return nil
	}
	h := handler.NewHandlerWithParameters(mockCli, nil, mockFileMgr)
	h.Start()
	downloadSucc := false
	select {
	case <-fileDownloadedCh:
		downloadSucc = true
		break
	case <-time.After(3 * time.Second):
		downloadSucc = false
		break
	}
	s.True(connected)
	s.True(addFileSendOK)
	s.True(downloadSucc)
	s.Equal(1, mockFileMgr.CallCount.Download)
}

func TestHandlerUnit(t *testing.T) {
	s := &testHandlerUnitSuite{
		fileserver: fakefileserver.StartT(t),
	}
	suite.Run(t, s)
}
