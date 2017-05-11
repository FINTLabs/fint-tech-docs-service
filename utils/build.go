package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/FINTProsjektet/fint-tech-docs-service/config"
)

// BuildPath return the path we're building the documentation in.
func BuildPath(name string) string {
	c := config.Get()
	return fmt.Sprintf("%s/%s/", c.WorkspaceDir, name)
}

// CleanWorkspace removes the workspaces directory.
func CleanWorkspace() {
	log.Println("Cleaning up workspace")

	c := config.Get()
	os.RemoveAll(c.WorkspaceDir)
}
