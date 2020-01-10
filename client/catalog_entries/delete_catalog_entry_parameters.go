// Code generated by go-swagger; DO NOT EDIT.

package catalog_entries

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

// NewDeleteCatalogEntryParams creates a new DeleteCatalogEntryParams object
// with the default values initialized.
func NewDeleteCatalogEntryParams() *DeleteCatalogEntryParams {
	var ()
	return &DeleteCatalogEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogEntryParamsWithTimeout creates a new DeleteCatalogEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCatalogEntryParamsWithTimeout(timeout time.Duration) *DeleteCatalogEntryParams {
	var ()
	return &DeleteCatalogEntryParams{

		timeout: timeout,
	}
}

// NewDeleteCatalogEntryParamsWithContext creates a new DeleteCatalogEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCatalogEntryParamsWithContext(ctx context.Context) *DeleteCatalogEntryParams {
	var ()
	return &DeleteCatalogEntryParams{

		Context: ctx,
	}
}

// NewDeleteCatalogEntryParamsWithHTTPClient creates a new DeleteCatalogEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCatalogEntryParamsWithHTTPClient(client *http.Client) *DeleteCatalogEntryParams {
	var ()
	return &DeleteCatalogEntryParams{
		HTTPClient: client,
	}
}

/*DeleteCatalogEntryParams contains all the parameters to send to the API endpoint
for the delete catalog entry operation typically these are written to a http.Request
*/
type DeleteCatalogEntryParams struct {

	/*ResourceID
	  resource's id

	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete catalog entry params
func (o *DeleteCatalogEntryParams) WithTimeout(timeout time.Duration) *DeleteCatalogEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalog entry params
func (o *DeleteCatalogEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalog entry params
func (o *DeleteCatalogEntryParams) WithContext(ctx context.Context) *DeleteCatalogEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalog entry params
func (o *DeleteCatalogEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalog entry params
func (o *DeleteCatalogEntryParams) WithHTTPClient(client *http.Client) *DeleteCatalogEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalog entry params
func (o *DeleteCatalogEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the delete catalog entry params
func (o *DeleteCatalogEntryParams) WithResourceID(resourceID string) *DeleteCatalogEntryParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the delete catalog entry params
func (o *DeleteCatalogEntryParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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