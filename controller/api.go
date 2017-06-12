package controller

import (
	"encoding/json"
	"net/http"
	"github.com/FINTProsjektet/fint-tech-docs-service/builder"
	"github.com/FINTProsjektet/fint-tech-docs-service/db"
)

func BuildAllProjects(w http.ResponseWriter, req *http.Request) {
	b := builder.NewBuilder()
	go b.BuildAllDocs()
	w.WriteHeader(http.StatusOK)
}

func GetAllProjects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mongo := db.NewMongo()
	defer mongo.Close()

	p := mongo.FindAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}
