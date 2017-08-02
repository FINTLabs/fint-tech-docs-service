package builder

import (
	"log"
	"os"
	"os/exec"

	"fmt"

	"github.com/FINTprosjektet/fint-tech-docs-service/db"
	"github.com/FINTprosjektet/fint-tech-docs-service/errors"
	"github.com/FINTprosjektet/fint-tech-docs-service/git"
	"github.com/FINTprosjektet/fint-tech-docs-service/util"
)

// BuildJavaDocs ...
func BuildJavaDocs(p *db.Project) error {
	g := git.New()

	g.Clone(p)
	dir, _ := os.Getwd()
	gradle := fmt.Sprintf("%s/%s/gradlew", dir, util.BuildPath(p.Name))
	buildGradle := fmt.Sprintf("%s/%s", dir, util.BuildPath(p.Name))
	javadocs := fmt.Sprintf("%s/%s/javadocs", dir,  util.BuildPath(p.Name))

	out, err := exec.Command(gradle, "-p", buildGradle, "javadoc").CombinedOutput()
	if err != nil {
		log.Printf("%s", out)
		log.Printf("Gradle build faild (%s)", err)
		return err
	}
	log.Printf("%s", out)

	log.Println("Copying javadocs to http server root")
	d := fmt.Sprintf("%s/public1/%s", dir, p.Name)
	errors.Handler("Remove old javadocs: ", os.RemoveAll(d))
	errors.Handler("Copy new javadocs: ", util.CopyDir(javadocs, d))

	return nil
}
