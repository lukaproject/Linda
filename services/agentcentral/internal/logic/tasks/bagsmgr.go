package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/comm"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"

	"github.com/lukaproject/xerr"
)

// BagsMgr
// 这是用来管理bags的BagsMgr
type BagsMgr interface {
	Init() error
	AddBag(bag *models.Bag)
	GetBag(bagName string) (*models.Bag, error)
	DeleteBag(bagName string) error

	GetTasksMgr(bagName string) TasksMgr
}

type bagsManager struct {
	tasksMgrs map[string]TasksMgr
	queCli    taskqueueclient.Client
}

func (mgr *bagsManager) Init() error {
	//TODO
	return nil
}

func (mgr *bagsManager) AddBag(bag *models.Bag) {
	db.NewDBOperations().Bags.Create(bag)
	comm.GetAsyncWorksInstance().AddBag(bag.BagName)
}

func (mgr *bagsManager) GetBag(bagName string) (bag *models.Bag, err error) {
	func() {
		xerr.Recover(&err)
		bag = db.NewDBOperations().Bags.Get(bagName)
	}()
	return
}

func (mgr *bagsManager) DeleteBag(bagName string) (err error) {
	func() {
		xerr.Recover(&err)
		db.NewDBOperations().Bags.Delete(bagName)
		comm.GetAsyncWorksInstance().DeleteBag(bagName)
	}()
	return
}

func (mgr *bagsManager) GetTasksMgr(bagName string) TasksMgr {
	return &tasksManager{
		BagName: bagName,
		queCli:  mgr.queCli,
	}
}
