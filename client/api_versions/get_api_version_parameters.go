// Code generated by go-swagger; DO NOT EDIT.

package api_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIVersionParams creates a new GetAPIVersionParams object
// with the default values initialized.
func NewGetAPIVersionParams() *GetAPIVersionParams {
	var ()
	return &GetAPIVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIVersionParamsWithTimeout creates a new GetAPIVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIVersionParamsWithTimeout(timeout time.Duration) *GetAPIVersionParams {
	var ()
	return &GetAPIVersionParams{

		timeout: timeout,
	}
}

// NewGetAPIVersionParamsWithContext creates a new GetAPIVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIVersionParamsWithContext(ctx context.Context) *GetAPIVersionParams {
	var ()
	return &GetAPIVersionParams{

		Context: ctx,
	}
}

// NewGetAPIVersionParamsWithHTTPClient creates a new GetAPIVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIVersionParamsWithHTTPClient(client *http.Client) *GetAPIVersionParams {
	var ()
	return &GetAPIVersionParams{
		HTTPClient: client,
	}
}

/*GetAPIVersionParams contains all the parameters to send to the API endpoint
for the get Api version operation typically these are written to a http.Request
*/
type GetAPIVersionParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api version params
func (o *GetAPIVersionParams) WithTimeout(timeout time.Duration) *GetAPIVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api version params
func (o *GetAPIVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api version params
func (o *GetAPIVersionParams) WithContext(ctx context.Context) *GetAPIVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api version params
func (o *GetAPIVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api version params
func (o *GetAPIVersionParams) WithHTTPClient(client *http.Client) *GetAPIVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api version params
func (o *GetAPIVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the get Api version params
func (o *GetAPIVersionParams) WithResourceID(resourceID string) *GetAPIVersionParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get Api version params
func (o *GetAPIVersionParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}