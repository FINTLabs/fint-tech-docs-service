package svc

import (
	"fmt"
	"os"

	"github.com/jasonlvhit/gocron"
)

// Builder ....
type Builder struct{}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func buildDocs() {
	mongo := NewMongo()
	defer mongo.Close()

	os.RemoveAll("./workspace")
	ps := mongo.FindDirty()
	for i := 0; i < len(ps); i++ {
		fmt.Printf("Building docs for %s\n", ps[i].Name)
		BuildJavaDocs(&ps[i])
		mongo.Clean(&ps[i])
	}

}

// NewBuilder ...
func NewBuilder() *Builder {
	return &Builder{}
}

// Start ...
func (b *Builder) Start() {
	gocron.Every(5).Seconds().Do(buildDocs)
	<-gocron.Start()
}
