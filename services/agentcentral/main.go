package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"Linda/services/agentcentral/apis"
	_ "Linda/services/agentcentral/docs"
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/db"
	"Linda/services/agentcentral/internal/logic"
)

var (
	configfile = flag.String("f", "etc/agentcentral.json", "agent central config file")
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
	flag.Parse()
	config.Initial(*configfile)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db.InitialWithDSN(config.Instance().PGSQL_DSN)
	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(
		httpSwagger.Handler()).Methods(http.MethodGet)

	apis.EnableHeartBeat(r)
	apis.EnableHealthCheck(r)
	apis.EnableBags(r)
	apis.EnableTasks(r)

	logic.InitAgentsMgr()
	logic.InitTasksMgr()

	port := fmt.Sprintf(":%d", config.Instance().Port)
	logrus.Infof("serve in %s", port)
	logrus.Fatal(http.ListenAndServe(port, r))
}
