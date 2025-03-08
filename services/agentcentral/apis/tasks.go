package apis

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic/tasks"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableTasks(r *mux.Router) {
	r.HandleFunc("/api/bags/{bagName}/tasks", addTask).Methods(http.MethodPost)
	r.HandleFunc("/api/bags/{bagName}/tasks", listTasks).Methods(http.MethodGet)
	r.HandleFunc("/api/bags/{bagName}/tasks/{taskName}", getTask).Methods(http.MethodGet)
}

// addTask godoc
//
//	@Summary		add task
//	@Description	add task
//	@Tags			tasks
//	@Param			bagName		path	string			true	"bag's name"
//	@Param			addTaskReq	body	apis.AddTaskReq	true	"add tasks's request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.AddTaskResp
//	@Router			/bags/{bagName}/tasks [post]
func addTask(w http.ResponseWriter, r *http.Request) {
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
//	@Tags			tasks
//	@Param			bagName		path	string	true	"bag's name"
//	@Param			taskName	path	string	true	"task's name"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.GetTaskResp
//	@Router			/bags/{bagName}/tasks/{taskName} [get]
func getTask(w http.ResponseWriter, r *http.Request) {
	bagName := mux.Vars(r)["bagName"]
	taskName := mux.Vars(r)["taskName"]
	taskModel := tasks.
		GetBagsMgrInstance().
		GetTasksMgr(bagName).GetTask(taskName)
	resp := GetTaskResp{}
	FromTaskModelToTask(taskModel, &resp.Task)
	w.Write(models.Serialize(resp))
}

// listTasks godoc
//
//	@Summary		list tasks
//	@Description	list tasks
//	@Tags			tasks
//	@Param			bagName	path	string	true	"bag's name"
//	@Param			perfix		query		string	false	"find all tasks which taskName with this prefix"
//	@Param			createAfter	query		int64	false	"find all tasks created after this time (ms)"
//	@Param			limit		query		int		false	"max count of tasks in result"
//	@Param			idAfter		query		string	false	"find all tasks which taskName greater or equal to this id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]apis.Task
//	@Router			/bags/{bagName}/tasks [get]
func listTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	bagName := mux.Vars(r)["bagName"]
	logger.Infof("query is %v", query)
	lqp, err := abstractions.NewListQueryPacker(query)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ch := db.NewDBOperations().Tasks.List(bagName, lqp)
	result := make([]Task, 0)
	for taskModel := range ch {
		var task = Task{}
		FromTaskModelToTask(taskModel, &task)
		result = append(result, task)
	}
	w.Write(models.Serialize(result))
}
