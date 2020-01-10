// Code generated by go-swagger; DO NOT EDIT.

package source_repositories

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

// NewUpdateSourceRepositoryParams creates a new UpdateSourceRepositoryParams object
// with the default values initialized.
func NewUpdateSourceRepositoryParams() *UpdateSourceRepositoryParams {
	var ()
	return &UpdateSourceRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSourceRepositoryParamsWithTimeout creates a new UpdateSourceRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSourceRepositoryParamsWithTimeout(timeout time.Duration) *UpdateSourceRepositoryParams {
	var ()
	return &UpdateSourceRepositoryParams{

		timeout: timeout,
	}
}

// NewUpdateSourceRepositoryParamsWithContext creates a new UpdateSourceRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSourceRepositoryParamsWithContext(ctx context.Context) *UpdateSourceRepositoryParams {
	var ()
	return &UpdateSourceRepositoryParams{

		Context: ctx,
	}
}

// NewUpdateSourceRepositoryParamsWithHTTPClient creates a new UpdateSourceRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSourceRepositoryParamsWithHTTPClient(client *http.Client) *UpdateSourceRepositoryParams {
	var ()
	return &UpdateSourceRepositoryParams{
		HTTPClient: client,
	}
}

/*UpdateSourceRepositoryParams contains all the parameters to send to the API endpoint
for the update source repository operation typically these are written to a http.Request
*/
type UpdateSourceRepositoryParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string
	/*Sourcerepository
	  SourceRepository object to be updated

	*/
	Sourcerepository models.SourceRepository

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update source repository params
func (o *UpdateSourceRepositoryParams) WithTimeout(timeout time.Duration) *UpdateSourceRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update source repository params
func (o *UpdateSourceRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update source repository params
func (o *UpdateSourceRepositoryParams) WithContext(ctx context.Context) *UpdateSourceRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update source repository params
func (o *UpdateSourceRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update source repository params
func (o *UpdateSourceRepositoryParams) WithHTTPClient(client *http.Client) *UpdateSourceRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update source repository params
func (o *UpdateSourceRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the update source repository params
func (o *UpdateSourceRepositoryParams) WithResourceID(resourceID string) *UpdateSourceRepositoryParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update source repository params
func (o *UpdateSourceRepositoryParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithSourcerepository adds the sourcerepository to the update source repository params
func (o *UpdateSourceRepositoryParams) WithSourcerepository(sourcerepository models.SourceRepository) *UpdateSourceRepositoryParams {
	o.SetSourcerepository(sourcerepository)
	return o
}

// SetSourcerepository adds the sourcerepository to the update source repository params
func (o *UpdateSourceRepositoryParams) SetSourcerepository(sourcerepository models.SourceRepository) {
	o.Sourcerepository = sourcerepository
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSourceRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if o.Sourcerepository != nil {
		if err := r.SetBodyParam(o.Sourcerepository); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
