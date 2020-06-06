// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FetchTodoItemsHandlerFunc turns a function with the right signature into a fetch todo items handler
type FetchTodoItemsHandlerFunc func(FetchTodoItemsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchTodoItemsHandlerFunc) Handle(params FetchTodoItemsParams) middleware.Responder {
	return fn(params)
}

// FetchTodoItemsHandler interface for that can handle valid fetch todo items params
type FetchTodoItemsHandler interface {
	Handle(FetchTodoItemsParams) middleware.Responder
}

// NewFetchTodoItems creates a new http.Handler for the fetch todo items operation
func NewFetchTodoItems(ctx *middleware.Context, handler FetchTodoItemsHandler) *FetchTodoItems {
	return &FetchTodoItems{Context: ctx, Handler: handler}
}

/*FetchTodoItems swagger:route GET / todo fetchTodoItems

fetch[]

fetch all tiems

*/
type FetchTodoItems struct {
	Context *middleware.Context
	Handler FetchTodoItemsHandler
}

func (o *FetchTodoItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFetchTodoItemsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
