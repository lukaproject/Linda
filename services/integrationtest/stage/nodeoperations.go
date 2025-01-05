package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"context"
	"net/http"
	"testing"

	"github.com/lukaproject/xerr"
)

type NodeOperations struct {
	t   *testing.T
	cli *swagger.APIClient
}

func (no *NodeOperations) JoinBag(bagName, nodeId string) {
	_, resp := xerr.Must2(no.cli.AgentsApi.AgentsJoinNodeIdPost(
		context.Background(), swagger.ApisNodeJoinReq{
			BagName: bagName,
		}, nodeId))

	xerr.MustOk[int](0, resp.StatusCode == http.StatusOK)
}

func (no *NodeOperations) FreeNode(nodeId string) {
	_, resp := xerr.Must2(no.cli.AgentsApi.AgentsFreeNodeIdPost(
		context.Background(), swagger.ApisNodeFreeReq{}, nodeId))

	xerr.MustOk[int](0, resp.StatusCode == http.StatusOK)
}
