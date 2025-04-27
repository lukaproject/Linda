package comm

import (
	"Linda/baselibs/abstractions"
	"Linda/baselibs/abstractions/xlog"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"net/url"
	"sync"

	"github.com/lukaproject/xerr"
)

var (
	asyncWorksInstance *AsyncWorks
	logger             = xlog.NewForPackage()
)

func InitAsyncWorksInstance() {
	asyncWorksInstance = &AsyncWorks{
		bagsLocks: &sync.Map{},
		cli:       taskqueueclient.NewRedisTaskQueueClient(config.Instance().Redis),
	}
	lqp := xerr.Must(abstractions.NewListQueryPacker(url.Values{}))
	ch := db.NewDBOperations().Bags.List(lqp)
	for bagModel := range ch {
		asyncWorksInstance.AddBag(bagModel.BagName)
	}
}

func GetAsyncWorksInstance() *AsyncWorks {
	return asyncWorksInstance
}
