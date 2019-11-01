package services

import (
	"errors"
	"fmt"
)

type Make struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Models []Model `json:"-"`
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
