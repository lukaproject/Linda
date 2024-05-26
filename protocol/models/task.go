package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskName        string `gorm:"primaryKey"`
	TaskDisplayName string
	BagName         string `gorm:"index"`
	ScriptPath      string
	WorkingDir      string
	AccessKey       string
	CreateTimeMs    int64
	FinishTimeMs    int64
	ScheduledTimeMs int64
}

func NewTask(taskDisplayName, bagName, scriptPath, workingDir string) (t *Task) {
	createTimeMs := time.Now().UnixMilli()
	t = &Task{
		TaskName:        uuid.NewString(),
		TaskDisplayName: taskDisplayName,
		BagName:         bagName,
		ScriptPath:      scriptPath,
		WorkingDir:      workingDir,
		CreateTimeMs:    createTimeMs,
	}
	return
}
