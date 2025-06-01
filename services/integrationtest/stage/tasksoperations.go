package stage

import (
	"Linda/baselibs/apiscall/swagger"
	"context"
	"testing"

	"github.com/lukaproject/xerr"
)

type TasksOperations struct {
	t   *testing.T
	cli *swagger.APIClient
}

func (to *TasksOperations) Add(bagName, taskDisplayName, scriptPath, script, workingDir string) (taskName string) {
	resp, _, err := to.cli.TasksApi.BagsBagNameTasksPost(
		context.Background(),
		swagger.ApisAddTaskReq{
			TaskDisplayName: taskDisplayName,
			ScriptPath:      scriptPath,
			Script:          script,
			WorkingDir:      workingDir,
		},
		bagName,
	)
	xerr.Must0(err)
	return resp.TaskName
}

func (to *TasksOperations) Get(bagName, taskName string) swagger.ApisGetTaskResp {
	resp, _, err := to.cli.TasksApi.BagsBagNameTasksTaskNameGet(
		context.Background(),
		bagName,
		taskName,
	)
	xerr.Must0(err)
	return resp
}

// True for finished, false for not finished.
func (to *TasksOperations) VerifyTaskIsFinished(bagName, taskName string, exitCode int32) bool {
	resp := to.Get(bagName, taskName)
	to.t.Logf(
		"task %s, finished time %d, create time %d, schedule time %d, exit code %d",
		resp.TaskName,
		resp.FinishTimeMs,
		resp.CreateTimeMs,
		resp.ScheduledTimeMs,
		resp.ExitCode)
	return resp.FinishTimeMs != 0 && resp.ExitCode == exitCode
}
