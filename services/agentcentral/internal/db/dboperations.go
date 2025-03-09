package db

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db/suboperations"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type DBOperations struct {
	dbi *gorm.DB

	NodeInfos *suboperations.NodeInfos
	Tasks     *suboperations.Tasks
	Bags      *suboperations.Bags
}

func (dbo *DBOperations) GetBagEnqueuedTaskNumber(bagName string) uint32 {
	var countType CountType
	xerr.Must0(dbo.dbi.
		Table("tasks").
		Select("COUNT(*)").
		Where("bag_name = ?", bagName).
		Where("order_id IS NOT NULL").
		Where("order_id != 0").
		Scan(&countType).Error)
	return countType.Count
}

func (dbo *DBOperations) ListBagNames() (ret []string) {
	ret = make([]string, 0)
	lst := ""
	for {
		part := make([]string, 0)
		xerr.Must0(dbo.dbi.
			Model(&models.Bag{}).
			Order("bag_name").
			Select("bag_name").
			Where("bag_name > ?", lst).
			Limit(10).
			Scan(&part).Error)
		if len(part) == 0 {
			break
		}
		ret = append(ret, part...)
		lst = part[len(part)-1]
	}
	return
}

func (dbo *DBOperations) ListBags() (ret []*models.Bag) {
	ret = make([]*models.Bag, 0)
	lst := ""
	for {
		part := make([]*models.Bag, 0)
		xerr.Must0(dbo.dbi.
			Model(&models.Bag{}).
			Order("bag_name").
			Where("bag_name > ?", lst).
			Limit(10).
			Scan(&part).Error)
		if len(part) == 0 {
			break
		}
		ret = append(ret, part...)
		lst = part[len(part)-1].BagName
	}
	return
}

func NewDBOperations() *DBOperations {
	dbi := Instance()
	dbo := &DBOperations{
		dbi: dbi,
	}
	dbo.NodeInfos = new(suboperations.NodeInfos)
	dbo.NodeInfos.Initial(dbi)
	dbo.Tasks = new(suboperations.Tasks)
	dbo.Tasks.Initial(dbi)
	dbo.Bags = new(suboperations.Bags)
	dbo.Bags.Initial(dbi)
	return dbo
}
