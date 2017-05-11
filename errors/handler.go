package errors

import "log"

// Handler returns `true` if there is no error. If there is an error it returns `false` and logs the error.
func Handler(m string, e error) bool {
	if e != nil {
		log.Printf("ERROR: %s -- %s", m, e)
		return false
	}
	return true
}
