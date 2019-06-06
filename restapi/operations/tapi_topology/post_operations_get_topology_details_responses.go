// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetTopologyDetailsOKCode is the HTTP code returned for type PostOperationsGetTopologyDetailsOK
const PostOperationsGetTopologyDetailsOKCode int = 200

/*PostOperationsGetTopologyDetailsOK Correct response

swagger:response postOperationsGetTopologyDetailsOK
*/
type PostOperationsGetTopologyDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiTopologyGetTopologyDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetTopologyDetailsOK creates PostOperationsGetTopologyDetailsOK with default headers values
func NewPostOperationsGetTopologyDetailsOK() *PostOperationsGetTopologyDetailsOK {

	return &PostOperationsGetTopologyDetailsOK{}
}

// WithPayload adds the payload to the post operations get topology details o k response
func (o *PostOperationsGetTopologyDetailsOK) WithPayload(payload *models.TapiTopologyGetTopologyDetailsOutput) *PostOperationsGetTopologyDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get topology details o k response
func (o *PostOperationsGetTopologyDetailsOK) SetPayload(payload *models.TapiTopologyGetTopologyDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetTopologyDetailsCreatedCode is the HTTP code returned for type PostOperationsGetTopologyDetailsCreated
const PostOperationsGetTopologyDetailsCreatedCode int = 201

/*PostOperationsGetTopologyDetailsCreated No response

swagger:response postOperationsGetTopologyDetailsCreated
*/
type PostOperationsGetTopologyDetailsCreated struct {
}

// NewPostOperationsGetTopologyDetailsCreated creates PostOperationsGetTopologyDetailsCreated with default headers values
func NewPostOperationsGetTopologyDetailsCreated() *PostOperationsGetTopologyDetailsCreated {

	return &PostOperationsGetTopologyDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetTopologyDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetTopologyDetailsBadRequest
const PostOperationsGetTopologyDetailsBadRequestCode int = 400

/*PostOperationsGetTopologyDetailsBadRequest Internal error

swagger:response postOperationsGetTopologyDetailsBadRequest
*/
type PostOperationsGetTopologyDetailsBadRequest struct {
}

// NewPostOperationsGetTopologyDetailsBadRequest creates PostOperationsGetTopologyDetailsBadRequest with default headers values
func NewPostOperationsGetTopologyDetailsBadRequest() *PostOperationsGetTopologyDetailsBadRequest {

	return &PostOperationsGetTopologyDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetTopologyDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}