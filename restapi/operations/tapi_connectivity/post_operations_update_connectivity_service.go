// Code generated by go-swagger; DO NOT EDIT.

package tapi_connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsUpdateConnectivityServiceHandlerFunc turns a function with the right signature into a post operations update connectivity service handler
type PostOperationsUpdateConnectivityServiceHandlerFunc func(PostOperationsUpdateConnectivityServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsUpdateConnectivityServiceHandlerFunc) Handle(params PostOperationsUpdateConnectivityServiceParams) middleware.Responder {
	return fn(params)
}

// PostOperationsUpdateConnectivityServiceHandler interface for that can handle valid post operations update connectivity service params
type PostOperationsUpdateConnectivityServiceHandler interface {
	Handle(PostOperationsUpdateConnectivityServiceParams) middleware.Responder
}

// NewPostOperationsUpdateConnectivityService creates a new http.Handler for the post operations update connectivity service operation
func NewPostOperationsUpdateConnectivityService(ctx *middleware.Context, handler PostOperationsUpdateConnectivityServiceHandler) *PostOperationsUpdateConnectivityService {
	return &PostOperationsUpdateConnectivityService{Context: ctx, Handler: handler}
}

/*PostOperationsUpdateConnectivityService swagger:route POST /operations/update-connectivity-service/ tapi-connectivity postOperationsUpdateConnectivityService

PostOperationsUpdateConnectivityService post operations update connectivity service API

*/
type PostOperationsUpdateConnectivityService struct {
	Context *middleware.Context
	Handler PostOperationsUpdateConnectivityServiceHandler
}

func (o *PostOperationsUpdateConnectivityService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsUpdateConnectivityServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
