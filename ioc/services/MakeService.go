package services

import (
	"errors"
	"fmt"
)

type BaseService struct {
	Name string
}

type Make struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var makes = make([]Make, 0, 3)

func init() {
	makes = append(makes, Make{
		Id:   1,
		Name: "Toyota",
	})

	makes = append(makes, Make{
		Id:   2,
		Name: "Mercedes",
	})

	makes = append(makes, Make{
		Id:   3,
		Name: "Chevrolet",
	})
}

func (s *BaseService) GetAllMakes() []Make {
	return makes
}

func (s *BaseService) GetMakeById(id int) (Make, error) {
	for _, v := range makes {
		if v.Id == id {
			return v, nil
		}
	}
	return Make{}, errors.New(fmt.Sprintf("Make with id %d not found", id))
}
