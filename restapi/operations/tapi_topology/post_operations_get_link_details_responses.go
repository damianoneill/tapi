// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetLinkDetailsOKCode is the HTTP code returned for type PostOperationsGetLinkDetailsOK
const PostOperationsGetLinkDetailsOKCode int = 200

/*PostOperationsGetLinkDetailsOK Correct response

swagger:response postOperationsGetLinkDetailsOK
*/
type PostOperationsGetLinkDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiTopologyGetLinkDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetLinkDetailsOK creates PostOperationsGetLinkDetailsOK with default headers values
func NewPostOperationsGetLinkDetailsOK() *PostOperationsGetLinkDetailsOK {

	return &PostOperationsGetLinkDetailsOK{}
}

// WithPayload adds the payload to the post operations get link details o k response
func (o *PostOperationsGetLinkDetailsOK) WithPayload(payload *models.TapiTopologyGetLinkDetailsOutput) *PostOperationsGetLinkDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get link details o k response
func (o *PostOperationsGetLinkDetailsOK) SetPayload(payload *models.TapiTopologyGetLinkDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetLinkDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetLinkDetailsCreatedCode is the HTTP code returned for type PostOperationsGetLinkDetailsCreated
const PostOperationsGetLinkDetailsCreatedCode int = 201

/*PostOperationsGetLinkDetailsCreated No response

swagger:response postOperationsGetLinkDetailsCreated
*/
type PostOperationsGetLinkDetailsCreated struct {
}

// NewPostOperationsGetLinkDetailsCreated creates PostOperationsGetLinkDetailsCreated with default headers values
func NewPostOperationsGetLinkDetailsCreated() *PostOperationsGetLinkDetailsCreated {

	return &PostOperationsGetLinkDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetLinkDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetLinkDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetLinkDetailsBadRequest
const PostOperationsGetLinkDetailsBadRequestCode int = 400

/*PostOperationsGetLinkDetailsBadRequest Internal error

swagger:response postOperationsGetLinkDetailsBadRequest
*/
type PostOperationsGetLinkDetailsBadRequest struct {
}

// NewPostOperationsGetLinkDetailsBadRequest creates PostOperationsGetLinkDetailsBadRequest with default headers values
func NewPostOperationsGetLinkDetailsBadRequest() *PostOperationsGetLinkDetailsBadRequest {

	return &PostOperationsGetLinkDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetLinkDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
