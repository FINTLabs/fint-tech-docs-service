package controller

import (
	"encoding/json"
	"net/http"
	"github.com/FINTProsjektet/fint-tech-docs-service/builder"
	"github.com/FINTProsjektet/fint-tech-docs-service/db"
)

// BuildAllProjects ...
func BuildAllProjects(w http.ResponseWriter, req *http.Request) {
	b := builder.New()
	go b.BuildAllJavaDocs()
	w.WriteHeader(http.StatusOK)
}

// GetAllProjects ...
func GetAllProjects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mongo := db.New()
	defer mongo.Close()

	p := mongo.FindAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}
