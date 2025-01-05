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

func (to *TasksOperations) Add(bagName, taskDisplayName, scriptPath, workingDir string) (taskName string) {
	resp, _, err := to.cli.TasksApi.BagsBagNameTasksPost(
		context.Background(),
		swagger.ApisAddTaskReq{
			TaskDisplayName: taskDisplayName,
			ScriptPath:      scriptPath,
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
