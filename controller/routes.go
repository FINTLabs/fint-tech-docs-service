package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRouters ...
func SetupRouters() *mux.Router {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public"))
	router.HandleFunc("/webhook", GitHubWebHook).Methods("POST")
	router.HandleFunc("/api/projects", GetAllProjects).Methods("GET")
	router.HandleFunc("/api/projects/build", BuildAllProjects).Methods("POST")
	router.Handle("/", fs)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}
