package suboperations

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type Tasks struct {
	dbi *gorm.DB
}

func (t *Tasks) Initial(dbi *gorm.DB) {
	t.dbi = dbi
}

func (t *Tasks) Create(taskModel *models.Task) (err error) {
	return t.dbi.Create(taskModel).Error
}

func (t *Tasks) UpdateOrderId(
	bagName, taskName string,
	orderId uint32,
) {
	xerr.Must0(t.dbi.
		Model(&models.Task{}).
		Where("bag_name = ?", bagName).
		Where("task_name = ?", taskName).
		Update("order_id", orderId).Error)
}

func (t *Tasks) UpdateFinishedTime(
	bagName string,
	taskNames []string,
	finishTimeMs int64,
) {
	xerr.Must0(t.dbi.
		Model(&models.Task{}).
		Where("task_name IN ?", taskNames).
		Where("order_id IS NOT NULL").
		Where("order_id != 0").
		Update("finish_time_ms", finishTimeMs).
		Error)
}

func (t *Tasks) UpdateScheduledTime(
	bagName string,
	taskNames []string,
	nodeId string,
	scheduledTimeMs int64,
) {
	xerr.Must0(t.dbi.
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

// get by primary key.
func (t *Tasks) Get(bagName, taskName string) *models.Task {
	task := &models.Task{
		BagName:  bagName,
		TaskName: taskName,
	}
	xerr.Must0(t.dbi.First(task).Error)
	return task
}

// List By MultiFields
// list all tasks which suit for thie fieldsMap,
// tips:
//
//	(key, value) means columnKey eq Value
func (t *Tasks) ListByMultiFields(fieldsMap map[string]any) (tasksResult []*models.Task) {
	tasksResult = make([]*models.Task, 0)
	xerr.Must0(t.dbi.
		Where(fieldsMap).
		Find(&tasksResult).Error)
	return
}

// List By List query packer.
func (t *Tasks) List(bagName string, lqp abstractions.ListQueryPacker) (responses chan *models.Task) {
	chanSize := 10
	responses = make(chan *models.Task, chanSize)
	go listQueryAsync(responses, lqp, t.dbi, "task_name")
	return
}
