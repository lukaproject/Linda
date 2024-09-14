package apis

import (
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/logic/tasks"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukaproject/xerr"
)

func EnableBags(r *mux.Router) {
	r.HandleFunc("/api/bags", addBag).Methods(http.MethodPost)
	r.HandleFunc("/api/bags", listBag).Methods(http.MethodGet)
	r.HandleFunc("/api/bags/{bagName}", getBag).Methods(http.MethodGet)
	r.HandleFunc("/api/bags/{bagName}", deleteBag).Methods(http.MethodDelete)
	r.HandleFunc("/api/bagnodes/{bagName}", listBagNodes).Methods(http.MethodGet)
}

// addBag godoc
//
//	@Summary		add bag
//	@Description	create a new bag
//	@Param			addBagReq	body	apis.AddBagReq	true	"bag's request"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.AddBagResp
//	@Router			/bags [post]
func addBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bag := AddBagReq{}
	models.ReadJSON(r.Body, &bag)
	bagModel := &models.Bag{BagDisplayName: bag.BagDisplayName}
	tasks.GetBagsMgrInstance().AddBag(bagModel)

	resp := AddBagResp{}
	FromBagModelToBag(bagModel, &resp.Bag)
	w.Write(models.Serialize(resp))
}

// getBag godoc
//
//	@Summary		get bag
//	@Description	get bag
//	@Param			bagName	path	string	true	"bag's name"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	apis.GetBagResp
//	@Router			/bags/{bagName} [get]
func getBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bagName := mux.Vars(r)["bagName"]
	bagModel := xerr.Must(tasks.GetBagsMgrInstance().GetBag(bagName))
	resp := GetBagResp{}
	FromBagModelToBag(bagModel, &resp.Bag)
	w.Write(models.Serialize(resp))
}

// delete Bag godoc
//
//	@Summary		delete bag
//	@Description	delete bag
//	@Param			bagName	path	string	true	"bag's name"
//	@Accept			json
//	@Produce		json
//	@Router			/bags/{bagName} [delete]
//	@Success		200	{object}	apis.DeleteBagResp
func deleteBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bagName := mux.Vars(r)["bagName"]
	xerr.Must0(tasks.GetBagsMgrInstance().DeleteBag(bagName))
}

// list Bag godoc
//
//	@Summary		list bag [no implementation]
//	@Description	list bag
//	@Accept			json
//	@Produce		json
//	@Router			/bags [get]
//	@Success		200	{object}	apis.ListBagsResp
func listBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	// TODO
}

// list bag nodes godoc
//
//	@Summary		list bag nodes [no implementation]
//	@Description	list all node ids which belong to this node
//	@Accept			json
//	@Produce		json
//	@Router			/bagnodes/{bagName} [get]
//	@Success		200	{object}	apis.ListBagNodesResp
func listBagNodes(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	_ = mux.Vars(r)["bagName"]
}
