package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/configor"
)

// Config ...
type Config struct {
	Port           string `default:"9000"`
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

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
