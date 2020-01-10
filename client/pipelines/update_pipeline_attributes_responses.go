// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kathra-project/kathra-resourcemanager-client-go/models"
)

// UpdatePipelineAttributesReader is a Reader for the UpdatePipelineAttributes structure.
type UpdatePipelineAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePipelineAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePipelineAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePipelineAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePipelineAttributesOK creates a UpdatePipelineAttributesOK with default headers values
func NewUpdatePipelineAttributesOK() *UpdatePipelineAttributesOK {
	return &UpdatePipelineAttributesOK{}
}

/*UpdatePipelineAttributesOK handles this case with default header values.

Returns the modified object
*/
type UpdatePipelineAttributesOK struct {
	Payload models.Pipeline
}

func (o *UpdatePipelineAttributesOK) Error() string {
	return fmt.Sprintf("[PATCH /pipelines/{resourceId}][%d] updatePipelineAttributesOK  %+v", 200, o.Payload)
}

func (o *UpdatePipelineAttributesOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *UpdatePipelineAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePipelineAttributesUnauthorized creates a UpdatePipelineAttributesUnauthorized with default headers values
func NewUpdatePipelineAttributesUnauthorized() *UpdatePipelineAttributesUnauthorized {
	return &UpdatePipelineAttributesUnauthorized{}
}

/*UpdatePipelineAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdatePipelineAttributesUnauthorized struct {
}

func (o *UpdatePipelineAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /pipelines/{resourceId}][%d] updatePipelineAttributesUnauthorized ", 401)
}

func (o *UpdatePipelineAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
