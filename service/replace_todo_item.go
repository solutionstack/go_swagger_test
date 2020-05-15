package service

import (
	"errors"
	"swagger_test/models"
	"swagger_test/restapi/operations/todo"
)

func (s *service) ReplaceTodoItem(id int64, item *todo.ReplaceTodoItemParams) (models.Item, error) {
	var md models.Item

	//check if item exist
	if _, ok := List[id]; !ok {
		return md, errors.New("404")
	}

	//delete existing
	delete(List, id)

	md = models.Item{
		Completed:   false,
		Description: item.Body.Description,
		ID:          item.Body.ID,
	}
	List[item.Body.ID] = md

	return md, nil
}
