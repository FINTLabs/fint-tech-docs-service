package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FINTProsjektet/fint-tech-docs-service/config"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
	"github.com/gorilla/mux"
	"github.com/FINTProsjektet/fint-tech-docs-service/controller"
	"github.com/FINTProsjektet/fint-tech-docs-service/builder"
)

func main() {

	config.Dump()

	utils.CleanWorkspace()

	startBuilder()

	router := controller.SetupRouters()
	startServer(config.Get(), router)
}

func startBuilder() {
	b := builder.NewBuilder()
	go b.Start()
}

func startServer(c config.Config, router *mux.Router) {
	log.Printf("FINT tech docs service is listening on port %s", c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.Port), router))
}
