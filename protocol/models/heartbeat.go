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

	RunningTaskIds  []string
	FinishedTaskIds []string

	Memory MemoryInfo
}

type HeartBeatFromServer struct {
	SeqId int64

	ScheduledTaskNames []string

	IsHeartBeatEnd bool
	HeartBeatEnd   HeartBeatEnd
}

type HeartBeatEnd struct {
	// 是否等待node上所有的任务都结束
	WaitingForAllTasksComplete bool
}

type HeartBeatEndResponse struct {
	Result string
}
