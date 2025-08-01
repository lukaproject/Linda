//go:generate swag init
package main

// generate swagger:
// gonow generate ./...
// build dist:
// gonow build -o bin/agentcentral
import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"Linda/baselibs/abstractions/serviceskit/generator"
	"Linda/baselibs/abstractions/xdebug"
	"Linda/baselibs/abstractions/xlog"
	"Linda/services/agentcentral/apis"
	"Linda/services/agentcentral/apis/middlewares"
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
	xlog.Initial()
	c := config.Instance()

	db.InitialWithDSN(c.PGSQL_DSN)
	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler()).Methods(http.MethodGet)

	generator.Initial()
	apis.EnableHeartBeat(r)
	apis.EnableHealthCheck(r)
	apis.EnableBags(r)
	apis.EnableTasks(r)
	apis.EnableAgents(r)
	apis.EnableInnerCall(r)
	if c.Env == "debug" {
		xdebug.EnablePprof(r)
	}
	r.Use(
		middlewares.LogRequest,
		middlewares.SetHeaderJSON,
		middlewares.HTTPRecover)

	logic.InitAgentsMgr()
	logic.InitTasksMgr()
	logic.InitAsyncWorks()

	port := fmt.Sprintf(":%d", c.Port)
	xlog.Infof("serve in %s, environments is %s", port, c.Env)
	if !c.SSL.Enabled {
		xlog.Fatal(http.ListenAndServe(port, r))
	} else {
		xlog.Fatal(http.ListenAndServeTLS(port, c.SSL.CertFile, c.SSL.KeyFile, r))
	}
}
