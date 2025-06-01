package models

// 这个请求由agent发出，用于启动agent和
// agentcentral之间的长连接。
type HeartBeatStart struct {
	Node NodeInfo
}

type HeartBeatStartResponse struct {
	Result string
}

type HeartBeatFromAgent struct {
	SeqId int64

	RunningTaskNames []string
	FinishedTasks    []FinishedTaskResult

	Node NodeInfo
}

type HeartBeatFromServer struct {
	SeqId int64

	ScheduledTasks []ScheduledTaskInfo
	// 需要下载的文件
	DownloadFiles []FileDescription

	HeartBeatEnd *HeartBeatEnd
	JoinBag      *JoinBag
	FreeNode     *FreeNode
}

type HeartBeatEnd struct {
	// 是否等待node上所有的任务都结束
	WaitingForAllTasksComplete bool
}

type HeartBeatEndResponse struct {
	Result string
}

type FreeNode struct {
	BagName string
}

type JoinBag struct {
	BagName string
}

type FileDescription struct {
	Uri          string
	LocationPath string
}

type UploadFiles struct {
	OperationId string
	Files       []FileDescription
}

type ScheduledTaskInfo struct {
	Name string
	// agent只有拿到这个AccessKey才能去访问task
	AccessKey string
}

type FinishedTaskResult struct {
	Name     string
	ExitCode int32
}
