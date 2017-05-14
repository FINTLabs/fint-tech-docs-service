package config

import (
	"github.com/jinzhu/configor"
)

// Config ...
type Config struct {
	Port           string `default:"9000"`
	WorkspaceDir   string `default:"./workspace"`
	BuildInternval uint64 `default:"30"`
	DBHost         string `default:"localhost"`
	GitHubSecret   string `default:"topsecret"`
}

// Get returns a Config struct
func Get() Config {
	c := Config{}
	configor.Load(&c, "./config.yml")

	if svc.FileExits(c.GitHubSecret)
	return c
}
