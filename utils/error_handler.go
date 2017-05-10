package eh

import "log"

// ErrorHandler returns `true` if there is no error. If there is an error it returns `false` and logs the error.
func ErrorHandler(e error) bool {
	if e != nil {
		log.Printf("ERROR: %s", e)
		return false
	}
	return true
}
