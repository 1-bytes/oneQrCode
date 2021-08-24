package logger

import "log"

// CheckError any errors will be saved to the log.
func CheckError(err error) {
	if err != nil {
		log.Printf("[Error]: %s", err)
	}
}
