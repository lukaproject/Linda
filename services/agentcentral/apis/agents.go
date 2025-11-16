package apis

import (
	"Linda/baselibs/codes/errno"
	"Linda/protocol/models"
	"Linda/services/agentcentral/apis/validator"
	"Linda/services/agentcentral/internal/logic/agents"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func EnableAgents(r *mux.Router) {
	r.HandleFunc("/api/agents/join/{nodeId}", nodeJoin).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/free/{nodeId}", nodeFree).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/info/{nodeId}", nodeInfo).Methods(http.MethodGet)
	r.HandleFunc("/api/agents/listids", listNodeIds).Methods(http.MethodGet)
	r.HandleFunc("/api/agents/list", listNodes).Methods(http.MethodGet)
	r.HandleFunc("/api/agents/uploadfiles", uploadFilesToNodes).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/{nodeId}/files/list", listNodeFiles).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/{nodeId}/files/get", getNodeFile).Methods(http.MethodPost)
}

// node join godoc
//
//	@Summary		join free node to a bag
//	@Description	join free node to a bag
//	@Tags			agents
//	@Param			nodeId		path	string				true	"node id"
//	@Param			nodeJoinReq	body	apis.NodeJoinReq	true	"Node join request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	OK
//	@Failure		409 {string}	Conflict
//	@Router			/agents/join/{nodeId} [post]
func nodeJoin(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	req := NodeJoinReq{}
	models.ReadJSON(r.Body, &req)
	err := agents.GetMgrInstance().AddNodeToBag(nodeId, req.BagName)
	statusCode := http.StatusOK
	if err != nil {
		logger.Errorf("node join bag failed, err=%v", err)
		if err == errno.ErrNodeBelongsToAnotherBag {
			statusCode = http.StatusConflict
		} else {
			statusCode = http.StatusInternalServerError
		}
	}
	w.WriteHeader(statusCode)
}

// node free godoc
//
//	@Summary		free node
//	@Description	free node
//	@Tags			agents
//	@Param			nodeId		path	string				true	"node id"
//	@Param			nodeFreeReq	body	apis.NodeFreeReq	true	"Node free request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	OK
//	@Router			/agents/free/{nodeId} [post]
func nodeFree(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	req := NodeFreeReq{}
	models.ReadJSON(r.Body, &req)
	agents.GetMgrInstance().FreeNode(nodeId)
	w.WriteHeader(http.StatusOK)
}

