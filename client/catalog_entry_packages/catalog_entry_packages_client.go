// Code generated by go-swagger; DO NOT EDIT.

package catalog_entry_packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog entry packages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog entry packages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddCatalogEntryPackage adds a new catalogentrypackage
*/
func (a *Client) AddCatalogEntryPackage(params *AddCatalogEntryPackageParams, authInfo runtime.ClientAuthInfoWriter) (*AddCatalogEntryPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCatalogEntryPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCatalogEntryPackage",
		Method:             "POST",
		PathPattern:        "/catalogentrypackages",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddCatalogEntryPackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddCatalogEntryPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addCatalogEntryPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCatalogEntryPackage deletes a registered catalogentrypackage
*/
func (a *Client) DeleteCatalogEntryPackage(params *DeleteCatalogEntryPackageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCatalogEntryPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCatalogEntryPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCatalogEntryPackage",
		Method:             "DELETE",
		PathPattern:        "/catalogentrypackages/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCatalogEntryPackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCatalogEntryPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCatalogEntryPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCatalogEntryPackage retrieves a specific catalogentrypackage object
*/
func (a *Client) GetCatalogEntryPackage(params *GetCatalogEntryPackageParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogEntryPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogEntryPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogEntryPackage",
		Method:             "GET",
		PathPattern:        "/catalogentrypackages/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogEntryPackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogEntryPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogEntryPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCatalogEntryPackages retrieves a list of accessible catalogentrypackages for authenticated user
*/
func (a *Client) GetCatalogEntryPackages(params *GetCatalogEntryPackagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCatalogEntryPackagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogEntryPackagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogEntryPackages",
		Method:             "GET",
		PathPattern:        "/catalogentrypackages",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogEntryPackagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCatalogEntryPackagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogEntryPackages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCatalogEntryPackage fullies update a registered catalogentrypackage
*/
func (a *Client) UpdateCatalogEntryPackage(params *UpdateCatalogEntryPackageParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCatalogEntryPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogEntryPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCatalogEntryPackage",
		Method:             "PUT",
		PathPattern:        "/catalogentrypackages/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCatalogEntryPackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCatalogEntryPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCatalogEntryPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCatalogEntryPackageAttributes partiallies update a registered catalogentrypackage
*/
func (a *Client) UpdateCatalogEntryPackageAttributes(params *UpdateCatalogEntryPackageAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCatalogEntryPackageAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogEntryPackageAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCatalogEntryPackageAttributes",
		Method:             "PATCH",
		PathPattern:        "/catalogentrypackages/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCatalogEntryPackageAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCatalogEntryPackageAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCatalogEntryPackageAttributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
