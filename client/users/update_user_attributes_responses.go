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

// UpdateUserAttributesReader is a Reader for the UpdateUserAttributes structure.
type UpdateUserAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserAttributesOK creates a UpdateUserAttributesOK with default headers values
func NewUpdateUserAttributesOK() *UpdateUserAttributesOK {
	return &UpdateUserAttributesOK{}
}

/*UpdateUserAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdateUserAttributesOK struct {
	Payload models.User
}

func (o *UpdateUserAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{resourceId}][%d] updateUserAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdateUserAttributesOK) GetPayload() models.User {
	return o.Payload
}

func (o *UpdateUserAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAttributesUnauthorized creates a UpdateUserAttributesUnauthorized with default headers values
func NewUpdateUserAttributesUnauthorized() *UpdateUserAttributesUnauthorized {
	return &UpdateUserAttributesUnauthorized{}
}

/*UpdateUserAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserAttributesUnauthorized struct {
}

func (o *UpdateUserAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /users/{resourceId}][%d] updateUserAttributesUnauthorized ", 401)
}

func (o *UpdateUserAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
