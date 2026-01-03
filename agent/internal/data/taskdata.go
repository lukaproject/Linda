package data

import (
	"Linda/agent/internal/localdb"
	"Linda/baselibs/codes/errno"
	"Linda/protocol/models"
	"encoding/json"
	"strings"

	"github.com/lukaproject/xerr"
)

const (
	_Slash         = "/"
	BucketTaskData = "tasks"
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
	Script       string

	// script running dir
	WorkingDir string
	// task located dir, such as stdout / stderr or others.
	TaskDir string
	State   TaskState
	Pid     int
}

func (t *TaskData) FromTaskModel(taskModel *models.Task) {
	t.Bag = taskModel.BagName
	t.Name = taskModel.TaskName
	t.PathToScript = taskModel.ScriptPath
	t.WorkingDir = taskModel.WorkingDir
	t.Script = taskModel.Script
	t.Resource = 1
}

// GetCommands
// return Script if task is a script only task, or return PathToScript
func (t *TaskData) GetCommands(defaultShell string) []string {
	if t.PathToScript != "" && t.Script != "" {
		panic(errno.ErrInvalidTaskData)
	}
	if t.PathToScript != "" {
		return []string{defaultShell, t.PathToScript}
	}
	return strings.Split(t.Script, " ")
}

func (t *TaskData) getPersistor() (p *localdb.Persistor[*StringType, *TaskData]) {
	p, err := localdb.GetPersistor[*StringType, *TaskData](BucketTaskData)
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

func (t *TaskData) key() *StringType {
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
