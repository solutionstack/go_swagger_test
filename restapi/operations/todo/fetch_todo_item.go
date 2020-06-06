// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FetchTodoItemHandlerFunc turns a function with the right signature into a fetch todo item handler
type FetchTodoItemHandlerFunc func(FetchTodoItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchTodoItemHandlerFunc) Handle(params FetchTodoItemParams) middleware.Responder {
	return fn(params)
}

// FetchTodoItemHandler interface for that can handle valid fetch todo item params
type FetchTodoItemHandler interface {
	Handle(FetchTodoItemParams) middleware.Responder
}

// NewFetchTodoItem creates a new http.Handler for the fetch todo item operation
func NewFetchTodoItem(ctx *middleware.Context, handler FetchTodoItemHandler) *FetchTodoItem {
	return &FetchTodoItem{Context: ctx, Handler: handler}
}

/*FetchTodoItem swagger:route GET /{id} todo fetchTodoItem

Get

get a single item

*/
type FetchTodoItem struct {
	Context *middleware.Context
	Handler FetchTodoItemHandler
}

func (o *FetchTodoItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFetchTodoItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
