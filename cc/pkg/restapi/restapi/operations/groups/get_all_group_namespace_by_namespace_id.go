// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllGroupNamespaceByNamespaceIDHandlerFunc turns a function with the right signature into a get all group namespace by namespace Id handler
type GetAllGroupNamespaceByNamespaceIDHandlerFunc func(GetAllGroupNamespaceByNamespaceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllGroupNamespaceByNamespaceIDHandlerFunc) Handle(params GetAllGroupNamespaceByNamespaceIDParams) middleware.Responder {
	return fn(params)
}

// GetAllGroupNamespaceByNamespaceIDHandler interface for that can handle valid get all group namespace by namespace Id params
type GetAllGroupNamespaceByNamespaceIDHandler interface {
	Handle(GetAllGroupNamespaceByNamespaceIDParams) middleware.Responder
}

// NewGetAllGroupNamespaceByNamespaceID creates a new http.Handler for the get all group namespace by namespace Id operation
func NewGetAllGroupNamespaceByNamespaceID(ctx *middleware.Context, handler GetAllGroupNamespaceByNamespaceIDHandler) *GetAllGroupNamespaceByNamespaceID {
	return &GetAllGroupNamespaceByNamespaceID{Context: ctx, Handler: handler}
}

/*GetAllGroupNamespaceByNamespaceID swagger:route GET /cc/v1/groups/groupNamespace/{namespaceId} Groups getAllGroupNamespaceByNamespaceId

get GroupNamespace.

Optional extended description in Markdown.

*/
type GetAllGroupNamespaceByNamespaceID struct {
	Context *middleware.Context
	Handler GetAllGroupNamespaceByNamespaceIDHandler
}

func (o *GetAllGroupNamespaceByNamespaceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllGroupNamespaceByNamespaceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
