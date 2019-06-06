// Code generated by go-swagger; DO NOT EDIT.

package tapi_notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetNotificationSubscriptionServiceListOKCode is the HTTP code returned for type PostOperationsGetNotificationSubscriptionServiceListOK
const PostOperationsGetNotificationSubscriptionServiceListOKCode int = 200

/*PostOperationsGetNotificationSubscriptionServiceListOK Correct response

swagger:response postOperationsGetNotificationSubscriptionServiceListOK
*/
type PostOperationsGetNotificationSubscriptionServiceListOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiNotificationGetNotificationSubscriptionServiceListOutput `json:"body,omitempty"`
}

// NewPostOperationsGetNotificationSubscriptionServiceListOK creates PostOperationsGetNotificationSubscriptionServiceListOK with default headers values
func NewPostOperationsGetNotificationSubscriptionServiceListOK() *PostOperationsGetNotificationSubscriptionServiceListOK {

	return &PostOperationsGetNotificationSubscriptionServiceListOK{}
}

// WithPayload adds the payload to the post operations get notification subscription service list o k response
func (o *PostOperationsGetNotificationSubscriptionServiceListOK) WithPayload(payload *models.TapiNotificationGetNotificationSubscriptionServiceListOutput) *PostOperationsGetNotificationSubscriptionServiceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get notification subscription service list o k response
func (o *PostOperationsGetNotificationSubscriptionServiceListOK) SetPayload(payload *models.TapiNotificationGetNotificationSubscriptionServiceListOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetNotificationSubscriptionServiceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetNotificationSubscriptionServiceListCreatedCode is the HTTP code returned for type PostOperationsGetNotificationSubscriptionServiceListCreated
const PostOperationsGetNotificationSubscriptionServiceListCreatedCode int = 201

/*PostOperationsGetNotificationSubscriptionServiceListCreated No response

swagger:response postOperationsGetNotificationSubscriptionServiceListCreated
*/
type PostOperationsGetNotificationSubscriptionServiceListCreated struct {
}

// NewPostOperationsGetNotificationSubscriptionServiceListCreated creates PostOperationsGetNotificationSubscriptionServiceListCreated with default headers values
func NewPostOperationsGetNotificationSubscriptionServiceListCreated() *PostOperationsGetNotificationSubscriptionServiceListCreated {

	return &PostOperationsGetNotificationSubscriptionServiceListCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetNotificationSubscriptionServiceListCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetNotificationSubscriptionServiceListBadRequestCode is the HTTP code returned for type PostOperationsGetNotificationSubscriptionServiceListBadRequest
const PostOperationsGetNotificationSubscriptionServiceListBadRequestCode int = 400

/*PostOperationsGetNotificationSubscriptionServiceListBadRequest Internal error

swagger:response postOperationsGetNotificationSubscriptionServiceListBadRequest
*/
type PostOperationsGetNotificationSubscriptionServiceListBadRequest struct {
}

// NewPostOperationsGetNotificationSubscriptionServiceListBadRequest creates PostOperationsGetNotificationSubscriptionServiceListBadRequest with default headers values
func NewPostOperationsGetNotificationSubscriptionServiceListBadRequest() *PostOperationsGetNotificationSubscriptionServiceListBadRequest {

	return &PostOperationsGetNotificationSubscriptionServiceListBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetNotificationSubscriptionServiceListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}