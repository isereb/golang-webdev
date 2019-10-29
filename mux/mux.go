package main

import (
	"github.com/isereb/golang-webdev/mux/controller"
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

	go serveMercedes(mercedes)
	serveToyota(toyota)
}

func serveMercedes(mercedes controller.Mercedes) {
	err := http.ListenAndServe(":11962", mercedes)
	if err != nil {
		log.Println("Could not start the mercedes http server: ", err)
	}
}

func serveToyota(toyota controller.Toyota) {
	err := http.ListenAndServe(":11937", toyota)
	if err != nil {
		log.Println("Could not start the toyota  http server: ", err)
	}
}
