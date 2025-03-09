package suboperations_test

import (
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"fmt"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

type CommonTestSuite struct {
	dsn string

	suite.Suite
}

func (cts *CommonTestSuite) HealthCheckAndSetup() {
	var err error
	func() {
		cts.dsn = config.TestConfig().PGSQL_DSN
		defer xerr.Recover(&err)
		db.InitialWithDSN(cts.dsn)
	}()
	if err != nil {
		cts.T().Logf("failed to connect db, err is %v", err)
		cts.T().Skip("skip due to db env is not prepared.")
		return
	}
	cts.T().Log("success init! begin to test real db-operations test suite.")
}

func (cts *CommonTestSuite) DropTables() {
	tables := []string{"tasks", "bags", "node_infos"}
	cts.T().Logf("drop tables %v", tables)
	for _, table := range tables {
		xerr.Must0(db.Instance().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error)
	}
}
