package utils

import "log"

func LogError(id string, err error) {
	log.Printf("ID: %s - Error: %s\n", id, err)
}
