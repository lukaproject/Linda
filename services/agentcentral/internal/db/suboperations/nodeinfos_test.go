package suboperations_test

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/db/dbtestcommon"
	"Linda/services/agentcentral/internal/logic/agents"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type nodeInfosTestSuite struct {
	dbtestcommon.CommonTestSuite
}

func (s *nodeInfosTestSuite) TestCreateNodeInfo_Success() {
	dbo := db.NewDBOperations()
	for i := range 10 {
		s.Nil(dbo.NodeInfos.Create(&models.NodeInfo{
			NodeId:   agents.GenNodeId(),
			NodeName: strconv.Itoa(i),
		}))
	}
}

func (s *nodeInfosTestSuite) TestCreateNodeInfo_PrimaryKeyConflict() {
	dbo := db.NewDBOperations()
	nodeId := agents.GenNodeId()
	s.Nil(dbo.NodeInfos.Create(&models.NodeInfo{
		NodeId:   nodeId,
		NodeName: "1",
	}))
	err := dbo.NodeInfos.Create(&models.NodeInfo{
		NodeId:   nodeId,
		NodeName: "2",
	})
	s.NotNil(err)
}

func (s *nodeInfosTestSuite) TestListNodeInfo_Prefix_Limit() {
	dbo := db.NewDBOperations()
	for i := range 10 {
		s.Nil(dbo.NodeInfos.Create(&models.NodeInfo{
			NodeId:   "prefix1_" + fmt.Sprintf("%05d", i),
			NodeName: strconv.Itoa(i),
		}))
	}
	for i := range 10 {
		s.Nil(dbo.NodeInfos.Create(&models.NodeInfo{
			NodeId:   "prefix2_" + fmt.Sprintf("%05d", i),
			NodeName: strconv.Itoa(i),
		}))
	}
	check := func(prefix string, expectCount, limit int) {
		query := url.Values{}
		query.Set("prefix", prefix)
		query.Set("limit", strconv.Itoa(limit))
		ch := dbo.NodeInfos.List(xerr.Must(abstractions.NewListQueryPacker(query)))
		cnt := 0
		for nodeInfo := range ch {
			cnt++
			s.True(strings.HasPrefix(nodeInfo.NodeId, prefix))
			s.T().Log(nodeInfo.NodeId)
		}
		s.Equal(min(expectCount, limit), cnt)
	}
	check("prefix1", 10, 5)
	check("prefix2", 10, 5)
	check("prefix1", 10, 15)
	check("prefix2", 10, 15)
}

func (s *nodeInfosTestSuite) SetupSuite() {
	s.HealthCheckAndSetup()
	s.DropTables()
	db.ReInitialWithDSN(s.DSN)
}

func TestNodeInfosTestSuite(t *testing.T) {
	s := &nodeInfosTestSuite{}
	suite.Run(t, s)
}
