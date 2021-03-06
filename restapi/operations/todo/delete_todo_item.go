// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTodoItemHandlerFunc turns a function with the right signature into a delete todo item handler
type DeleteTodoItemHandlerFunc func(DeleteTodoItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTodoItemHandlerFunc) Handle(params DeleteTodoItemParams) middleware.Responder {
	return fn(params)
}

// DeleteTodoItemHandler interface for that can handle valid delete todo item params
type DeleteTodoItemHandler interface {
	Handle(DeleteTodoItemParams) middleware.Responder
}

// NewDeleteTodoItem creates a new http.Handler for the delete todo item operation
func NewDeleteTodoItem(ctx *middleware.Context, handler DeleteTodoItemHandler) *DeleteTodoItem {
	return &DeleteTodoItem{Context: ctx, Handler: handler}
}

/*DeleteTodoItem swagger:route DELETE /{id} todo deleteTodoItem

delete

delete a single item

*/
type DeleteTodoItem struct {
	Context *middleware.Context
	Handler DeleteTodoItemHandler
}

func (o *DeleteTodoItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteTodoItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
