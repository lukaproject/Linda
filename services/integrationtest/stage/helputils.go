package stage

import (
	"Linda/baselibs/testcommon/testenv"
	"fmt"
	"testing"
)

// HealthCheck skip test if health check failed, and return false,
// otherwise return false.
// this function only work after dev-env setup.
func HealthCheck(t *testing.T, port int) bool {
	if !testenv.HealthCheck(
		fmt.Sprintf(
			"http://localhost:%d/api/healthcheck",
			port,
		),
	) {
		// skip e2e tests
		t.Skip("dev-env is not available, skip")
		return false
	}
	return true
}
