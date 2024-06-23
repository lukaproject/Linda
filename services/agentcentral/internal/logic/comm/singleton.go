package comm

import "sync"

var (
	asyncWorksInstance *AsyncWorks
)

func InitAsyncWorksInstance() {
	asyncWorksInstance = &AsyncWorks{
		bagsLocks: &sync.Map{},
	}
}

func GetAsyncWorksInstance() *AsyncWorks {
	return asyncWorksInstance
}
