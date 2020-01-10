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

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// NewUpdateImplementationVersionAttributesParams creates a new UpdateImplementationVersionAttributesParams object
// with the default values initialized.
func NewUpdateImplementationVersionAttributesParams() *UpdateImplementationVersionAttributesParams {
	var ()
	return &UpdateImplementationVersionAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateImplementationVersionAttributesParamsWithTimeout creates a new UpdateImplementationVersionAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateImplementationVersionAttributesParamsWithTimeout(timeout time.Duration) *UpdateImplementationVersionAttributesParams {
	var ()
	return &UpdateImplementationVersionAttributesParams{

		timeout: timeout,
	}
}

// NewUpdateImplementationVersionAttributesParamsWithContext creates a new UpdateImplementationVersionAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateImplementationVersionAttributesParamsWithContext(ctx context.Context) *UpdateImplementationVersionAttributesParams {
	var ()
	return &UpdateImplementationVersionAttributesParams{

		Context: ctx,
	}
}

// NewUpdateImplementationVersionAttributesParamsWithHTTPClient creates a new UpdateImplementationVersionAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateImplementationVersionAttributesParamsWithHTTPClient(client *http.Client) *UpdateImplementationVersionAttributesParams {
	var ()
	return &UpdateImplementationVersionAttributesParams{
		HTTPClient: client,
	}
}

/*UpdateImplementationVersionAttributesParams contains all the parameters to send to the API endpoint
for the update implementation version attributes operation typically these are written to a http.Request
*/
type UpdateImplementationVersionAttributesParams struct {

	/*Implementationversion
	  ImplementationVersion object to be updated

	*/
	Implementationversion models.ImplementationVersion
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) WithTimeout(timeout time.Duration) *UpdateImplementationVersionAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) WithContext(ctx context.Context) *UpdateImplementationVersionAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) WithHTTPClient(client *http.Client) *UpdateImplementationVersionAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImplementationversion adds the implementationversion to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) WithImplementationversion(implementationversion models.ImplementationVersion) *UpdateImplementationVersionAttributesParams {
	o.SetImplementationversion(implementationversion)
	return o
}

// SetImplementationversion adds the implementationversion to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) SetImplementationversion(implementationversion models.ImplementationVersion) {
	o.Implementationversion = implementationversion
}

// WithResourceID adds the resourceID to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) WithResourceID(resourceID string) *UpdateImplementationVersionAttributesParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update implementation version attributes params
func (o *UpdateImplementationVersionAttributesParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateImplementationVersionAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Implementationversion != nil {
		if err := r.SetBodyParam(o.Implementationversion); err != nil {
			return err
		}
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}