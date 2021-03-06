// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceTodoItemHandlerFunc turns a function with the right signature into a replace todo item handler
type ReplaceTodoItemHandlerFunc func(ReplaceTodoItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceTodoItemHandlerFunc) Handle(params ReplaceTodoItemParams) middleware.Responder {
	return fn(params)
}

// ReplaceTodoItemHandler interface for that can handle valid replace todo item params
type ReplaceTodoItemHandler interface {
	Handle(ReplaceTodoItemParams) middleware.Responder
}

// NewReplaceTodoItem creates a new http.Handler for the replace todo item operation
func NewReplaceTodoItem(ctx *middleware.Context, handler ReplaceTodoItemHandler) *ReplaceTodoItem {
	return &ReplaceTodoItem{Context: ctx, Handler: handler}
}

/*ReplaceTodoItem swagger:route PUT /{id} todo replaceTodoItem

replace

replace a single item

*/
type ReplaceTodoItem struct {
	Context *middleware.Context
	Handler ReplaceTodoItemHandler
}

func (o *ReplaceTodoItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceTodoItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
