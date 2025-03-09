package suboperations

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type Bags struct {
	dbi *gorm.DB
}

func (b *Bags) Initial(dbi *gorm.DB) {
	b.dbi = dbi
}

func (b *Bags) Create(bag *models.Bag) error {
	return b.dbi.Create(bag).Error
}

// Get Get by primary key
func (b *Bags) Get(bagName string) (bag *models.Bag) {
	bag = &models.Bag{
		BagName: bagName,
	}
	xerr.Must0(b.dbi.First(bag).Error)
	return
}

// Delete Delete by primary key
func (b *Bags) Delete(bagName string) {
	xerr.Must0(b.dbi.Delete(&models.Bag{
		BagName: bagName,
	}).Error)
}

func (b *Bags) List(lqp abstractions.ListQueryPacker) (responses chan *models.Bag) {
	chanSize := 10
	responses = make(chan *models.Bag, chanSize)
	go listQueryAsync(responses, lqp, b.dbi, "bag_name")
	return
}
