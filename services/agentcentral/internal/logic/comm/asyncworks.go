package comm

import (
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"sync"
	"time"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

type AsyncWorks struct {
	bagsLocks *sync.Map
	cli       taskqueueclient.Client
}

// 同一时刻对于每一个bag，只能有一个task入队
func (aw *AsyncWorks) TaskEnque(
	taskName string,
	bagName string,
	priority uint16,
) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()
	NewRLocker(xerr.MustOk[*sync.Mutex](aw.bagsLocks.Load(bagName))).
		Run(func() {
			logrus.Infof("bag %s, task %s, enque start", bagName, taskName)
			dbo := db.NewDBOperations()
			count := dbo.GetBagEnqueuedTaskNumber(bagName)
			dbo.UpdateTaskOrderId(bagName, taskName, count+1)
			xerr.Must0(aw.cli.Enque(taskName, bagName, priority, count+1))
			logrus.Infof("bag %s, task %s, enque success", bagName, taskName)
		})
}

func (aw *AsyncWorks) PersistFinishedTasks(bagName string, taskNames []string) {
	db.NewDBOperations().UpdateTasksFinishTime(bagName, taskNames, time.Now().UnixMilli())
}

func (aw *AsyncWorks) PersistScheduledTasks(bagName string, taskNames []string, nodeId string) {
	db.NewDBOperations().UpdateTasksScheduledTime(bagName, taskNames, nodeId, time.Now().UnixMilli())
}

// operations for locks, so it is not a async method.

// bags locks CURD.
func (aw *AsyncWorks) AddBag(bagName string) {
	aw.bagsLocks.Store(bagName, &sync.Mutex{})
	logrus.Debugf("add bag %s 's lock", bagName)
}

func (aw *AsyncWorks) DeleteBag(bagName string) {
	aw.bagsLocks.Delete(bagName)
	logrus.Debugf("remove bag %s 's lock", bagName)
}
