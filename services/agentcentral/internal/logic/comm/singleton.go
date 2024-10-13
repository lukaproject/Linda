package comm

import (
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"sync"
)

var (
	asyncWorksInstance *AsyncWorks
	fileSaverInstnce   FileSaver
)

func InitAsyncWorksInstance() {
	asyncWorksInstance = &AsyncWorks{
		bagsLocks: &sync.Map{},
		cli:       taskqueueclient.NewRedisTaskQueueClient(config.Instance().Redis),
	}

	bagNames := db.NewDBOperations().ListBagNames()
	for _, bagName := range bagNames {
		asyncWorksInstance.AddBag(bagName)
	}
}

func GetAsyncWorksInstance() *AsyncWorks {
	return asyncWorksInstance
}

func InitFileSaver() {
	fileSaverInstnce = NewLocalFileSaver()
}

func GetFileSaverInstance() FileSaver {
	return fileSaverInstnce
}
