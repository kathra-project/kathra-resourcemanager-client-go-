// Code generated by go-swagger; DO NOT EDIT.

package implementation_versions

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

// NewDeleteImplementationVersionParams creates a new DeleteImplementationVersionParams object
// with the default values initialized.
func NewDeleteImplementationVersionParams() *DeleteImplementationVersionParams {
	var ()
	return &DeleteImplementationVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImplementationVersionParamsWithTimeout creates a new DeleteImplementationVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteImplementationVersionParamsWithTimeout(timeout time.Duration) *DeleteImplementationVersionParams {
	var ()
	return &DeleteImplementationVersionParams{

		timeout: timeout,
	}
}

// NewDeleteImplementationVersionParamsWithContext creates a new DeleteImplementationVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteImplementationVersionParamsWithContext(ctx context.Context) *DeleteImplementationVersionParams {
	var ()
	return &DeleteImplementationVersionParams{

		Context: ctx,
	}
}

// NewDeleteImplementationVersionParamsWithHTTPClient creates a new DeleteImplementationVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteImplementationVersionParamsWithHTTPClient(client *http.Client) *DeleteImplementationVersionParams {
	var ()
	return &DeleteImplementationVersionParams{
		HTTPClient: client,
	}
}

/*DeleteImplementationVersionParams contains all the parameters to send to the API endpoint
for the delete implementation version operation typically these are written to a http.Request
*/
type DeleteImplementationVersionParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete implementation version params
func (o *DeleteImplementationVersionParams) WithTimeout(timeout time.Duration) *DeleteImplementationVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete implementation version params
func (o *DeleteImplementationVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete implementation version params
func (o *DeleteImplementationVersionParams) WithContext(ctx context.Context) *DeleteImplementationVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete implementation version params
func (o *DeleteImplementationVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete implementation version params
func (o *DeleteImplementationVersionParams) WithHTTPClient(client *http.Client) *DeleteImplementationVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete implementation version params
func (o *DeleteImplementationVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the delete implementation version params
func (o *DeleteImplementationVersionParams) WithResourceID(resourceID string) *DeleteImplementationVersionParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the delete implementation version params
func (o *DeleteImplementationVersionParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImplementationVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
