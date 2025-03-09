package db_test

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"fmt"
	"net/url"
	"sort"
	"testing"

	"slices"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type realDBOperationsTestSuite struct {
	suite.Suite

	dsn string
}

func (s *realDBOperationsTestSuite) TestListBagNames() {
	dbo := db.NewDBOperations()
	n := 55
	bags := make([]*models.Bag, n)
	for i := range n {
		bags[i] = &models.Bag{
			BagDisplayName: fmt.Sprintf("test-bag-%d", i),
		}
		dbo.Bags.Create(bags[i])
	}

	ch := dbo.Bags.List(xerr.Must(abstractions.NewListQueryPacker(url.Values{})))
	result := make([]string, 0)
	for bagModel := range ch {
		result = append(result, bagModel.BagName)
	}
	slices.Sort(result)
	s.Len(bags, len(result))
	sort.Slice(bags, func(i, j int) bool {
		return bags[i].BagName < bags[j].BagName
	})
	for i := range n {
		s.Equal(result[i], bags[i].BagName)
	}
}

func (s *realDBOperationsTestSuite) SetupTest() {
	tables := []string{"tasks", "bags", "node_infos"}
	s.T().Logf("drop tables %v", tables)
	for _, table := range tables {
		xerr.Must0(db.Instance().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error)
	}
	db.ReInitialWithDSN(s.dsn)
}

func TestRealDBOperationsTestSuite(t *testing.T) {
	var err error
	s := &realDBOperationsTestSuite{}
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
