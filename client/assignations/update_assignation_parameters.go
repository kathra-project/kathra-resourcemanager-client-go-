// Code generated by go-swagger; DO NOT EDIT.

package assignations

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

	Assignation "github.com/kathra-project/kathra-core-model-go/models"
)

// NewUpdateAssignationParams creates a new UpdateAssignationParams object
// with the default values initialized.
func NewUpdateAssignationParams() *UpdateAssignationParams {
	var ()
	return &UpdateAssignationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAssignationParamsWithTimeout creates a new UpdateAssignationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAssignationParamsWithTimeout(timeout time.Duration) *UpdateAssignationParams {
	var ()
	return &UpdateAssignationParams{

		timeout: timeout,
	}
}

// NewUpdateAssignationParamsWithContext creates a new UpdateAssignationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAssignationParamsWithContext(ctx context.Context) *UpdateAssignationParams {
	var ()
	return &UpdateAssignationParams{

		Context: ctx,
	}
}

// NewUpdateAssignationParamsWithHTTPClient creates a new UpdateAssignationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAssignationParamsWithHTTPClient(client *http.Client) *UpdateAssignationParams {
	var ()
	return &UpdateAssignationParams{
		HTTPClient: client,
	}
}

/*UpdateAssignationParams contains all the parameters to send to the API endpoint
for the update assignation operation typically these are written to a http.Request
*/
type UpdateAssignationParams struct {

	/*Assignation
	  Assignation object to be updated

	*/
	Assignation Assignation.Assignation
	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update assignation params
func (o *UpdateAssignationParams) WithTimeout(timeout time.Duration) *UpdateAssignationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update assignation params
func (o *UpdateAssignationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update assignation params
func (o *UpdateAssignationParams) WithContext(ctx context.Context) *UpdateAssignationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update assignation params
func (o *UpdateAssignationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update assignation params
func (o *UpdateAssignationParams) WithHTTPClient(client *http.Client) *UpdateAssignationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update assignation params
func (o *UpdateAssignationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignation adds the assignation to the update assignation params
func (o *UpdateAssignationParams) WithAssignation(assignation Assignation.Assignation) *UpdateAssignationParams {
	o.SetAssignation(assignation)
	return o
}

// SetAssignation adds the assignation to the update assignation params
func (o *UpdateAssignationParams) SetAssignation(assignation Assignation.Assignation) {
	o.Assignation = assignation
}

// WithResourceID adds the resourceID to the update assignation params
func (o *UpdateAssignationParams) WithResourceID(resourceID string) *UpdateAssignationParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update assignation params
func (o *UpdateAssignationParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAssignationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Assignation != nil {
		if err := r.SetBodyParam(o.Assignation); err != nil {
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
