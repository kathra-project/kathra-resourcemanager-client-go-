// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// AddGroupReader is a Reader for the AddGroup structure.
type AddGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddGroupOK creates a AddGroupOK with default headers values
func NewAddGroupOK() *AddGroupOK {
	return &AddGroupOK{}
}

/*AddGroupOK handles this case with default header values.

Returns the created object
*/
type AddGroupOK struct {
	Payload models.Group
}

func (o *AddGroupOK) Error() string {
	return fmt.Sprintf("[POST /groups][%d] addGroupOK  %+v", 200, o.Payload)
}

func (o *AddGroupOK) GetPayload() models.Group {
	return o.Payload
}

func (o *AddGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGroupUnauthorized creates a AddGroupUnauthorized with default headers values
func NewAddGroupUnauthorized() *AddGroupUnauthorized {
	return &AddGroupUnauthorized{}
}

/*AddGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type AddGroupUnauthorized struct {
}

func (o *AddGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /groups][%d] addGroupUnauthorized ", 401)
}

func (o *AddGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
