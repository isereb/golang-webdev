package main

import (
	"./rest"
	"./services"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/isereb/golang-webdev/logging/util"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port := flag.Int("port", 8081, "-port=8081")
	flag.Parse()

	util.Info("Starting a program on port", *port)

	svc := make(map[string]services.BaseService, 1)
	svc["make_service"] = services.BaseService{
		Name: "make_service",
	}

	handler := rest.BaseHandler{
		Services:   &svc,
		HandlerLog: log.New(os.Stdout, "HANDLER\t", log.LUTC),
	}

	r := mux.NewRouter()
	r.HandleFunc("/makes", handler.GetAllMakes).Methods("GET")
	r.HandleFunc("/make/{make}", handler.GetMake).Methods("GET")
	r.HandleFunc("/make/{make}/models", handler.GetAllModels).Methods("GET")
	r.HandleFunc("/make/{make}/model/{model}", handler.GetModel).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + fmt.Sprint(*port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
