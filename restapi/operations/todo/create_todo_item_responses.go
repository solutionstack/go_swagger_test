// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_test/models"
)

// CreateTodoItemCreatedCode is the HTTP code returned for type CreateTodoItemCreated
const CreateTodoItemCreatedCode int = 201

/*CreateTodoItemCreated Created

swagger:response createTodoItemCreated
*/
type CreateTodoItemCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewCreateTodoItemCreated creates CreateTodoItemCreated with default headers values
func NewCreateTodoItemCreated() *CreateTodoItemCreated {

	return &CreateTodoItemCreated{}
}

// WithPayload adds the payload to the create todo item created response
func (o *CreateTodoItemCreated) WithPayload(payload *models.Item) *CreateTodoItemCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo item created response
func (o *CreateTodoItemCreated) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoItemCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTodoItemDefault error occured

swagger:response createTodoItemDefault
*/
type CreateTodoItemDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTodoItemDefault creates CreateTodoItemDefault with default headers values
func NewCreateTodoItemDefault(code int) *CreateTodoItemDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTodoItemDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create todo item default response
func (o *CreateTodoItemDefault) WithStatusCode(code int) *CreateTodoItemDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create todo item default response
func (o *CreateTodoItemDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create todo item default response
func (o *CreateTodoItemDefault) WithPayload(payload *models.Error) *CreateTodoItemDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo item default response
func (o *CreateTodoItemDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoItemDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}