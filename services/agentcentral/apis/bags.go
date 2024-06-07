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
}

// addBag godoc
//
//	@Summary		add bag
//	@Description	create a new bag
//	@Param			addBagReq	body	apis.AddBagReq	true	"bag's request"
//	@Accept			json
//	@Produce		json
//	@Router			/bags [post]
func addBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bag := &AddBagReq{}
	models.ReadJSON(r, bag)
	bagModel := models.NewBag(bag.BagDisplayName)
	tasks.GetBagsMgrInstance().AddBag(bagModel)
	w.Write(models.Serialize(bagModel))
}

// getBag godoc
//
//	@Summary		get bag
//	@Description	get bag
//	@Param			bagName	path	string	true	"bag's name"
//	@Accept			json
//	@Produce		json
//	@Router			/bags/{bagName} [get]
func getBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bagName := mux.Vars(r)["bagName"]
	bagModel := xerr.Must(tasks.GetBagsMgrInstance().GetBag(bagName))
	w.Write(models.Serialize(bagModel))
}

// delete Bag godoc
//
//	@Summary		get bag
//	@Description	get bag
//	@Param			bagName	path	string	true	"bag's name"
//	@Accept			json
//	@Produce		json
//	@Router			/bags/{bagName} [delete]
func deleteBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	bagName := mux.Vars(r)["bagName"]
	xerr.Must0(tasks.GetBagsMgrInstance().DeleteBag(bagName))
}

// list Bag godoc
//
//	@Summary		get bag
//	@Description	get bag
//	@Accept			json
//	@Produce		json
//	@Router			/bags [get]
func listBag(w http.ResponseWriter, r *http.Request) {
	defer httpRecover(w, r)
	// TODO
}
