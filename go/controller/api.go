package controller // import "github.com/FINTprosjektet/fint-tech-docs-service/controller"

import (
	"encoding/json"
	"github.com/FINTprosjektet/fint-tech-docs-service/builder"
	"github.com/FINTprosjektet/fint-tech-docs-service/db"
	"net/http"
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
