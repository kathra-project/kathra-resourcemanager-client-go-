// Code generated by go-swagger; DO NOT EDIT.

package binary_repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// GetBinaryRepositoryReader is a Reader for the GetBinaryRepository structure.
type GetBinaryRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBinaryRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBinaryRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBinaryRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBinaryRepositoryOK creates a GetBinaryRepositoryOK with default headers values
func NewGetBinaryRepositoryOK() *GetBinaryRepositoryOK {
	return &GetBinaryRepositoryOK{}
}

/*GetBinaryRepositoryOK handles this case with default header values.

Returns the object
*/
type GetBinaryRepositoryOK struct {
	Payload models.BinaryRepository
}

func (o *GetBinaryRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /binaryrepositories/{resourceId}][%d] getBinaryRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetBinaryRepositoryOK) GetPayload() models.BinaryRepository {
	return o.Payload
}

func (o *GetBinaryRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBinaryRepositoryUnauthorized creates a GetBinaryRepositoryUnauthorized with default headers values
func NewGetBinaryRepositoryUnauthorized() *GetBinaryRepositoryUnauthorized {
	return &GetBinaryRepositoryUnauthorized{}
}

/*GetBinaryRepositoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBinaryRepositoryUnauthorized struct {
}

func (o *GetBinaryRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /binaryrepositories/{resourceId}][%d] getBinaryRepositoryUnauthorized ", 401)
}

func (o *GetBinaryRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
