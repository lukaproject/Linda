package data

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/localdb"
	"encoding/json"

	"github.com/lukaproject/xerr"
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
	BagName  string
	NodeName string
}

func (nd *NodeData) getPersistor() (p *localdb.Persistor[*StringType, *NodeData]) {
	p, err := localdb.GetPersistor[*StringType, *NodeData](BucketNodeData)
	if err != nil {
		logger.Errorf("get %s persistor failed, err %v", BucketNodeData, err)
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

func (nd *NodeData) SetUp() {
	nd.Load()
	nd.NodeName = config.Instance().NodeName
	nd.Store()
}
