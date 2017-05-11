package svc

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"

	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
)

// Git ...
type Git struct{}

// NewGit returns an instance of Git
func NewGit() *Git {
	return &Git{}
}

// Clone ...
func (g *Git) Clone(p *types.Project) {

	log.Printf("Cloning repository %s", p.Name)
	r, err := git.PlainClone(utils.BuildPath(p.Name), false, &git.CloneOptions{
		URL:      p.CloneURL,
		Progress: os.Stdout,
	})
	_, err = r.Head()

	if err != nil {
		log.Printf("Failed to clone %s", p.CloneURL)
	}
}
