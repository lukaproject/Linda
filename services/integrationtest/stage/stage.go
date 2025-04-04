package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/lukaproject/xerr"
)

type Stage struct {
	t   *testing.T
	Cli *swagger.APIClient

	TasksOperations
	FileOperations
	NodeOperations
}

func (s *Stage) SetUp(t *testing.T, cli *swagger.APIClient) {
	s.Cli = cli
	s.t = t
	s.TasksOperations = TasksOperations{
		t:   t,
		cli: cli,
	}
	s.FileOperations = FileOperations{
		t:                     t,
		cli:                   cli,
		fileServiceFEEndPoint: fmt.Sprintf("http://172.17.0.1:%d", FileServiceFEPort),
	}
	s.NodeOperations = NodeOperations{
		t:   t,
		cli: cli,
	}
}

func (s *Stage) CreateBag(bagDisplayName string) (bagName string) {
	resp, httpResp := xerr.Must2(
		s.Cli.BagsApi.BagsPost(
			context.Background(),
			swagger.ApisAddBagReq{
				BagDisplayName: bagDisplayName,
			}))
	if httpResp.StatusCode != http.StatusOK {
		s.t.Logf("create bag failed, bag display name %s, status code %d", bagDisplayName, httpResp.StatusCode)
		return ""
	}
	bagName = resp.BagName
	return
}

func (s *Stage) DeleteBag(bagName string) {
	resp, httpResp := xerr.Must2(
		s.Cli.BagsApi.BagsBagNameDelete(context.Background(), bagName))
	if httpResp.StatusCode != http.StatusOK {
		s.t.Logf(
			"delete bag failed, bag display name %s, status code %d, error msg = %s",
			bagName, httpResp.StatusCode, resp.ErrorMsg)
	}
}

func (s *Stage) WaitForNodeJoinFinished(nodeId, bagName string) (ch chan any) {
	ch = make(chan any, 1)
	go func() {
		for {
			nodeInfo, httpResp := xerr.Must2(s.Cli.AgentsApi.AgentsInfoNodeIdGet(context.Background(), nodeId))
			if httpResp.StatusCode != http.StatusOK {
				s.t.Logf("get node info failed, %d", httpResp.StatusCode)
			}
			s.t.Log(nodeInfo)
			if nodeInfo.BagName != "" {
				s.t.Logf("join bag %s", nodeInfo.BagName)
				break
			}
			<-time.After(5 * time.Second)
		}
		close(ch)
	}()
	return ch
}

func (s *Stage) WaitForNodeFree(nodeId string) (ch chan any) {
	ch = make(chan any, 1)
	go func() {
		for {
			nodeInfo, httpResp := xerr.Must2(s.Cli.AgentsApi.AgentsInfoNodeIdGet(context.Background(), nodeId))
			if httpResp.StatusCode != http.StatusOK {
				s.t.Logf("get node info failed, %d", httpResp.StatusCode)
			}
			s.t.Log(nodeInfo)
			if nodeInfo.BagName == "" {
				s.t.Logf("free node %s", nodeId)
				break
			}
			<-time.After(5 * time.Second)
		}
		close(ch)
	}()
	return ch
}

func (s *Stage) SelectOneNodeJoinToBag(bagName string) (selectedNodeId string) {
	for {
		nodeInfos := s.NodeOperations.ListNodes(1000)
		for _, nodeInfo := range nodeInfos {
			nodeId := nodeInfo.NodeId
			if nodeInfo.BagName == "" {
				if !s.NodeOperations.JoinBagWithTimeout(bagName, nodeId, time.Second*20) {
					continue
				} else {
					s.t.Logf("successfully send join bag %s request to node %s", bagName, nodeId)
					<-s.WaitForNodeJoinFinished(nodeId, bagName)
					selectedNodeId = nodeId
					s.t.Logf("success join node %s to bag %s", nodeId, bagName)
					return
				}
			}
		}
		<-time.After(3 * time.Second)
	}
}

// DownloadFromURL
// 现在Swagger生成出来的client没法下载bytes，所以只能暂时用这种方法代替
func (s *Stage) DownloadFromURL(url string) []byte {
	resp := xerr.Must(http.Get(url))
	defer resp.Body.Close()
	return xerr.Must(io.ReadAll(resp.Body))
}

func NewStageT(t *testing.T) *Stage {
	conf := swagger.NewConfiguration()
	conf.BasePath = fmt.Sprintf("http://localhost:%d/api", AgentCentralPort)
	cli := swagger.NewAPIClient(conf)

	currentStage := &Stage{}
	currentStage.SetUp(t, cli)
	return currentStage
}
