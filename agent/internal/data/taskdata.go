package data

import (
	"Linda/agent/internal/localdb"
	"Linda/protocol/models"
	"encoding/json"

	"github.com/lukaproject/xerr"
)

const (
	_Slash         = "/"
	BucketTaskData = ""
)

type TaskState int

const (
	TaskState_Waiting = iota
	TaskState_Preparing
	TaskState_Running
	TaskState_Finished
)

type TaskData struct {
	Name         string
	Bag          string
	Resource     int
	PathToScript string

	// script running dir
	WorkingDir string
	// task data located dir, such as stdout / stderr or others.
	TaskDir string
	State   TaskState
}

func (t *TaskData) FromTaskModel(taskModel *models.Task) {
	t.Bag = taskModel.BagName
	t.Name = taskModel.TaskName
	t.PathToScript = taskModel.ScriptPath
	t.WorkingDir = taskModel.WorkingDir
	t.Resource = 1
}

func (t *TaskData) getPersistor() (p *localdb.Persistor[*KeyType, *TaskData]) {
	p, err := localdb.GetPersistor[*KeyType, *TaskData](BucketTaskData)
	if err != nil {
		logger.Errorf("get %s persistor failed, err %v", BucketTaskData, err)
		panic(err)
	}
	return p
}

func (t *TaskData) Serialize() []byte {
	return xerr.Must(json.Marshal(t))
}

func (t *TaskData) Deserialize(b []byte) error {
	return json.Unmarshal(b, t)
}

func (t *TaskData) key() *KeyType {
	return NewKey(t.Bag + _Slash + t.Name)
}

func (t *TaskData) Load() {
	key := t.key()
	xerr.Must0(t.getPersistor().Get(key, t))
}

func (t *TaskData) Store() {
	key := t.key()
	xerr.Must0(t.getPersistor().Set(key, t))
}
