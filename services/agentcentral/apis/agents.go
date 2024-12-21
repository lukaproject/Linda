package apis

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/logic/agents"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func EnableAgents(r *mux.Router) {
	r.HandleFunc("/api/agents/join/{nodeId}", nodeJoin).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/free/{nodeId}", nodeFree).Methods(http.MethodPost)
	r.HandleFunc("/api/agents/info/{nodeId}", nodeInfo).Methods(http.MethodGet)
	r.HandleFunc("/api/agents/list", listNodes).Methods(http.MethodGet)
	r.HandleFunc("/api/agents/uploadfiles", uploadFilesToNodes).Methods(http.MethodPost)
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
//	@Router			/agents/join/{nodeId} [post]
func nodeJoin(w http.ResponseWriter, r *http.Request) {
	nodeId := mux.Vars(r)["nodeId"]
	req := NodeJoinReq{}
	models.ReadJSON(r.Body, &req)
	agents.GetMgrInstance().AddNodeToBag(nodeId, req.BagName)
	w.WriteHeader(http.StatusOK)
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

// nodes list godoc
//
//	@Summary	list nodes, return all node ids
//	@Tags		agents
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]string
//	@Router		/agents/list [get]
func listNodes(w http.ResponseWriter, r *http.Request) {
	w.Write(models.Serialize(agents.GetMgrInstance().ListNodeIds()))
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
			logrus.Infof("upload files to node %s", nodeId)
			return nil
		})
	}
}
