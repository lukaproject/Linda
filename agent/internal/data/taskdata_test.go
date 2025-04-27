package data_test

import (
	"Linda/agent/internal/data"
	"Linda/baselibs/codes/errno"
	"testing"

	"github.com/lukaproject/xerr"
	"github.com/stretchr/testify/suite"
)

const (
	mockShell = "mockshell"
)

type testTaskDataTestSuite struct {
	suite.Suite
}

func (s *testTaskDataTestSuite) TestGetCommands_Script() {
	for _, testCase := range []struct {
		td             data.TaskData
		expectCommands []string
	}{
		{
			td: data.TaskData{
				Script: "pwd",
			},
			expectCommands: []string{"pwd"},
		},
		{
			td: data.TaskData{
				Script: "echo 1",
			},
			expectCommands: []string{"echo", "1"},
		},
		{
			td: data.TaskData{
				Script: "echo  1",
			},
			expectCommands: []string{"echo", "", "1"},
		},
	} {
		s.EqualValues(testCase.expectCommands, testCase.td.GetCommands(mockShell))
	}
}

func (s *testTaskDataTestSuite) TestGetCommands_ScriptPath() {
	for _, testCase := range []struct {
		td             data.TaskData
		expectCommands []string
	}{
		{
			td: data.TaskData{
				PathToScript: "/bin/test.sh",
			},
			expectCommands: []string{mockShell, "/bin/test.sh"},
		},
	} {
		s.EqualValues(testCase.expectCommands, testCase.td.GetCommands(mockShell))
	}
}

func (s *testTaskDataTestSuite) TestGetCommands_Invalid() {
	td := data.TaskData{
		PathToScript: "/bin/test.sh",
		Script:       "echo 1",
	}
	var err error = nil
	func() {
		defer xerr.Recover(&err)
		td.GetCommands(mockShell)
	}()
	s.NotNil(err)
	s.Equal(errno.ErrInvalidTaskData.Error(), err.Error())
}

func TestTaskDataMain(t *testing.T) {
	suite.Run(t, new(testTaskDataTestSuite))
}
