package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/lukaproject/xerr"
)

type NodeOperations struct {
	t   *testing.T
	cli *swagger.APIClient
}

func (no *NodeOperations) JoinBag(bagName, nodeId string) {
	xerr.MustOk[int](0, no.joinBag(bagName, nodeId) == http.StatusOK)
}

func (no *NodeOperations) JoinBagWithStatusCode(bagName, nodeId string) int {
	return no.joinBag(bagName, nodeId)
}

func (no *NodeOperations) FreeNode(nodeId string) {
	_, resp := xerr.Must2(no.cli.AgentsApi.AgentsFreeNodeIdPost(
		context.Background(), swagger.ApisNodeFreeReq{}, nodeId))

	xerr.MustOk[int](0, resp.StatusCode == http.StatusOK)
}

// JoinBagWithTimeout
// 用 JoinBagWithTimeout 可以在 timeout 的时间内不停的尝试 JoinBag
// 如果失败返回false, 成功返回 true
func (no *NodeOperations) JoinBagWithTimeout(bagName, nodeId string, timeout time.Duration) (success bool) {
	endTime := time.Now().Add(timeout)
	for time.Now().Before(endTime) {
		statusCode := no.joinBag(bagName, nodeId)
		if statusCode == http.StatusOK {
			success = true
			break
		}
		if statusCode == http.StatusConflict {
			break
		}
		<-time.After(1 * time.Second)
	}
	return
}

func (no *NodeOperations) ListNodes(limit int32) []swagger.ApisNodeInfo {
	nodeInfos, resp := xerr.Must2(
		no.cli.AgentsApi.AgentsListGet(
			context.Background(),
			&swagger.AgentsApiAgentsListGetOpts{
				Limit: optional.NewInt32(limit),
			}))
	if resp.StatusCode != http.StatusOK {
		no.t.Logf("list nodes info failed, %d", resp.StatusCode)
	}
	return nodeInfos
}

func (no *NodeOperations) GetNodeInfo(nodeId string) (swagger.ApisNodeInfo, error) {
	nodeInfo, resp := xerr.Must2(no.cli.AgentsApi.AgentsInfoNodeIdGet(
		context.Background(), nodeId,
	))
	if resp.StatusCode != http.StatusOK {
		return nodeInfo, errors.New("Get node info failed")
	}
	return nodeInfo, nil
}

func (no *NodeOperations) joinBag(bagName, nodeId string) (statusCode int) {
	_, resp, _ := no.cli.AgentsApi.AgentsJoinNodeIdPost(
		context.Background(), swagger.ApisNodeJoinReq{
			BagName: bagName,
		}, nodeId)
	statusCode = resp.StatusCode
	return
}
