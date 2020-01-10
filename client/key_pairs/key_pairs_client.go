// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new key pairs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for key pairs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddKeyPair adds a new keypair
*/
func (a *Client) AddKeyPair(params *AddKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*AddKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addKeyPair",
		Method:             "POST",
		PathPattern:        "/keypairs",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddKeyPairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addKeyPair: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteKeyPair deletes a registered keypair
*/
func (a *Client) DeleteKeyPair(params *DeleteKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteKeyPair",
		Method:             "DELETE",
		PathPattern:        "/keypairs/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteKeyPairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteKeyPair: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKeyPair retrieves a specific keypair object
*/
func (a *Client) GetKeyPair(params *GetKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*GetKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getKeyPair",
		Method:             "GET",
		PathPattern:        "/keypairs/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKeyPairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getKeyPair: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKeyPairs retrieves a list of accessible keypairs for authenticated user
*/
func (a *Client) GetKeyPairs(params *GetKeyPairsParams, authInfo runtime.ClientAuthInfoWriter) (*GetKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getKeyPairs",
		Method:             "GET",
		PathPattern:        "/keypairs",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKeyPairsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKeyPairsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getKeyPairs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateKeyPair fullies update a registered keypair
*/
func (a *Client) UpdateKeyPair(params *UpdateKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateKeyPair",
		Method:             "PUT",
		PathPattern:        "/keypairs/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateKeyPairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateKeyPair: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateKeyPairAttributes partiallies update a registered keypair
*/
func (a *Client) UpdateKeyPairAttributes(params *UpdateKeyPairAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateKeyPairAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKeyPairAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateKeyPairAttributes",
		Method:             "PATCH",
		PathPattern:        "/keypairs/{resourceId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateKeyPairAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateKeyPairAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateKeyPairAttributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
