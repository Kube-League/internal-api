package utils

import "log"

type Log struct {
	Id string
}

func (l *Log) Error(err error) {
	log.Printf("ID: %s - Error: %s\n", l.Id, err)
}
