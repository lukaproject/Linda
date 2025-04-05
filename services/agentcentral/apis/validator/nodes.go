package validator

import (
	"Linda/baselibs/abstractions"
	"Linda/protocol/models"
	"Linda/services/agentcentral/internal/logic/agents"
	"net/http"
)

func NodesListRequest(r *http.Request) (chan *models.NodeInfo, error) {
	query := r.URL.Query()
	logger.Infof("query is %v", query)
	lqp, err := abstractions.NewListQueryPacker(query)
	if err != nil {
		return nil, err
	}
	ch, err := agents.GetMgrInstance().List(lqp)
	if err != nil {
		return nil, err
	}
	return ch, err
}
