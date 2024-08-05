package handler_test

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/handler"
	"Linda/baselibs/testcommon/testenv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type testHandlerSuite struct {
	suite.Suite
}

func (s *testHandlerSuite) TestNormal() {
	h := handler.NewHandler(config.TestConfig())
	h.Start()
	<-time.After(20 * time.Second)
}

func TestHandler(t *testing.T) {
	if !testenv.HealthCheck("http://localhost:5883/api/healthcheck") {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return
	}
	s := &testHandlerSuite{}
	suite.Run(t, s)
}