// node info godoc
//
//	@Summary		get node info
//	@Description	get node info by node id
//	@Tags			agents
//	@Param			nodeId	path	string	true	"node id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.NodeInfo
//	@Router			/agents/info/{nodeId} [get]
func nodeInfo(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	infoModel := agents.GetMgrInstance().GetNodeInfo(nodeId)
	if infoModel != nil {
		var nodeInfo NodeInfo
		FromNodeInfoModelToNodeInfo(infoModel, &nodeInfo)
		w.Write(models.Serialize(nodeInfo))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// nodeIds list godoc
//
//	@Summary		list nodes, return node ids by query
//	@Description	list nodes, return node ids by query, query format support prefix=, createAfter=, idAfter=, limit=.
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			prefix		query		string	false	"find all ids with this prefix"
//	@Param			createAfter	query		int64	false	"find all ids created after this time (ms)"
//	@Param			limit		query		int		false	"max count of node ids in result"
//	@Param			idAfter		query		string	false	"find all node ids which id greater or equal to this id"
//	@Success		200			{object}	[]string
//	@Router			/agents/listids [get]
func listNodeIds(w http.ResponseWriter, r *http.Request) {
	ch, err := validator.NodesListRequest(r)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := make([]string, 0)
	for nodeinfo := range ch {
		result = append(result, nodeinfo.NodeId)
	}
	w.Write(models.Serialize(result))
}

// nodes list godoc
//
//	@Summary		list nodes, return node infos by query
//	@Description	list nodes, return node infos by query, query format support prefix=, createAfter=, idAfter=, limit=.
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			prefix		query		string	false	"find all infos with this prefix"
//	@Param			createAfter	query		int64	false	"find all infos created after this time (ms)"
//	@Param			limit		query		int		false	"max count of node infos in result"
//	@Param			idAfter		query		string	false	"find all node infos which id greater or equal to this id"
//	@Success		200			{object}	[]NodeInfo
//	@Router			/agents/list [get]
func listNodes(w http.ResponseWriter, r *http.Request) {
	ch, err := validator.NodesListRequest(r)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := make([]NodeInfo, 0)
	for nodeinfo := range ch {
		result = append(result, NodeInfo{
			NodeId:          nodeinfo.NodeId,
			BagName:         nodeinfo.BagName,
			MaxRunningTasks: nodeinfo.MaxRunningTasks,
		})
	}
	w.Write(models.Serialize(result))
}

// upload files to nodes godoc
//
//	@Summary		upload files to nodes
//	@Description	upload files to nodes
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			uploadFilesReq	body	apis.UploadFilesReq	true	"upload files request"
//	@Router			/agents/uploadfiles [post]
func uploadFilesToNodes(w http.ResponseWriter, r *http.Request) {
	req := UploadFilesReq{}
	models.ReadJSON(r.Body, &req)
	fileDescriptions := make([]models.FileDescription, 0)
	for _, file := range req.Files {
		fileDescriptions = append(fileDescriptions, models.FileDescription{
			Uri:          file.Uri,
			LocationPath: file.LocationPath,
		})
	}

	for _, nodeId := range req.Nodes {
		agents.GetMgrInstance().CallAgent(nodeId, func(ag agents.Agent) error {
			ag.AddFilesUploadToNode(fileDescriptions)
			logger.Infof("upload files to node %s", nodeId)
			return nil
		})
	}
}

// Request Phase: A file operation is requested via API
// Queueing Phase: The request is queued to be sent to the agent
// Transmission Phase: Request is sent via heartbeat protocol
// Waiting Phase: The manager waits for a response (this is where your selected code runs)
// Response Phase: Agent processes and responds
// list node files godoc
//
//	@Summary		list files on a node
//	@Description	list files in a directory on a specific node
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			nodeId		path	string					true	"node id"
//	@Param			request		body	apis.ListFilesReq		true	"List files request"
//	@Success		200			{object}	apis.ListFilesResp
//	@Failure		500			{string}	string	"Internal server error"
//	@Failure		408			{string}	string	"Request timeout"
//	@Router			/agents/{nodeId}/files/list [post]
func listNodeFiles(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	req := ListFilesReq{}
	models.ReadJSON(r.Body, &req)

	operationId := generateOperationId()

	err := agents.GetMgrInstance().CallAgent(nodeId, func(ag agents.Agent) error {
		return ag.AddFileListRequest(operationId, req.LocationPath)
	})

	if err != nil {
		logger.Errorf("failed to request file list from node %s: %v", nodeId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Wait for response (implement with timeout)
	response, err := agents.GetMgrInstance().WaitForFileListResponse(nodeId, operationId, 30*time.Second)
	if err != nil {
		logger.Errorf("timeout waiting for file list response from node %s: %v", nodeId, err)
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}

	if response.Error != "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response.Error))
		return
	}

	resp := ListFilesResp{
		Files: make([]FileInfo, len(response.Files)),
	}

	for i, file := range response.Files {
		resp.Files[i] = FileInfo{
			Name:    file.Name,
			Path:    file.Path,
			Size:    file.Size,
			ModTime: file.ModTime,
			IsDir:   file.IsDir,
		}
	}

	w.Write(models.Serialize(resp))
}

func getNodeFile(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	req := GetFileReq{}
	models.ReadJSON(r.Body, &req)

	operationId := generateOperationId()
	logger.Errorf("processing get file : %s", operationId)
	err := agents.GetMgrInstance().CallAgent(nodeId, func(ag agents.Agent) error {
		return ag.AddFileGetRequest(operationId, req.LocationPath)
	})

	if err != nil {
		logger.Errorf("failed to request file from node %s: %v", nodeId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Wait for response (implement with timeout)
	response, err := agents.GetMgrInstance().WaitForFileGetResponse(nodeId, operationId, 30*time.Second)
	if err != nil {
		logger.Errorf("timeout waiting for file response from node %s: %v", nodeId, err)
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}

	if response.Error != "" {
		logger.Errorf("agent returned error for file get: %s", response.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Errorf("processed get file get: %s", operationId)
	resp := GetFileResp{
		Content: FileContent{
			FileInfo: FileInfo{
				Name:    response.Content.FileInfo.Name,
				Path:    response.Content.FileInfo.Path,
				Size:    response.Content.FileInfo.Size,
				ModTime: response.Content.FileInfo.ModTime,
				IsDir:   response.Content.FileInfo.IsDir,
			},
			Content: response.Content.Content,
		},
	}

	w.Write(models.Serialize(resp))
}

func generateOperationId() string {
	id := uuid.New()
	return fmt.Sprintf("op-%s", id.String())
}
