//go:generate mockgen -package=mocks -destination=../mocks/service.go swagger_test/service Service

//implements foo
package service

import (
	"swagger_test/models"
	"swagger_test/restapi/operations"
	"swagger_test/restapi/operations/todo"
)


var List map[int64]models.Item = map[int64]models.Item{}

type service struct {
}
type Service interface {
	CreateTodoItem(item *todo.CreateTodoItemParams) (models.Item, error)
	FetchTodoItems() []*models.Item
	FetchTodoItem(id int64) (*models.Item, error)
	DeleteTodoItem(id int64) error
	ReplaceTodoItem(id int64, item *todo.ReplaceTodoItemParams) (models.Item, error)
	Configure(api *operations.SwaggerAPI)
}

func NewService() Service {

	svc := &service{}

	return svc
}
