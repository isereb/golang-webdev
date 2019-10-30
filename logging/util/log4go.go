package util

import (
	"log"
	"os"
)

type Logger string

var debugLogger = getLogger(logSpec{
	label:    "DEBUG",
	priority: 5000,
})

var infoLogger = getLogger(logSpec{
	label:    "INFO",
	priority: 10000,
})

var warnLogger = getLogger(logSpec{
	label:    "WARN",
	priority: 20000,
})

type logSpec struct {
	label    string
	priority int
}

func getLogger(level logSpec) *log.Logger {
	return log.New(os.Stdin, level.label+"\t", log.LstdFlags)
}
