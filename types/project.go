package types

import (
	"fmt"

	"github.com/google/go-github/github"
)

// Project type stores information on the projects we want to build
type Project struct {
	CloneURL    string `json:"git-url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ReadMeURL   string `json:"readme-url,omitempty"`
	HTMLUrl     string `json:"html-url,omitempty"`
	Dirty       bool   `json:"dirty"`
}

// Build sets up a the Project based on the github.PushEventRepository object
func (p *Project) Build(r *github.PushEventRepository) {
	p.Dirty = true
	p.Description = r.GetDescription()
	p.CloneURL = r.GetCloneURL()
	p.HTMLUrl = r.GetHTMLURL()
	p.ReadMeURL = fmt.Sprintf("%s#readme", r.GetHTMLURL())
	p.Name = r.GetName()
}
