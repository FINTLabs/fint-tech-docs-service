package config // import "github.com/FINTprosjektet/fint-tech-docs-service/config"

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/configor"
)

// Config ...
type Config struct {
	Port           string `default:"9000"`
	Webroot        string `default:"/var/local/public"`
	WorkspaceDir   string `default:"./workspace"`
	BuildInternval uint64 `default:"30"`
	DBHost         string `default:"localhost"`
	GithubSecret   string `default:"/run/secrets/github_webhook"`
}

// Get returns a Config struct
func Get() Config {
	c := Config{}
	configor.Load(&c, "config.yml")

	if fileExists(c.GithubSecret) {
		log.Println("Using docker secret")
		s, _ := ioutil.ReadFile(c.GithubSecret)
		c.GithubSecret = string(s)
	}

	return c
}

// Dump prints out the configuration.
func Dump() {
	c := Get()
	log.Printf("Config.Port=%s", c.Port)
	log.Printf("Config.Webroot=%s", c.Webroot)
	log.Printf("Config.WorkspaceDir=%s", c.WorkspaceDir)
	log.Printf("Config.DBHost=%s", c.DBHost)
	log.Printf("Config.BuildInternval=%d", c.BuildInternval)
	log.Printf("Config.GithubSecret=%s", c.GithubSecret)
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
