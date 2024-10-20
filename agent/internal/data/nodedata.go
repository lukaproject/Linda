package data

import (
	"Linda/agent/internal/localdb"
	"encoding/json"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
)

const (
	BucketNodeData = "Linda_NodeData"
)

var (
	KeyNodeData = NewKey("Linda_NodeData_Key")
)

// NodeData
// NodeData会存储所有node相关的会被持久化的信息
type NodeData struct {
	BagName string
}

func (nd *NodeData) getPersistor() (p *localdb.Persistor[*KeyType, *NodeData]) {
	p, err := localdb.GetPersistor[*KeyType, *NodeData](BucketNodeData)
	if err != nil {
		logrus.Errorf("get NodeData persistor failed, err %v", err)
		panic(err)
	}
	return p
}

func (nd *NodeData) Load() {
	nd.getPersistor().Get(KeyNodeData, nd)
}

func (nd *NodeData) Store() {
	nd.getPersistor().Set(KeyNodeData, nd)
}

func (nd *NodeData) Delete() {
	nd.getPersistor().Delete(KeyNodeData)
}

func (nd *NodeData) Serialize() []byte {
	return xerr.Must(json.Marshal(nd))
}

func (nd *NodeData) Deserialize(b []byte) (err error) {
	return json.Unmarshal(b, nd)
}
