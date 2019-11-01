package services

import (
	"errors"
	"strconv"
)

type Model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Make *Make  `json:"make"`
}

func (s *BaseService) GetAllModels(markId int) ([]Model, error) {

	for _, v := range makes {
		if v.Id == markId {
			return v.Models, nil
		}
	}

	return nil, errors.New("Failed to find mark with id = " + strconv.Itoa(markId))
}

func (s *BaseService) GetModelById(markId int, modelId int) (Model, error) {

	for _, m := range makes {
		if m.Id == markId {
			for _, md := range m.Models {
				if md.Id == modelId {
					return md, nil
				}
			}
		}
	}

	return Model{}, errors.New("Failed to find model with markId = " + strconv.Itoa(markId) +
		" and modelId = " + strconv.Itoa(modelId))
}
