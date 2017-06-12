package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// LogPwd logs the current working directory
func LogPwd() {
	dir, _ := os.Getwd()
	log.Printf("Working directory: %s", strings.Replace(dir, " ", "\\ ", -1))
}

// ParseBool tries to convert a string to bool. If an error occurs it returns false.
func ParseBool(s string) bool {
	log.Printf("param: %s", s)
	b, e := strconv.ParseBool(s)

	if e != nil {
		log.Println("error parsing bool")
		b = false
	}
	return b
}
