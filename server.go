package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"gopkg.in/rjz/githubhook.v0"
	"gopkg.in/src-d/go-git.v4"
)

// GetBuildPath returns ...
func GetBuildPath(name string) string {
	return "./workspace/" + name + "/"
}

// BuildJavaDocs builds javadocs
func BuildJavaDocs(name string) error {
	os.Chdir(GetBuildPath(name))
	pwd, _ := os.Getwd()
	log.Println(pwd)
	out, err := exec.Command("./gradlew", "javadoc").CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("%s", out)
	log.Println("Finished building javadocs")
	return nil
}

// GitHubWebHook bladi
func GitHubWebHook(w http.ResponseWriter, req *http.Request) {
	secret := []byte("topsecret")
	hook, err := githubhook.Parse(secret, req)

	p := types.Project{}
	p.Description = "hello"
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}
	evt := github.PushEvent{}
	if err := json.Unmarshal(hook.Payload, &evt); err != nil {
		fmt.Println("Invalid JSON?", err)
	}
	log.Printf("Cloning repository %s", *evt.Repo.FullName)
	os.RemoveAll(GetBuildPath(evt.Repo.GetName()))
	r, err := git.PlainClone(GetBuildPath(evt.Repo.GetName()), false, &git.CloneOptions{
		URL:      evt.Repo.GetCloneURL(),
		Progress: os.Stdout,
	})
	_, err = r.Head()

	if err != nil {
		log.Printf("Failed to clone %s", evt.Repo.GetCloneURL())
	}

	err = BuildJavaDocs(evt.Repo.GetName())
	if err != nil {
		log.Fatal("Unable to build JavaDocs")
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", GitHubWebHook).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
