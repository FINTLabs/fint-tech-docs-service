package db

import (
	"log"

	"github.com/FINTProsjektet/fint-tech-docs-service/config"
	"github.com/FINTProsjektet/fint-tech-docs-service/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Mongo is a class for saving `Projects` to MongoService
type Mongo struct {
	session    *mgo.Session
	collection *mgo.Collection
	err        error
}

// New creates a new MongoService
func New() *Mongo {
	c := config.Get()

	m := &Mongo{}
	var err error
	m.session, err = mgo.Dial(c.DBHost)
	if errors.Handler("Dialing Mongo", err) {
		m.session.SetMode(mgo.Monotonic, true)
		m.collection = m.session.DB("docs").C("project")
		return m
	}

	return nil
}

// Close closes the session to Mongo
func (m *Mongo) Close() {
	m.session.Close()
}

// FindAll ...
func (m *Mongo) FindAll() []Project {
	p := []Project{}
	q := m.collection.Find(bson.M{})
	err := q.All(&p)

	if errors.Handler("FindAll prosjects", err) {
		return p
	}

	return []Project{}

}

// FindDirtyJavaDocs ...
func (m *Mongo) FindDirtyJavaDocs() []Project {
	p := []Project{}
	q := m.collection.Find(bson.M{"dirty": true})
	err := q.All(&p)

	if errors.Handler("Find Dirty", err) {
		return p
	}

	return []Project{}

}

// Clean sets the dirty flag to true
func (m *Mongo) Clean(p *Project) {
	p.Dirty = false
	m.collection.Update(bson.M{"cloneurl": p.CloneURL}, p)
}

// Save inserts the new `Project` or updates it if it exists
func (m *Mongo) Save(p *Project) {
	var err error
	//p := types.Project{}
	//p.Build(r)

	log.Printf("Saving project to MongoDB: %s", p.Name)

	if m.exists(p) {
		err = m.collection.Update(bson.M{"cloneurl": p.CloneURL}, p)
	} else {
		err = m.collection.Insert(p)
	}

	errors.Handler("Saving project", err)
}

func (m *Mongo) exists(p *Project) bool {
	r := Project{}
	err := m.collection.Find(bson.M{"cloneurl": p.CloneURL}).One(&r)
	log.Printf("Check if project exists: %s", p.Name)

	return errors.Handler("Project exists", err)
}
