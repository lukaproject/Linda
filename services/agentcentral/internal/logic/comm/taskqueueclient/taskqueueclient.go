package taskqueueclient

type Client interface {
	Enque(taskName string, bagName string, priority uint16, orderId uint32) (err error)
	Deque(bagName string) (taskName string, err error)
}
