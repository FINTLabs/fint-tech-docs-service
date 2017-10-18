package controller // import "github.com/FINTprosjektet/fint-tech-docs-service/controller"

import (
	"encoding/json"
	"fmt"
	"github.com/FINTprosjektet/fint-tech-docs-service/config"
	"github.com/FINTprosjektet/fint-tech-docs-service/db"
	"github.com/FINTprosjektet/fint-tech-docs-service/util"
	"github.com/google/go-github/github"
	"gopkg.in/rjz/githubhook.v0"
	"log"
	"net/http"
)

// GitHubWebHook ...
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
		p.JavaDocs = util.ParseBool(req.URL.Query().Get("javadocs"))
		p.Bintray = util.ParseBool(req.URL.Query().Get("bintray"))
		p.FintCoreModel = util.ParseBool(req.URL.Query().Get("fint_core_model"))
		p.Lang = req.URL.Query().Get("lang")
		p.Build(evt.Repo)

		mongo := db.New()
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
