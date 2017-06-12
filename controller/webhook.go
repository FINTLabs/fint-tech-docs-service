package controller

import (
	"log"
	"gopkg.in/rjz/githubhook.v0"
	"encoding/json"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
	"net/http"
	"github.com/FINTProsjektet/fint-tech-docs-service/config"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/FINTProsjektet/fint-tech-docs-service/db"
)

func GitHubWebHook(w http.ResponseWriter, req *http.Request) {

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

		p := db.Project{}
		p.JavaDocs = utils.ParseBool(req.URL.Query().Get("javadocs"))
		p.Bintray = utils.ParseBool(req.URL.Query().Get("bintray"))
		p.FintCoreModel = utils.ParseBool(req.URL.Query().Get("fint_core_model"))
		p.Lang = req.URL.Query().Get("lang")
		p.Build(evt.Repo)

		mongo := db.NewMongo()
		defer mongo.Close()
		mongo.Save(&p)
	}

	w.WriteHeader(http.StatusOK)
}

func errorResponse(e error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	em := fmt.Sprintf("Failed processing hook: %s", e)
	log.Print(em)
	json.NewEncoder(w).Encode(ErrorResponse{Message: em})
	return
}
