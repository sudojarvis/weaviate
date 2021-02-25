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

package schema

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

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewSchemaObjectsPropertiesAddParams creates a new SchemaObjectsPropertiesAddParams object
// with the default values initialized.
func NewSchemaObjectsPropertiesAddParams() *SchemaObjectsPropertiesAddParams {
	var ()
	return &SchemaObjectsPropertiesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsPropertiesAddParamsWithTimeout creates a new SchemaObjectsPropertiesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaObjectsPropertiesAddParamsWithTimeout(timeout time.Duration) *SchemaObjectsPropertiesAddParams {
	var ()
	return &SchemaObjectsPropertiesAddParams{

		timeout: timeout,
	}
}

// NewSchemaObjectsPropertiesAddParamsWithContext creates a new SchemaObjectsPropertiesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaObjectsPropertiesAddParamsWithContext(ctx context.Context) *SchemaObjectsPropertiesAddParams {
	var ()
	return &SchemaObjectsPropertiesAddParams{

		Context: ctx,
	}
}

// NewSchemaObjectsPropertiesAddParamsWithHTTPClient creates a new SchemaObjectsPropertiesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaObjectsPropertiesAddParamsWithHTTPClient(client *http.Client) *SchemaObjectsPropertiesAddParams {
	var ()
	return &SchemaObjectsPropertiesAddParams{
		HTTPClient: client,
	}
}

/*SchemaObjectsPropertiesAddParams contains all the parameters to send to the API endpoint
for the schema objects properties add operation typically these are written to a http.Request
*/
type SchemaObjectsPropertiesAddParams struct {

	/*Body*/
	Body *models.Property
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) WithTimeout(timeout time.Duration) *SchemaObjectsPropertiesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) WithContext(ctx context.Context) *SchemaObjectsPropertiesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) WithHTTPClient(client *http.Client) *SchemaObjectsPropertiesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) WithBody(body *models.Property) *SchemaObjectsPropertiesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithClassName adds the className to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) WithClassName(className string) *SchemaObjectsPropertiesAddParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects properties add params
func (o *SchemaObjectsPropertiesAddParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsPropertiesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
