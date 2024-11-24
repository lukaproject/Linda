package apis

type Task struct {
	TaskName        string
	TaskDisplayName string
	BagName         string
	ScriptPath      string
	Priority        int16
	WorkingDir      string
	CreateTimeMs    int64
	FinishTimeMs    int64
	ScheduledTimeMs int64
	NodeId          string
}

type Bag struct {
	BagName        string
	BagDisplayName string
	CreateTimeMs   int64
	UpdateTimeMs   int64
}

type AddBagReq struct {
	BagDisplayName string `example:"test-bagDisplayName"`
}

type AddTaskReq struct {
	TaskDisplayName string `example:"test-taskDisplayName"`
	ScriptPath      string `example:"/bin/test.sh"`
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
	ErrorMsg string
}

type ListBagsResp struct {
	Bags []Bag
}

type ListBagNodesResp struct {
	NodeIds []string
}

type NodeJoinReq struct {
	BagName string
}

type NodeFreeReq struct {
}

type NodeInfo struct {
	BagName         string
	MaxRunningTasks int
}

type UploadFilesReq struct {
	Files []struct {
		// File's URI
		Uri string
		// the location of file in node
		LocationPath string
	}
	// nodes id list which will receive these files.
	Nodes []string
}
