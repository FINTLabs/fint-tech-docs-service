package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRouters ...
func SetupRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", GitHubWebHook).Methods("POST")
	router.HandleFunc("/api/projects", GetAllProjects).Methods("GET")
	router.HandleFunc("/api/projects/build", BuildAllProjects).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(notFound)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	return router
}
