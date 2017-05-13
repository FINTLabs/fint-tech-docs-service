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

	/*
		DB struct {
			Name     string
			User     string `default:"root"`
			Password string `required:"true" env:"DBPassword"`
			Port     uint   `default:"3306"`
		}
	*/
}

// Get returns a Config struct
func Get() Config {
	c := Config{}
	configor.Load(&c, "./config.yml")

	return c
}
