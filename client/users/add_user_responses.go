// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddUserOK creates a AddUserOK with default headers values
func NewAddUserOK() *AddUserOK {
	return &AddUserOK{}
}

/*AddUserOK handles this case with default header values.

Returns the created object
*/
type AddUserOK struct {
	Payload models.User
}

func (o *AddUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserOK  %+v", 200, o.Payload)
}

func (o *AddUserOK) GetPayload() models.User {
	return o.Payload
}

func (o *AddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserUnauthorized creates a AddUserUnauthorized with default headers values
func NewAddUserUnauthorized() *AddUserUnauthorized {
	return &AddUserUnauthorized{}
}

/*AddUserUnauthorized handles this case with default header values.

Unauthorized
*/
type AddUserUnauthorized struct {
}

func (o *AddUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserUnauthorized ", 401)
}

func (o *AddUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}