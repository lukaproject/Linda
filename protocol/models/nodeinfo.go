package models

import (
	"time"

	"gorm.io/gorm"
)

type NodeInfo struct {
	NodeId string `gorm:"primaryKey"`
	// 属于哪个bag name集群
	BagName         string
	MaxRunningTasks int
	// 可以由用户传入的用于获取对应Node的NodeName
	NodeName string

	CreateTimeMs int64
	UpdateTimeMs int64
}

func (nodeInfo *NodeInfo) BeforeCreate(_ *gorm.DB) (err error) {
	createTimeMs := time.Now().UnixMilli()
	nodeInfo.CreateTimeMs = createTimeMs
	nodeInfo.UpdateTimeMs = createTimeMs
	return
}

func (nodeInfo *NodeInfo) BeforeSave(tx *gorm.DB) (err error) {
	nodeInfo.UpdateTimeMs = time.Now().UnixMilli()
	return
}

type MemoryInfo struct {
	FreeMB  int64
	InUseMB int64
	TotalMB int64
}
