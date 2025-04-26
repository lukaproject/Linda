package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	TaskName        string `gorm:"primaryKey"`
	TaskDisplayName string
	BagName         string `gorm:"index"`
	ScriptPath      string
	Script          string
	Priority        int16
	WorkingDir      string
	CreateTimeMs    int64
	FinishTimeMs    int64
	ScheduledTimeMs int64
	NodeId          string

	TaskBusiness
}

type TaskBusiness struct {
	AccessKey string
	OrderId   uint32
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreateTimeMs = time.Now().UnixMilli()
	if t.TaskName == "" {
		t.TaskName = uuid.NewString()
	}
	return
}

func (t *Task) GetPrimaryKeyColumn() string {
	return "task_name"
}
