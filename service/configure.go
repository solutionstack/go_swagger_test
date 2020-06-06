package service

import (
	"github.com/go-openapi/runtime/middleware"
	"swagger_test/models"
	"swagger_test/restapi/operations"
	"swagger_test/restapi/operations/todo"
)

func (s *service) Configure(api *operations.SwaggerAPI) {

	api.TodoCreateTodoItemHandler = todo.CreateTodoItemHandlerFunc(func(params todo.CreateTodoItemParams) middleware.Responder {

		//create new item
		if item, err := s.CreateTodoItem(&params); err == nil {
			return todo.NewCreateTodoItemCreated().WithPayload(&item)
		}

		return todo.NewCreateTodoItemInternalServerError()
	})


	api.TodoFetchTodoItemsHandler = todo.FetchTodoItemsHandlerFunc(func(params todo.FetchTodoItemsParams) middleware.Responder {

		items := s.FetchTodoItems()
		return todo.NewFetchTodoItemsOK().WithPayload(items)
	})

	api.TodoFetchTodoItemHandler = todo.FetchTodoItemHandlerFunc(func(params todo.FetchTodoItemParams) middleware.Responder {
		msg := "not found"
		item, err := s.FetchTodoItem(params.ID)
		if err != nil {
			return todo.NewFetchTodoItemNotFound().WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return todo.NewFetchTodoItemOK().WithPayload(item)
	})

	api.TodoReplaceTodoItemHandler = todo.ReplaceTodoItemHandlerFunc(func(params todo.ReplaceTodoItemParams) middleware.Responder {
		msg := "not found"
		item, err := s.ReplaceTodoItem(params.ID, &params)
		if err != nil && err.Error() =="404"{
			return todo.NewReplaceTodoItemNotFound().WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return todo.NewReplaceTodoItemOK().WithPayload(&item)
	})
	api.TodoDeleteTodoItemHandler = todo.DeleteTodoItemHandlerFunc(func(params todo.DeleteTodoItemParams) middleware.Responder {
		msg := "not found"
		 err := s.DeleteTodoItem(params.ID)
		if err != nil && err.Error() =="404"{
			return todo.NewDeleteTodoItemNotFound().WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return todo.NewDeleteTodoItemNoContent()
	})
}
