package suboperations

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// Update Scheduled Time
func (t *Tasks) UpdateScheduledTime(
	bagName string,
	taskNames []string,
	accessKeys []string,
	nodeId string,
	scheduledTimeMs int64,
) {
	n := len(taskNames)
	xerr.MustOk[int](0, n == len(accessKeys))
	scheduledTasks := make([]map[string]any, n)
	for i := range n {
		scheduledTasks[i] = map[string]any{
			"task_name":         taskNames[i],
			"bag_name":          bagName,
			"access_key":        accessKeys[i],
			"node_id":           nodeId,
			"scheduled_time_ms": scheduledTimeMs,
		}
	}

	clauses := clause.OnConflict{
		Columns: []clause.Column{
			{Name: "task_name"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"access_key",
			"node_id",
			"scheduled_time_ms",
		}),
	}

	xerr.Must0(
		t.dbi.
			Model(&models.Task{}).
			Clauses(clauses).
			Where("order_id IS NOT NULL").
			Where("order_id != 0").
			Where("finish_time_ms IS NOT NULL").
			Create(scheduledTasks).
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

func (t *Tasks) GetByAccessKey(bagName, taskName, accessKey string) *models.Task {
	task := &models.Task{
		BagName:  bagName,
		TaskName: taskName,
		TaskBusiness: models.TaskBusiness{
			AccessKey: accessKey,
		},
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
	go listQueryAsync(responses, lqp, t.dbi.Where("bag_name = ?", bagName), "task_name")
	return
}
