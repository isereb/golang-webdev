package main

import (
	"fmt"
	log "github.com/isereb/golang-webdev/logging/util"
)

func main() {
	log.Debug("HELLOOOOOO from debug")
	log.Info("HELLOOOOOO")
	log.Warn("oh no from warn")
	fmt.Println()
}
