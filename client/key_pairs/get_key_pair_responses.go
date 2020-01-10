// Code generated by go-swagger; DO NOT EDIT.

package key_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// GetKeyPairReader is a Reader for the GetKeyPair structure.
type GetKeyPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeyPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeyPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKeyPairUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeyPairOK creates a GetKeyPairOK with default headers values
func NewGetKeyPairOK() *GetKeyPairOK {
	return &GetKeyPairOK{}
}

/*GetKeyPairOK handles this case with default header values.

Returns the object
*/
type GetKeyPairOK struct {
	Payload models.KeyPair
}

func (o *GetKeyPairOK) Error() string {
	return fmt.Sprintf("[GET /keypairs/{resourceId}][%d] getKeyPairOK  %+v", 200, o.Payload)
}

func (o *GetKeyPairOK) GetPayload() models.KeyPair {
	return o.Payload
}

func (o *GetKeyPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeyPairUnauthorized creates a GetKeyPairUnauthorized with default headers values
func NewGetKeyPairUnauthorized() *GetKeyPairUnauthorized {
	return &GetKeyPairUnauthorized{}
}

/*GetKeyPairUnauthorized handles this case with default header values.

Unauthorized
*/
type GetKeyPairUnauthorized struct {
}

func (o *GetKeyPairUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keypairs/{resourceId}][%d] getKeyPairUnauthorized ", 401)
}

func (o *GetKeyPairUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
