// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetGovProposalsProposalIDReader is a Reader for the GetGovProposalsProposalID structure.
type GetGovProposalsProposalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDOK creates a GetGovProposalsProposalIDOK with default headers values
func NewGetGovProposalsProposalIDOK() *GetGovProposalsProposalIDOK {
	return &GetGovProposalsProposalIDOK{}
}

/*GetGovProposalsProposalIDOK handles this case with default header values.

OK
*/
type GetGovProposalsProposalIDOK struct {
	Payload *models.TextProposal
}

func (o *GetGovProposalsProposalIDOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}][%d] getGovProposalsProposalIdOK  %+v", 200, o.Payload)
}

func (o *GetGovProposalsProposalIDOK) GetPayload() *models.TextProposal {
	return o.Payload
}

func (o *GetGovProposalsProposalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TextProposal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDBadRequest creates a GetGovProposalsProposalIDBadRequest with default headers values
func NewGetGovProposalsProposalIDBadRequest() *GetGovProposalsProposalIDBadRequest {
	return &GetGovProposalsProposalIDBadRequest{}
}

/*GetGovProposalsProposalIDBadRequest handles this case with default header values.

Invalid proposal id
*/
type GetGovProposalsProposalIDBadRequest struct {
}

func (o *GetGovProposalsProposalIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}][%d] getGovProposalsProposalIdBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDInternalServerError creates a GetGovProposalsProposalIDInternalServerError with default headers values
func NewGetGovProposalsProposalIDInternalServerError() *GetGovProposalsProposalIDInternalServerError {
	return &GetGovProposalsProposalIDInternalServerError{}
}

/*GetGovProposalsProposalIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDInternalServerError struct {
}

func (o *GetGovProposalsProposalIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}][%d] getGovProposalsProposalIdInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}