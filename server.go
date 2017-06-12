package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FINTProsjektet/fint-tech-docs-service/config"
	"github.com/FINTProsjektet/fint-tech-docs-service/services"
	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"gopkg.in/rjz/githubhook.v0"
)

func errorResponse(e error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	em := fmt.Sprintf("Failed processing hook: %s", e)
	log.Print(em)
	json.NewEncoder(w).Encode(types.ErrorResponse{Message: em})
	return
}

func gitHubWebHook(w http.ResponseWriter, req *http.Request) {

	if req.Header.Get("X-Github-Event") == "push" {

		c := config.Get()
		secret := []byte(c.GithubSecret)
		log.Printf("Secret: %s", secret)

		w.Header().Set("Content-Type", "application/json")

		hook, err := githubhook.Parse(secret, req)
		if err != nil {
			errorResponse(err, w)
		}

		evt := github.PushEvent{}
		if err := json.Unmarshal(hook.Payload, &evt); err != nil {
			errorResponse(err, w)
		}

		p := types.Project{}
		p.JavaDocs = utils.ParseBool(req.URL.Query().Get("javadocs"))
		p.Bintray = utils.ParseBool(req.URL.Query().Get("bintray"))
		p.FintCoreModel = utils.ParseBool(req.URL.Query().Get("fint_core_model"))
		p.Lang = req.URL.Query().Get("lang")
		p.Build(evt.Repo)

		mongo := svc.NewMongo()
		defer mongo.Close()
		mongo.Save(&p)
	}

	w.WriteHeader(http.StatusOK)
}

func buildAllProjects(w http.ResponseWriter, req *http.Request) {
	builder := svc.NewBuilder()
	go builder.BuildAllDocs()
	w.WriteHeader(http.StatusOK)
}

func getAllProjects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mongo := svc.NewMongo()
	defer mongo.Close()

	p := mongo.FindAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func main() {
	c := config.Get()
	log.Printf("Config.Port=%s", c.Port)
	log.Printf("Config.WorkspaceDir=%s", c.WorkspaceDir)
	log.Printf("Config.DBHost=%s", c.DBHost)
	log.Printf("Config.BuildInternval=%d", c.BuildInternval)
	log.Printf("Config.GithubSecret=%s", c.GithubSecret)
	utils.LogPwd()

	utils.CleanWorkspace()

	router := mux.NewRouter()
	router.HandleFunc("/webhook", gitHubWebHook).Methods("POST")
	router.HandleFunc("/api/projects", getAllProjects).Methods("GET")
	router.HandleFunc("/api/projects/build", buildAllProjects).Methods("POST")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	b := svc.NewBuilder()
	go b.Start()

	log.Printf("FINT tech docs service is listening on port %s", c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.Port), router))
}
