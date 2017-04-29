/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateEventsListHandlerFunc turns a function with the right signature into a weaviate events list handler
type WeaviateEventsListHandlerFunc func(WeaviateEventsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateEventsListHandlerFunc) Handle(params WeaviateEventsListParams) middleware.Responder {
	return fn(params)
}

// WeaviateEventsListHandler interface for that can handle valid weaviate events list params
type WeaviateEventsListHandler interface {
	Handle(WeaviateEventsListParams) middleware.Responder
}

// NewWeaviateEventsList creates a new http.Handler for the weaviate events list operation
func NewWeaviateEventsList(ctx *middleware.Context, handler WeaviateEventsListHandler) *WeaviateEventsList {
	return &WeaviateEventsList{Context: ctx, Handler: handler}
}

/*WeaviateEventsList swagger:route GET /events events weaviateEventsList

Lists events.

*/
type WeaviateEventsList struct {
	Context *middleware.Context
	Handler WeaviateEventsListHandler
}

func (o *WeaviateEventsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateEventsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}