//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

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

// NewThingsValidateParams creates a new ThingsValidateParams object
// with the default values initialized.
func NewThingsValidateParams() *ThingsValidateParams {
	var ()
	return &ThingsValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsValidateParamsWithTimeout creates a new ThingsValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsValidateParamsWithTimeout(timeout time.Duration) *ThingsValidateParams {
	var ()
	return &ThingsValidateParams{

		timeout: timeout,
	}
}

// NewThingsValidateParamsWithContext creates a new ThingsValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsValidateParamsWithContext(ctx context.Context) *ThingsValidateParams {
	var ()
	return &ThingsValidateParams{

		Context: ctx,
	}
}

// NewThingsValidateParamsWithHTTPClient creates a new ThingsValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsValidateParamsWithHTTPClient(client *http.Client) *ThingsValidateParams {
	var ()
	return &ThingsValidateParams{
		HTTPClient: client,
	}
}

/*ThingsValidateParams contains all the parameters to send to the API endpoint
for the things validate operation typically these are written to a http.Request
*/
type ThingsValidateParams struct {

	/*Body*/
	Body *models.Thing

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things validate params
func (o *ThingsValidateParams) WithTimeout(timeout time.Duration) *ThingsValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things validate params
func (o *ThingsValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things validate params
func (o *ThingsValidateParams) WithContext(ctx context.Context) *ThingsValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things validate params
func (o *ThingsValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things validate params
func (o *ThingsValidateParams) WithHTTPClient(client *http.Client) *ThingsValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things validate params
func (o *ThingsValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the things validate params
func (o *ThingsValidateParams) WithBody(body *models.Thing) *ThingsValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the things validate params
func (o *ThingsValidateParams) SetBody(body *models.Thing) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
