package util

import (
	"log"
	"os"
)

type Logger string

var DebugLogger = getLogger(logSpec{
	label:    "DEBUG",
	priority: 5000,
})

var InfoLogger = getLogger(logSpec{
	label:    "INFO",
	priority: 10000,
})

var WarnLogger = getLogger(logSpec{
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
