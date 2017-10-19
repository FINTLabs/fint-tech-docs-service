package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FINTprosjektet/fint-tech-docs-service/builder"
	"github.com/FINTprosjektet/fint-tech-docs-service/config"
	"github.com/FINTprosjektet/fint-tech-docs-service/controller"
	"github.com/FINTprosjektet/fint-tech-docs-service/util"
)

func main() {
	log.Println("Firing up all cylinders... . .. . ..")
	
	config.Dump()

	util.CleanWorkspace()

	startBuilder()

	cfg := config.Get()
	router := controller.SetupRouters(cfg.Webroot)
	startServer(cfg, router)
}

func startBuilder() {
	b := builder.New()
	go b.Start()
}

func startServer(c config.Config, router http.Handler) {
	log.Printf("FINT tech docs service is listening on the port %s", c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.Port), router))
}
