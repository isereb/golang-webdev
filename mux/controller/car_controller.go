package controller

import (
	"fmt"
	"io"
	"net/http"
)

// Car

type Car interface {
	GetName() string
	GetWheels() int
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}

// Mercedes

type Mercedes struct {
	Name   string
	Wheels int
}

func (car Mercedes) GetName() string {
	return "Mercedes"
}

func (car Mercedes) GetWheels() int {
	return 2
}

// Toyota

type Toyota struct {
	Name   string
	Wheels int
}

func (car Toyota) GetName() string {
	return "Toyota"
}

func (car Toyota) GetWheels() int {
	return 4
}

func getText(car Car) string {
	return fmt.Sprintf("My favourite Car is %s. It has %d wheels.", car.GetName(), car.GetWheels())
}

func (car Mercedes) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, getText(car))
}

func (car Toyota) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, getText(car))
}
