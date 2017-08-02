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
	util.LogPwd()
	os.Chdir(util.BuildPath(p.Name))
	util.LogPwd()

	out, err := exec.Command("./gradlew", "javadoc").CombinedOutput()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("%s", out)
	log.Println("Finished building javadocs")

	log.Println("Copying javadocs to http server root")
	d := fmt.Sprintf("./../../public/%s", p.Name)
	errors.Handler("Remove old javadocs: ", os.RemoveAll(d))
	errors.Handler("Copy new javadocs: ", util.CopyDir("./javadocs", d))
	errors.Handler("Go back home: ", os.Chdir("./../../"))
	util.LogPwd()

	return nil
}
