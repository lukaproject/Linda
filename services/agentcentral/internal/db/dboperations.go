package db

import (
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type DBOperations struct {
	dbi *gorm.DB
}

func (dbo *DBOperations) AddBag(bag *models.Bag) {
	xerr.Must0(dbo.dbi.Create(bag).Error)
}

func (dbo *DBOperations) GetBagByBagName(bagName string) (bag *models.Bag) {
	bag = &models.Bag{
		BagName: bagName,
	}
	xerr.Must0(dbo.dbi.First(bag).Error)
	return
}

func (dbo *DBOperations) DeleteBagByBagName(bagName string) {
	xerr.Must0(dbo.dbi.Delete(&models.Bag{
		BagName: bagName,
	}).Error)
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

func (dbo *DBOperations) UpdateTaskOrderId(bagName string, taskName string, orderId uint32) {
	xerr.Must0(dbo.dbi.
		Model(&models.Task{}).
		Where("task_name = ?", taskName).
		Update("order_id", orderId).Error)
}

func (dbo *DBOperations) AddTask(task *models.Task) {
	xerr.Must0(dbo.dbi.Create(task).Error)
}

func (dbo *DBOperations) UpdateTasksFinishTime(bagName string, taskNames []string, finishTimeMs int64) {
	logger.Infof("tasks %v, bagName %s, finished in %d", taskNames, bagName, finishTimeMs)
	xerr.Must0(dbo.dbi.
		Model(&models.Task{}).
		Where("task_name IN ?", taskNames).
		Where("order_id IS NOT NULL").
		Where("order_id != 0").
		Update("finish_time_ms", finishTimeMs).
		Error)
}

func (dbo *DBOperations) UpdateTasksScheduledTime(bagName string, taskNames []string, nodeId string, scheduledTimeMs int64) {
	xerr.Must0(dbo.dbi.
		Model(&models.Task{}).
		Where("task_name IN ?", taskNames).
		Where("order_id IS NOT NULL").
		Where("order_id != 0").
		Where("finish_time_ms IS NOT NULL").
		Updates(map[string]any{
			"scheduled_time_ms": scheduledTimeMs,
			"node_id":           nodeId,
		}).
		Error)
}

func (dbo *DBOperations) GetTaskByTaskNameAndBagName(taskName string, bagName string) (task *models.Task) {
	task = &models.Task{}
	xerr.Must0(dbo.dbi.
		Where("task_name = ?", taskName).
		Where("bag_name = ?", bagName).
		First(task).Error)
	return
}

func (dbo *DBOperations) GetTaskByMultiFields(fieldsMap map[string]any) (tasksResult []*models.Task) {
	tasksResult = make([]*models.Task, 0)
	xerr.Must0(dbo.dbi.
		Where(fieldsMap).
		Find(&tasksResult).Error)
	return
}

func (dbo *DBOperations) CreateNodeInfo(nodeInfo *models.NodeInfo) error {
	return dbo.dbi.Create(nodeInfo).Error
}

func (dbo *DBOperations) DeleteNodeInfoByNodeId(nodeId string) {
	xerr.Must0(dbo.dbi.Delete(&models.NodeInfo{
		NodeId: nodeId,
	}).Error)
}

func (dbo *DBOperations) DeleteNodeInfoByNodeName(nodeName string) {
	xerr.Must0(dbo.dbi.Delete(&models.NodeInfo{
		NodeName: nodeName,
	}).Error)
}

func (dbo *DBOperations) GetNodeInfoByNodeId(nodeId string) (nodeInfo *models.NodeInfo) {
	nodeInfo = &models.NodeInfo{
		NodeId: nodeId,
	}
	xerr.Must0(dbo.dbi.First(nodeInfo).Error)
	return
}

func NewDBOperations() *DBOperations {
	return &DBOperations{
		dbi: Instance(),
	}
}
