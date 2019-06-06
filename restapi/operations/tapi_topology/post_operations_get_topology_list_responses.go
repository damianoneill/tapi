// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetTopologyListOKCode is the HTTP code returned for type PostOperationsGetTopologyListOK
const PostOperationsGetTopologyListOKCode int = 200

/*PostOperationsGetTopologyListOK Correct response

swagger:response postOperationsGetTopologyListOK
*/
type PostOperationsGetTopologyListOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiTopologyGetTopologyListOutput `json:"body,omitempty"`
}

// NewPostOperationsGetTopologyListOK creates PostOperationsGetTopologyListOK with default headers values
func NewPostOperationsGetTopologyListOK() *PostOperationsGetTopologyListOK {

	return &PostOperationsGetTopologyListOK{}
}

// WithPayload adds the payload to the post operations get topology list o k response
func (o *PostOperationsGetTopologyListOK) WithPayload(payload *models.TapiTopologyGetTopologyListOutput) *PostOperationsGetTopologyListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get topology list o k response
func (o *PostOperationsGetTopologyListOK) SetPayload(payload *models.TapiTopologyGetTopologyListOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetTopologyListCreatedCode is the HTTP code returned for type PostOperationsGetTopologyListCreated
const PostOperationsGetTopologyListCreatedCode int = 201

/*PostOperationsGetTopologyListCreated No response

swagger:response postOperationsGetTopologyListCreated
*/
type PostOperationsGetTopologyListCreated struct {
}

// NewPostOperationsGetTopologyListCreated creates PostOperationsGetTopologyListCreated with default headers values
func NewPostOperationsGetTopologyListCreated() *PostOperationsGetTopologyListCreated {

	return &PostOperationsGetTopologyListCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyListCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetTopologyListBadRequestCode is the HTTP code returned for type PostOperationsGetTopologyListBadRequest
const PostOperationsGetTopologyListBadRequestCode int = 400

/*PostOperationsGetTopologyListBadRequest Internal error

swagger:response postOperationsGetTopologyListBadRequest
*/
type PostOperationsGetTopologyListBadRequest struct {
}

// NewPostOperationsGetTopologyListBadRequest creates PostOperationsGetTopologyListBadRequest with default headers values
func NewPostOperationsGetTopologyListBadRequest() *PostOperationsGetTopologyListBadRequest {

	return &PostOperationsGetTopologyListBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
