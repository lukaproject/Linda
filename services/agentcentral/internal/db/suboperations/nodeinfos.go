package suboperations

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type NodeInfos struct {
	dbi *gorm.DB
}

func (ni *NodeInfos) Initial(dbi *gorm.DB) {
	ni.dbi = dbi
}

func (ni *NodeInfos) Get(nodeId string) (nodeInfo *models.NodeInfo) {
	nodeInfo = &models.NodeInfo{
		NodeId: nodeId,
	}
	xerr.Must0(ni.dbi.First(nodeInfo).Error)
	return
}

// Delete
// Delete by nodeId(primary key).
func (ni *NodeInfos) Delete(nodeId string) (err error) {
	return ni.dbi.Delete(&models.NodeInfo{
		NodeId: nodeId,
	}).Error
}

func (ni *NodeInfos) DeleteByNodeName(nodeName string) (err error) {
	return ni.dbi.Delete(&models.NodeInfo{
		NodeName: nodeName,
	}).Error
}

func (ni *NodeInfos) Create(nodeInfo *models.NodeInfo) (err error) {
	return ni.dbi.Create(nodeInfo).Error
}

func (ni *NodeInfos) List(lqp abstractions.ListQueryPacker) (responses chan *models.NodeInfo) {
	chanSize := 10
	responses = make(chan *models.NodeInfo, chanSize)
	go func(
		responseChan chan *models.NodeInfo,
		listQueryPacker abstractions.ListQueryPacker,
	) {
		dbCurrent := listQueryPacker.PackListQuery("node_id", ni.dbi.Model(models.NodeInfo{}))
		rows, err := dbCurrent.Rows()
		if err != nil {
			return
		}
		for rows.Next() {
			var nodeInfo = &models.NodeInfo{}
			if err = dbCurrent.ScanRows(rows, nodeInfo); err != nil {
				break
			}
			responseChan <- nodeInfo
		}
		close(responseChan)
	}(responses, lqp)
	return
}
