package controller // import "github.com/FINTprosjektet/fint-tech-docs-service/controller"

import (
	"net/http"
	"log"
	"os"
)

func router(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v %+v", r.Method, r.URL)
	if r.Method == "POST" {
		if r.URL.Path == "/webhook" {
			GitHubWebHook(w,r)
		} else if r.URL.Path == "/api/projects/build" {
			BuildAllProjects(w,r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else if r.Method == "GET" {
		if r.URL.Path == "/api/projects" {
			GetAllProjects(w,r)
		} else if _, err := os.Stat("./public" + r.URL.Path); err == nil {
			log.Printf("Serving file ./public%s", r.URL.Path)
			http.ServeFile(w, r, "./public" + r.URL.Path)
		} else {
			http.ServeFile(w, r, "./public/index.html")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// SetupRouters ...
func SetupRouters() http.Handler {
	log.Println("Setting up HTTP handler...")
	return http.HandlerFunc(router)
}
