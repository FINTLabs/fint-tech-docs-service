package utils

import (
	"log"
	"os"
	"strings"
)

// LogPwd logs the current working directory
func LogPwd() {
	dir, _ := os.Getwd()
	log.Printf("Working directory: %s", strings.Replace(dir, " ", "\\ ", -1))
}
