package main

import (
	"github.com/isereb/golang-webdev/mux/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var mercedes controller.Mercedes
	var toyota controller.Toyota

	mercedesMux := http.NewServeMux()
	mercedesMux.Handle("/car", mercedes)

	toyotaMux := http.NewServeMux()
	toyotaMux.Handle("/car", toyota)

	go serveCar(mercedes)
	serveCar(toyota)
}

func serveCar(car controller.Car) {
	var port int
	switch car.(type) {
	case controller.Mercedes:
		port = 11962 // "1" + year mercedes was founded
	case controller.Toyota:
		port = 11937 // "1" + year toyota was founded
	}

	err := http.ListenAndServe(":"+fmt.Sprint(port), car)
	if err != nil {
		log.Println("Could not start the mercedes http server: ", err)
	}
}
