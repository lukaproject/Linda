package apis

import "Linda/protocol/models"

type Task struct {
	TaskName        string `json:"taskName"`
	TaskDisplayName string `json:"taskDisplayName"`
	BagName         string `json:"bagName"`
	ScriptPath      string `json:"scriptPath"`
	Script          string `json:"script"`
	Priority        int16  `json:"priority"`
	WorkingDir      string `json:"workingDir"`
	CreateTimeMs    int64  `swaggertype:"integer" format:"int64" json:"createTimeMs"`
	FinishTimeMs    int64  `swaggertype:"integer" format:"int64" json:"finishTimeMs"`
	ScheduledTimeMs int64  `swaggertype:"integer" format:"int64" json:"scheduledTimeMs"`
	NodeId          string `json:"nodeId"`
	ExitCode        int32  `swaggertype:"integer" format:"int32" json:"exitCode"`
}

type Bag struct {
	BagName        string `json:"bagName"`
	BagDisplayName string `json:"bagDisplayName"`
	CreateTimeMs   int64  `swaggertype:"integer" format:"int64" json:"createTimeMs"`
	UpdateTimeMs   int64  `swaggertype:"integer" format:"int64" json:"updateTimeMs"`
}

type AddBagReq struct {
	BagDisplayName string `example:"test-bagDisplayName"`
}

// AddTaskReq:
// script 和 scriptPath 必须有一个存在并且不能够同时存在
type AddTaskReq struct {
	TaskDisplayName string `example:"test-taskDisplayName"`
	ScriptPath      string `example:"/bin/test.sh"`
	Script          string `example:"echo 1"`
	WorkingDir      string `example:"/bin/testWorkingDir/working"`
}

type AddTaskResp struct {
	Task
}

type GetTaskResp struct {
	Task
}

type AddBagResp struct {
	Bag
}

type GetBagResp struct {
	Bag
}

type DeleteBagResp struct {
	ErrorMsg string `json:"errorMsg"`
}

type ListBagNodesResp struct {
	NodeIds []string `json:"nodeIds"`
}

type NodeJoinReq struct {
	BagName string `json:"bagName"`
}

type NodeFreeReq struct {
}

type NodeInfo struct {
	NodeId          string `json:"nodeId"`
	BagName         string `json:"bagName"`
	MaxRunningTasks int    `json:"maxRunningTasks"`
}

type UploadFilesReq struct {
	Files []struct {
		// File's URI
		Uri string `json:"uri"`
		// the location of file in node
		LocationPath string `json:"locationPath"`
	} `json:"files"`
	// nodes id list which will receive these files.
	Nodes []string `json:"nodes"`
}

type ListFilesReq struct {
	// File's location path in node
	LocationPath string `json:"locationPath"`
}

type GetFileReq struct {
	// File's location path in node
	LocationPath string `json:"locationPath"`
}

type ListFilesResp struct {
	Files []models.FileInfo `json:"files"`
}

type GetFileResp struct {
	Content models.FileContent
}
