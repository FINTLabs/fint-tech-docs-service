package db

import (
	"fmt"

	"context"

	"github.com/google/go-github/github"
)

// Project type stores information on the projects we want to build
type Project struct {
	CloneURL      string `json:"git,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	ReadMeURL     string `json:"readme,omitempty"`
	HTMLUrl       string `json:"html,omitempty"`
	MavenURL      string `json:"maven,omitempty"`
	MavenBadge    string `json:"maven_badge,omitempty"`
	Latest        string `json:"latest,omitempty"`
	LatestURL     string `json:"latest_url,omitempty"`
	LatestTime    string `json:"latest_time,omitempty"`
	Lang          string `json:"lang,omitempty"`
	JavaDocs      bool   `json:"java_docs"`
	Bintray       bool   `json:"bintray"`
	Dirty         bool   `json:"dirty"`
	FintCoreModel bool   `json:"fint_core_model"`
}

// Build sets up a the Project based on the github.PushEventRepository object
func (p *Project) Build(r *github.PushEventRepository) {
	p.Dirty = true
	p.Description = r.GetDescription()
	p.CloneURL = r.GetCloneURL()
	p.HTMLUrl = r.GetHTMLURL()
	p.ReadMeURL = fmt.Sprintf("%s#readme", r.GetHTMLURL())
	p.Name = r.GetName()
	p.MavenURL = fmt.Sprintf("https://bintray.com/fint/maven/%s/_latestVersion", r.GetName())
	p.MavenBadge = fmt.Sprintf("https://api.bintray.com/packages/fint/maven/%s/images/download.svg", r.GetName())
	lastestInfo(p, r)
}

func lastestInfo(p *Project, r *github.PushEventRepository) {
	client := github.NewClient(nil)
	ctx := context.Background()
	release, _, err := client.Repositories.GetLatestRelease(ctx, r.Owner.GetName(), r.GetName())
	if err == nil {
		p.Latest = release.GetName()
		p.LatestURL = release.GetHTMLURL()
		p.LatestTime = release.GetPublishedAt().Format("02.01.2006 15:04:05")
	}
}
