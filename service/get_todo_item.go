package service

import (
	"errors"
	"swagger_test/models"
)

func (s *service) FetchTodoItem(id int64) (*models.Item, error) {

	item, ok := List[id]

	if ok{
		return  &item, nil
	}

	return &models.Item{}, errors.New("not found")
}
