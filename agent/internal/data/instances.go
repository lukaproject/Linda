package data

import "sync"

type InstanceManager struct {
	// 我们只需要保证不要同时写就OK了，同时读没关系
	ndMut sync.Mutex
	*NodeData
}

var (
	iMgrInstance *InstanceManager
)

func Initial() {
	iMgrInstance = &InstanceManager{
		NodeData: &NodeData{},
	}

	iMgrInstance.NodeData.Load()
}

func Instance() *InstanceManager {
	return iMgrInstance
}

func (im *InstanceManager) UpdateNodeData(updateFunc func(*NodeData) *NodeData) {
	im.ndMut.Lock()
	defer im.ndMut.Unlock()
	im.NodeData = updateFunc(im.NodeData)
}
