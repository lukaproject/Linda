package apis

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/logic/tasks"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableTasks(r *mux.Router) {
	r.HandleFunc("/api/bags/{bagName}/tasks", addTask).Methods(http.MethodPost)
	r.HandleFunc("/api/bags/{bagName}/tasks/{taskName}", getTask).Methods(http.MethodGet)
}

// addTask godoc
//
//	@Summary		add task
//	@Description	add task
//	@Param			bagName		path	string			true	"bag's name"
//	@Param			addTaskReq	body	apis.AddTaskReq	true	"add tasks's request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.AddTaskResp
//	@Router			/bags/{bagName}/tasks [post]
func addTask(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)

	bagName := mux.Vars(r)["bagName"]
	addTaskReq := AddTaskReq{}
	models.ReadJSON(r.Body, &addTaskReq)
	task := &models.Task{
		TaskDisplayName: addTaskReq.TaskDisplayName,
		BagName:         bagName,
		ScriptPath:      addTaskReq.ScriptPath,
		WorkingDir:      addTaskReq.WorkingDir,
	}

	tasks.
		GetBagsMgrInstance().
		GetTasksMgr(bagName).
		AddTask(task)

	taskModel := tasks.
		GetBagsMgrInstance().
		GetTasksMgr(bagName).
		GetTask(task.TaskName)

	resp := AddTaskResp{}
	FromTaskModelToTask(taskModel, &resp.Task)
	w.Write(models.Serialize(resp))
}

// getTask godoc
//
//	@Summary		get task
//	@Description	get task
//	@Param			bagName		path	string	true	"bag's name"
//	@Param			taskName	path	string	true	"task's name"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.GetTaskResp
//	@Router			/bags/{bagName}/tasks/{taskName} [get]
func getTask(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bagName := mux.Vars(r)["bagName"]
	taskName := mux.Vars(r)["taskName"]
	taskModel := tasks.
		GetBagsMgrInstance().
		GetTasksMgr(bagName).GetTask(taskName)
	resp := GetTaskResp{}
	FromTaskModelToTask(taskModel, &resp.Task)
	w.Write(models.Serialize(resp))
}
