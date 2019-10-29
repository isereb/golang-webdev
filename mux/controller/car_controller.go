package controller

import (
	"fmt"
	"io"
	"net/http"
)

type Car interface {
	Name() string
	Wheels() uint8
}

type Mercedes string

func (m *Mercedes) Name() string {
	return "Mercedes"
}
func (m *Mercedes) Wheels() uint8 {
	return 2
}

type Toyota string

func (m *Toyota) Name() string {
	return "Toyota"
}
func (m *Toyota) Wheels() uint8 {
	return 4
}

func getText(car Car) string {
	return fmt.Sprintf("My favourite Car is %s. It has %d wheels.", car.Name(), string(car.Wheels()))
}

func (car *Toyota) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, getText(car))
}

func (car *Mercedes) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, getText(car))
}
