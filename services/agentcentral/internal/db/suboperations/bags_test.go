package suboperations_test

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/db/dbtestcommon"
	"fmt"
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type bagsTestSuite struct {
	dbtestcommon.CommonTestSuite
}

func (s *bagsTestSuite) TestBagCURD() {
	dbo := db.NewDBOperations()
	n := 10
	bags := make([]*models.Bag, n)
	for i := range n {
		bags[i] = &models.Bag{
			BagDisplayName: fmt.Sprintf("test-bag-%d", i),
		}
		dbo.Bags.Create(bags[i])
	}

	for i := range n {
		result := dbo.Bags.Get(bags[i].BagName)
		s.Equal(bags[i].BagDisplayName, result.BagDisplayName)
	}

	dbo.Bags.Delete(bags[3].BagName)

	var err error = nil
	func() {
		defer xerr.Recover(&err)
		dbo.Bags.Get(bags[3].BagName)
	}()
	s.Equal(err, gorm.ErrRecordNotFound)
}

func (s *bagsTestSuite) TestList() {
	dbo := db.NewDBOperations()
	for i := range 10 {
		s.Nil(dbo.Bags.Create(&models.Bag{
			BagName: "prefix1_" + fmt.Sprintf("%05d", i),
		}))
	}
	for i := range 10 {
		s.Nil(dbo.Bags.Create(&models.Bag{
			BagName: "prefix2_" + fmt.Sprintf("%05d", i),
		}))
	}
	check := func(prefix string, expectCount, limit int) {
		query := url.Values{}
		query.Set("prefix", prefix)
		query.Set("limit", strconv.Itoa(limit))
		ch := dbo.Bags.List(xerr.Must(abstractions.NewListQueryPacker(query)))
		cnt := 0
		for task := range ch {
			cnt++
			s.True(strings.HasPrefix(task.BagName, prefix))
		}
		s.Equal(min(expectCount, limit), cnt)
	}
	check("prefix1", 10, 5)
	check("prefix2", 10, 5)
	check("prefix1", 10, 15)
	check("prefix2", 10, 15)
}

func (s *bagsTestSuite) TestListBagNames() {
	dbo := db.NewDBOperations()
	n := 55
	bags := make([]*models.Bag, n)
	for i := range n {
		bags[i] = &models.Bag{
			BagName: fmt.Sprintf("test-list-bag-%d", i),
		}
		dbo.Bags.Create(bags[i])
	}

	ch := dbo.Bags.List(xerr.Must(abstractions.NewListQueryPacker(url.Values{
		"prefix": []string{"test-list-bag"},
	})))
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

func (s *bagsTestSuite) SetupSuite() {
	s.HealthCheckAndSetup()
	s.DropTables()
	db.ReInitialWithDSN(s.DSN)
}

func TestBagsTestSuite(t *testing.T) {
	s := &bagsTestSuite{}
	suite.Run(t, s)
}
