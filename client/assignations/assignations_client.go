// Code generated by go-swagger; DO NOT EDIT.

package assignations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new assignations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for assignations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddAssignation adds a new assignation
*/
func (a *Client) AddAssignation(params *AddAssignationParams, authInfo runtime.ClientAuthInfoWriter) (*AddAssignationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAssignationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addAssignation",
		Method:             "POST",
		PathPattern:        "/assignations",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddAssignationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddAssignationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addAssignation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAssignation deletes a registered assignation
*/
func (a *Client) DeleteAssignation(params *DeleteAssignationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAssignationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAssignationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAssignation",
		Method:             "DELETE",
		PathPattern:        "/assignations/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAssignationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAssignationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAssignation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAssignation retrieves a specific assignation object
*/
func (a *Client) GetAssignation(params *GetAssignationParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignation",
		Method:             "GET",
		PathPattern:        "/assignations/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAssignationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAssignations retrieves a list of accessible assignations for authenticated user
*/
func (a *Client) GetAssignations(params *GetAssignationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignations",
		Method:             "GET",
		PathPattern:        "/assignations",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAssignationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAssignation fullies update a registered assignation
*/
func (a *Client) UpdateAssignation(params *UpdateAssignationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssignationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAssignationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAssignation",
		Method:             "PUT",
		PathPattern:        "/assignations/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateAssignationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAssignationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAssignation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateAssignationAttributes partiallies update a registered assignation
*/
func (a *Client) UpdateAssignationAttributes(params *UpdateAssignationAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssignationAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAssignationAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAssignationAttributes",
		Method:             "PATCH",
		PathPattern:        "/assignations/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateAssignationAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAssignationAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAssignationAttributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
