package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FINTProsjektet/fint-tech-docs-service/services"
	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"gopkg.in/rjz/githubhook.v0"
)

func gitHubWebHook(w http.ResponseWriter, req *http.Request) {
	secret := []byte("topsecret")
	hook, err := githubhook.Parse(secret, req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		em := fmt.Sprintf("Failed processing hook: %s", err)
		log.Print(em)
		json.NewEncoder(w).Encode(types.ErrorResponse{Message: em})
		return
	}
	evt := github.PushEvent{}
	if err := json.Unmarshal(hook.Payload, &evt); err != nil {
		fmt.Println("Invalid JSON?", err)
	}
	mongo := svc.NewMongo()
	defer mongo.Close()
	mongo.Save(evt.Repo)

	w.WriteHeader(http.StatusOK)
}

func getAllProjects(w http.ResponseWriter, req *http.Request) {
	mongo := svc.NewMongo()
	defer mongo.Close()

	p := mongo.FindAll()
	log.Print(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", gitHubWebHook).Methods("POST")
	router.HandleFunc("/api/projects", getAllProjects).Methods("GET")

	b := svc.NewBuilder()
	go b.Start()

	log.Fatal(http.ListenAndServe(":12345", router))

}
