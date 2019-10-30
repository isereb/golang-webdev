package main

import "github.com/isereb/golang-webdev/logging/util"

func main() {
	util.InfoLogger.Println("HELLOOOOOO")
	util.DebugLogger.Println("HELLOOOOOO from debug")
	util.WarnLogger.Println("oh no from warn")
}
