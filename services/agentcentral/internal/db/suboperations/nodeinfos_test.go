package suboperations_test

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
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
	suite.Suite

	dsn string
}

func (s *nodeInfosTestSuite) TestCreateNodeInfo_Success() {
	dbo := db.NewDBOperations()
	for i := 0; i < 10; i++ {
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
	for i := 0; i < 10; i++ {
		s.Nil(dbo.NodeInfos.Create(&models.NodeInfo{
			NodeId:   "prefix1_" + fmt.Sprintf("%05d", i),
			NodeName: strconv.Itoa(i),
		}))
	}
	for i := 0; i < 10; i++ {
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
	tables := []string{"tasks", "bags", "node_infos"}
	s.T().Logf("drop tables %v", tables)
	for _, table := range tables {
		xerr.Must0(db.Instance().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error)
	}
	db.ReInitialWithDSN(s.dsn)
}

func TestNodeInfosTestSuite(t *testing.T) {
	var err error
	s := &nodeInfosTestSuite{}
	func() {
		s.dsn = config.TestConfig().PGSQL_DSN
		defer xerr.Recover(&err)
		db.InitialWithDSN(s.dsn)
	}()
	if err != nil {
		t.Logf("failed to connect db, err is %v", err)
		return
	}
	t.Log("success init! begin to test real db-operations test suite.")
	suite.Run(t, s)
}
