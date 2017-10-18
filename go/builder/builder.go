package builder // import "github.com/FINTprosjektet/fint-tech-docs-service/builder" 

import (
	"log"

	"github.com/FINTprosjektet/fint-tech-docs-service/config"
	"github.com/FINTprosjektet/fint-tech-docs-service/db"
	"github.com/FINTprosjektet/fint-tech-docs-service/util"
	"github.com/jasonlvhit/gocron"
)

// Builder ...
type Builder struct{}

func buildDirtyJavaDocs() {
	mongo := db.New()
	defer mongo.Close()

	util.CleanWorkspace()

	log.Println("Start building documentation")
	ps := mongo.FindDirtyJavaDocs()
	for i := 0; i < len(ps); i++ {
		if ps[i].JavaDocs {
			log.Printf("Building docs for %s\n", ps[i].Name)
			BuildJavaDocs(&ps[i])
		} else {
			log.Printf("%s is not a JavaDoc prosject.", ps[i].Name)
		}
		mongo.Clean(&ps[i])
	}
	log.Println("Finished building documentation")

}

// BuildAllJavaDocs forces a build of all docs in the database
func (b *Builder) BuildAllJavaDocs() {
	mongo := db.New()
	defer mongo.Close()

	util.CleanWorkspace()

	log.Println("Start building documentation")
	ps := mongo.FindAll()
	for i := 0; i < len(ps); i++ {
		if ps[i].JavaDocs {
			log.Printf("Building docs for %s\n", ps[i].Name)
			BuildJavaDocs(&ps[i])
		} else {
			log.Printf("%s is not a JavaDoc prosject.", ps[i].Name)
		}
	}
	log.Println("Finished building documentation")
}

// New ...
func New() *Builder {
	return &Builder{}
}

// Start ...
func (b *Builder) Start() {
	log.Println("Starting documentation builder")

	c := config.Get()
	gocron.Every(c.BuildInternval).Seconds().Do(buildDirtyJavaDocs)
	<-gocron.Start()
}
