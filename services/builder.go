package svc

import (
	"log"

	"github.com/FINTProsjektet/fint-tech-docs-service/config"
	"github.com/FINTProsjektet/fint-tech-docs-service/utils"
	"github.com/jasonlvhit/gocron"
)

// Builder ....
type Builder struct{}

func buildDirtyDocs() {
	mongo := NewMongo()
	defer mongo.Close()

	utils.CleanWorkspace()

	log.Println("Start building documentation")
	ps := mongo.FindDirty()
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

// BuildAllDocs forces a build of all docs in the database
func (b *Builder) BuildAllDocs() {
	mongo := NewMongo()
	defer mongo.Close()

	utils.CleanWorkspace()

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

// NewBuilder ...
func NewBuilder() *Builder {
	return &Builder{}
}

// Start ...
func (b *Builder) Start() {
	log.Println("Starting documentation builder")

	c := config.Get()
	gocron.Every(c.BuildInternval).Seconds().Do(buildDirtyDocs)
	<-gocron.Start()
}
