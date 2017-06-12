package git

import (
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4"

	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
	"github.com/FINTProsjektet/fint-tech-docs-service/db"
)

// Git ...
type Git struct{}

// New returns an instance of Git
func New() *Git {
	return &Git{}
}

// Clone ...
func (g *Git) Clone(p *db.Project) {

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
