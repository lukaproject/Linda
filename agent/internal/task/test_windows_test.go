package task

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
)

type taskTestWindowsSuite struct {
	testBase
}

func TestTaskTestWindowsSuiteMain(t *testing.T) {
	if runtime.GOOS == "windows" {
		suite.Run(t, new(taskTestWindowsSuite))
	} else {
		t.Skip("only running in windows")
	}
}
