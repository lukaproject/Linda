package apis

import (
	"Linda/protocol/models"
)

// Provide some transfer functions to transfrom api types
// to protocol model types.

func FromTaskModelToTask(taskModel *models.Task, task *Task) {
	task.BagName = taskModel.BagName
	task.CreateTimeMs = taskModel.CreateTimeMs
	task.FinishTimeMs = taskModel.FinishTimeMs
	task.NodeId = taskModel.NodeId
	task.Priority = taskModel.Priority
	task.ScheduledTimeMs = taskModel.ScheduledTimeMs
	task.ScriptPath = taskModel.ScriptPath
	task.Script = taskModel.Script
	task.TaskDisplayName = taskModel.TaskDisplayName
	task.TaskName = taskModel.TaskName
	task.WorkingDir = taskModel.WorkingDir
}

func FromBagModelToBag(bagModel *models.Bag, bag *Bag) {
	bag.BagDisplayName = bagModel.BagDisplayName
	bag.BagName = bagModel.BagName
	bag.CreateTimeMs = bagModel.CreateTimeMs
	bag.UpdateTimeMs = bagModel.UpdateTimeMs
}

func FromNodeInfoModelToNodeInfo(nodeInfoModel *models.NodeInfo, nodeInfo *NodeInfo) {
	nodeInfo.BagName = nodeInfoModel.BagName
	nodeInfo.MaxRunningTasks = nodeInfoModel.MaxRunningTasks
}
