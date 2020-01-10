// Code generated by go-swagger; DO NOT EDIT.

package assignations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// UpdateAssignationReader is a Reader for the UpdateAssignation structure.
type UpdateAssignationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAssignationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAssignationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAssignationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAssignationOK creates a UpdateAssignationOK with default headers values
func NewUpdateAssignationOK() *UpdateAssignationOK {
	return &UpdateAssignationOK{}
}

/*UpdateAssignationOK handles this case with default header values.

Returns the modified object
*/
type UpdateAssignationOK struct {
	Payload models.Assignation
}

func (o *UpdateAssignationOK) Error() string {
	return fmt.Sprintf("[PUT /assignations/{resourceId}][%d] updateAssignationOK  %+v", 200, o.Payload)
}

func (o *UpdateAssignationOK) GetPayload() models.Assignation {
	return o.Payload
}

func (o *UpdateAssignationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAssignationUnauthorized creates a UpdateAssignationUnauthorized with default headers values
func NewUpdateAssignationUnauthorized() *UpdateAssignationUnauthorized {
	return &UpdateAssignationUnauthorized{}
}

/*UpdateAssignationUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateAssignationUnauthorized struct {
}

func (o *UpdateAssignationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /assignations/{resourceId}][%d] updateAssignationUnauthorized ", 401)
}

func (o *UpdateAssignationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}