package service

import (
	"swagger_test/models"
	"swagger_test/restapi/operations/todo"
)


func (s *service) CreateTodoItem(item *todo.CreateTodoItemParams) (models.Item, error) {

	var md models.Item

	//check if item exist
	if _, ok := List[item.Body.ID]; !ok { //add only if it doesn't exist

		md = models.Item{
			Completed:   false,
			Description: item.Body.Description,
			ID:          item.Body.ID,
		}
		List[item.Body.ID] = md
	}

	return md, nil
}
