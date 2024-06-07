package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"

	"github.com/lukaproject/xerr"
)

// BagsMgr
// 这是用来管理bags的BagsMgr
type BagsMgr interface {
	AddBag(bag *models.Bag)
	GetBag(bagName string) (*models.Bag, error)
	DeleteBag(bagName string) error

	GetTasksMgr(bagName string) TasksMgr
}

type bagsManager struct{}

func (mgr *bagsManager) AddBag(bag *models.Bag) {
	dbi := db.Instance()
	xerr.Must0(dbi.Save(bag).Error)
}

func (mgr *bagsManager) GetBag(bagName string) (bag *models.Bag, err error) {
	dbi := db.Instance()
	bag = &models.Bag{
		BagName: bagName,
	}
	err = dbi.First(bag).Error
	return
}

func (mgr *bagsManager) DeleteBag(bagName string) (err error) {
	return db.Instance().Delete(&models.Bag{BagName: bagName}).Error
}

func (mgr *bagsManager) GetTasksMgr(bagName string) TasksMgr {
	return &tasksManager{
		BagName: bagName,
	}
}
