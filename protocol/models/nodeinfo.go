package models

type NodeInfo struct {
	// 属于哪个tvms集群
	BagId string
}

type MemoryInfo struct {
	FreeMB  int64
	InUseMB int64
	TotalMB int64
}
