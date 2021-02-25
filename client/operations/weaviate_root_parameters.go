//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewWeaviateRootParams creates a new WeaviateRootParams object
// with the default values initialized.
func NewWeaviateRootParams() *WeaviateRootParams {

	return &WeaviateRootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateRootParamsWithTimeout creates a new WeaviateRootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateRootParamsWithTimeout(timeout time.Duration) *WeaviateRootParams {

	return &WeaviateRootParams{

		timeout: timeout,
	}
}

// NewWeaviateRootParamsWithContext creates a new WeaviateRootParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateRootParamsWithContext(ctx context.Context) *WeaviateRootParams {

	return &WeaviateRootParams{

		Context: ctx,
	}
}

// NewWeaviateRootParamsWithHTTPClient creates a new WeaviateRootParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateRootParamsWithHTTPClient(client *http.Client) *WeaviateRootParams {

	return &WeaviateRootParams{
		HTTPClient: client,
	}
}

/*WeaviateRootParams contains all the parameters to send to the API endpoint
for the weaviate root operation typically these are written to a http.Request
*/
type WeaviateRootParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate root params
func (o *WeaviateRootParams) WithTimeout(timeout time.Duration) *WeaviateRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate root params
func (o *WeaviateRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate root params
func (o *WeaviateRootParams) WithContext(ctx context.Context) *WeaviateRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate root params
func (o *WeaviateRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate root params
func (o *WeaviateRootParams) WithHTTPClient(client *http.Client) *WeaviateRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate root params
func (o *WeaviateRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
