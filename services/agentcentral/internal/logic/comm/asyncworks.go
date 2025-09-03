package comm

import (
	"Linda/baselibs/abstractions/xctx"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"sync"
	"time"

	"github.com/lukaproject/xerr"
)

type AsyncWorks struct {
	bagsLocks     *sync.Map
	quesManageCli taskqueueclient.QuesManageClient
}

// 同一时刻对于每一个bag，只能有一个task入队
func (aw *AsyncWorks) TaskEnque(
	taskName string,
	bagName string,
	priority uint16,
) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()
	xctx.NewLocker(xerr.MustOk[*sync.Mutex](aw.bagsLocks.Load(bagName))).
		Run(func() {
			logger.Infof("bag %s, task %s, enque start", bagName, taskName)
			dbo := db.NewDBOperations()
			count := dbo.GetBagEnqueuedTaskNumber(bagName)
			dbo.Tasks.UpdateOrderId(bagName, taskName, count+1)
			xerr.Must0(xerr.Must(aw.quesManageCli.Get(bagName)).Enque(taskName, priority, count+1))
			logger.Infof("bag %s, task %s, enque success", bagName, taskName)
		})
}

func (aw *AsyncWorks) TaskDeque(bagName string) (taskName string, err error) {
	queCli, err := aw.quesManageCli.Get(bagName)
	if err != nil {
		return
	}
	return queCli.Deque()
}

func (aw *AsyncWorks) PersistFinishedTasks(bagName string, tasks []models.FinishedTaskResult) {
	db.NewDBOperations().Tasks.PersistFinishedTasks(bagName, tasks, time.Now().UnixMilli())
}

func (aw *AsyncWorks) PersistScheduledTasks(bagName string, tasks []models.ScheduledTaskInfo, nodeId string) {
	taskNames := make([]string, 0)
	accessKeys := make([]string, 0)
	for _, taskInfo := range tasks {
		taskNames = append(taskNames, taskInfo.Name)
		accessKeys = append(accessKeys, taskInfo.AccessKey)
	}
	db.NewDBOperations().Tasks.UpdateScheduledTime(bagName, taskNames, accessKeys, nodeId, time.Now().UnixMilli())
}

// operations for locks, so it is not a async method.

// bags locks CURD.
func (aw *AsyncWorks) AddBag(bagName string) {
	aw.bagsLocks.Store(bagName, &sync.Mutex{})
	logger.Debugf("add bag %s 's lock", bagName)
}

func (aw *AsyncWorks) DeleteBag(bagName string) {
	aw.bagsLocks.Delete(bagName)
	logger.Debugf("remove bag %s 's lock", bagName)
}

func (aw *AsyncWorks) Initial() {

}
