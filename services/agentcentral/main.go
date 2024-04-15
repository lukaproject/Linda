package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"Linda/services/agentcentral/apis"
	_ "Linda/services/agentcentral/docs"
	"Linda/services/agentcentral/internal/logic"
)

// @title			AgentCentral API
//
// @version		0.dev
// @description	This is agent central swagger API
// @termsOfService	http://swagger.io/terms/
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/api
func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(
		httpSwagger.Handler()).Methods(http.MethodGet)

	apis.EnableHeartBeat(r)
	apis.EnableHealthCheck(r)
	logic.InitAgentsMgr()

	port := ":5883"
	logrus.Infof("serve in %s", port)
	logrus.Fatal(http.ListenAndServe(port, r))
}
