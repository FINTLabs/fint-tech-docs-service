package svc

import (
	"log"
	"os"
	"os/exec"

	"fmt"

	"github.com/FINTProsjektet/fint-tech-docs-service/errors"
	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
)

// BuildJavaDocs ...
func BuildJavaDocs(p *types.Project) error {
	g := NewGit()

	g.Clone(p)
	utils.LogPwd()
	os.Chdir(utils.BuildPath(p.Name))
	utils.LogPwd()

	out, err := exec.Command("./gradlew", "javadoc").CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("%s", out)
	log.Println("Finished building javadocs")

	log.Println("Copying javadocs to http server root")
	d := fmt.Sprintf("./../../public/%s", p.Name)
	errors.Handler("Remove old javadocs: ", os.RemoveAll(d))
	errors.Handler("Copy new javadocs: ", utils.CopyDir("./javadocs", d))
	errors.Handler("Go back home: ", os.Chdir("./../../"))
	utils.LogPwd()

	return nil
}
