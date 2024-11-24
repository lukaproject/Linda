package models

type NodeInfo struct {
	NodeId string `gorm:"primaryKey"`
	// 属于哪个bag name集群
	BagName         string
	MaxRunningTasks int
	// 可以由用户传入的用于获取对应Node的NodeName
	NodeName string `gorm:"uniqueIndex"`
}

type MemoryInfo struct {
	FreeMB  int64
	InUseMB int64
	TotalMB int64
}
