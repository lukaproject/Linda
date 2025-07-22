package taskqueueclient

// QuesMangeClient
// 用来操作queue的增删改查以及状态修改
type QuesManageClient interface {
	Create(bagName string) (err error)
	Delete(bagName string) (err error)
	Get(bagName string) (queClient QueClient, err error)
}

// QueClient
// 用来执行入队和出队
type QueClient interface {
	Enque(taskName string, priority uint16, orderId uint32) (err error)
	Deque() (taskName string, err error)
	// 一次deque count个tasks
	Deques(count int64) (taskNames []string, err error)
}
