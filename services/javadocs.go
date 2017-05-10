package svc

import (
	"log"
	"os"
	"os/exec"

	"fmt"

	"github.com/FINTProsjektet/fint-tech-docs-service/types"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
)

func buildPath(name string) string {
	return "./workspace/" + name + "/"
}

// BuildJavaDocs ...
func BuildJavaDocs(p *types.Project) error {
	g := NewGit()

	g.Clone(p)
	os.Chdir(buildPath(p.Name))

	out, err := exec.Command("./gradlew", "javadoc").CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("%s", out)
	log.Println("Finished building javadocs")

	d := fmt.Sprintf("../../public/%s", p.Name)
	os.RemoveAll(d)
	eh.CopyDir("./javadocs", d)
	os.Chdir("../../")
	return nil
}
