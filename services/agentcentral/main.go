package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"Linda/services/agentcentral/apis"
	_ "Linda/services/agentcentral/docs"
)

// @title AgentCentral API
// @version 0.dev
// @description This is agent central swagger API
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(
		httpSwagger.Handler()).Methods(http.MethodGet)

	apis.EnableHeartBeat(r)
	log.Fatal(http.ListenAndServe(":5883", r))
}
