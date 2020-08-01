// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle/lcd/models"
)

// PostStakingDelegatorsDelegatorAddrDelegationsReader is a Reader for the PostStakingDelegatorsDelegatorAddrDelegations structure.
type PostStakingDelegatorsDelegatorAddrDelegationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStakingDelegatorsDelegatorAddrDelegationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStakingDelegatorsDelegatorAddrDelegationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStakingDelegatorsDelegatorAddrDelegationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostStakingDelegatorsDelegatorAddrDelegationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStakingDelegatorsDelegatorAddrDelegationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsOK creates a PostStakingDelegatorsDelegatorAddrDelegationsOK with default headers values
func NewPostStakingDelegatorsDelegatorAddrDelegationsOK() *PostStakingDelegatorsDelegatorAddrDelegationsOK {
	return &PostStakingDelegatorsDelegatorAddrDelegationsOK{}
}

/*PostStakingDelegatorsDelegatorAddrDelegationsOK handles this case with default header values.

OK
*/
type PostStakingDelegatorsDelegatorAddrDelegationsOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsOK) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] postStakingDelegatorsDelegatorAddrDelegationsOK  %+v", 200, o.Payload)
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsBadRequest creates a PostStakingDelegatorsDelegatorAddrDelegationsBadRequest with default headers values
func NewPostStakingDelegatorsDelegatorAddrDelegationsBadRequest() *PostStakingDelegatorsDelegatorAddrDelegationsBadRequest {
	return &PostStakingDelegatorsDelegatorAddrDelegationsBadRequest{}
}

/*PostStakingDelegatorsDelegatorAddrDelegationsBadRequest handles this case with default header values.

Invalid delegator address or delegation request body
*/
type PostStakingDelegatorsDelegatorAddrDelegationsBadRequest struct {
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] postStakingDelegatorsDelegatorAddrDelegationsBadRequest ", 400)
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsUnauthorized creates a PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized with default headers values
func NewPostStakingDelegatorsDelegatorAddrDelegationsUnauthorized() *PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized {
	return &PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized{}
}

/*PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized handles this case with default header values.

Key password is wrong
*/
type PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized struct {
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] postStakingDelegatorsDelegatorAddrDelegationsUnauthorized ", 401)
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsInternalServerError creates a PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError with default headers values
func NewPostStakingDelegatorsDelegatorAddrDelegationsInternalServerError() *PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError {
	return &PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError{}
}

/*PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError struct {
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] postStakingDelegatorsDelegatorAddrDelegationsInternalServerError ", 500)
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostStakingDelegatorsDelegatorAddrDelegationsBody post staking delegators delegator addr delegations body
swagger:model PostStakingDelegatorsDelegatorAddrDelegationsBody
*/
type PostStakingDelegatorsDelegatorAddrDelegationsBody struct {

	// amount
	Amount *models.Coin `json:"amount,omitempty"`

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// delegator address
	DelegatorAddress models.Address `json:"delegator_address,omitempty"`

	// validator address
	ValidatorAddress models.ValidatorAddress `json:"validator_address,omitempty"`
}

// Validate validates this post staking delegators delegator addr delegations body
func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDelegatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValidatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	if o.Amount != nil {
		if err := o.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation" + "." + "amount")
			}
			return err
		}
	}

	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) validateDelegatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.DelegatorAddress) { // not required
		return nil
	}

	if err := o.DelegatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "delegator_address")
		}
		return err
	}

	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) validateValidatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.ValidatorAddress) { // not required
		return nil
	}

	if err := o.ValidatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "validator_address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostStakingDelegatorsDelegatorAddrDelegationsBody) UnmarshalBinary(b []byte) error {
	var res PostStakingDelegatorsDelegatorAddrDelegationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}