package main

func main() {

	logger := getLogger(logSpec{
		label:    "MYLOG",
		priority: 4999,
	})

	logger.Println("Hi!")

	infoLogger.Println("HELLOOOOOO")
	debugLogger.Println("HELLOOOOOO from debug")
	warnLogger.Println("oh no from warn")

	// to run use
	// go run main.go log4go.go
}
