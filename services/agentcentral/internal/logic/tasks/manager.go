package tasks

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"

	"github.com/lukaproject/xerr"
)

// Mgr
// 这是用来管理tasks和bags的Mgr
type Mgr interface {
	AddBag(bag *models.Bag)
	GetBag(bagName string) (*models.Bag, error)
	DeleteBag(bagName string) error
}

type manager struct{}

func (mgr *manager) AddBag(bag *models.Bag) {
	dbi := db.Instance()
	xerr.Must0(dbi.Save(bag).Error)
}

func (mgr *manager) GetBag(bagName string) (bag *models.Bag, err error) {
	dbi := db.Instance()
	bag = &models.Bag{
		BagName: bagName,
	}
	err = dbi.First(bag).Error
	return
}

func (mgr *manager) DeleteBag(bagName string) (err error) {
	return db.Instance().Delete(&models.Bag{BagName: bagName}).Error
}
