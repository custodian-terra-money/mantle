// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams creates a new PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams object
// with the default values initialized.
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams() *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	var ()
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithTimeout creates a new PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithTimeout(timeout time.Duration) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	var ()
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams{

		timeout: timeout,
	}
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithContext creates a new PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithContext(ctx context.Context) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	var ()
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams{

		Context: ctx,
	}
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithHTTPClient creates a new PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParamsWithHTTPClient(client *http.Client) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	var ()
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams{
		HTTPClient: client,
	}
}

/*PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams contains all the parameters to send to the API endpoint
for the post msgauth granters granter address grantees grantee address grants msg type revoke operation typically these are written to a http.Request
*/
type PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams struct {

	/*GranteeAddress*/
	GranteeAddress string
	/*GranterAddress*/
	GranterAddress string
	/*MsgType*/
	MsgType string
	/*RevokeRequestBody*/
	RevokeRequestBody *models.RevokeGrantReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithTimeout(timeout time.Duration) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithContext(ctx context.Context) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithHTTPClient(client *http.Client) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGranteeAddress adds the granteeAddress to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithGranteeAddress(granteeAddress string) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetGranteeAddress(granteeAddress)
	return o
}

// SetGranteeAddress adds the granteeAddress to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetGranteeAddress(granteeAddress string) {
	o.GranteeAddress = granteeAddress
}

// WithGranterAddress adds the granterAddress to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithGranterAddress(granterAddress string) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetGranterAddress(granterAddress)
	return o
}

// SetGranterAddress adds the granterAddress to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetGranterAddress(granterAddress string) {
	o.GranterAddress = granterAddress
}

// WithMsgType adds the msgType to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithMsgType(msgType string) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetMsgType(msgType)
	return o
}

// SetMsgType adds the msgType to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetMsgType(msgType string) {
	o.MsgType = msgType
}

// WithRevokeRequestBody adds the revokeRequestBody to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WithRevokeRequestBody(revokeRequestBody *models.RevokeGrantReq) *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams {
	o.SetRevokeRequestBody(revokeRequestBody)
	return o
}

// SetRevokeRequestBody adds the revokeRequestBody to the post msgauth granters granter address grantees grantee address grants msg type revoke params
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) SetRevokeRequestBody(revokeRequestBody *models.RevokeGrantReq) {
	o.RevokeRequestBody = revokeRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeRevokeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param granteeAddress
	if err := r.SetPathParam("granteeAddress", o.GranteeAddress); err != nil {
		return err
	}

	// path param granterAddress
	if err := r.SetPathParam("granterAddress", o.GranterAddress); err != nil {
		return err
	}

	// path param msgType
	if err := r.SetPathParam("msgType", o.MsgType); err != nil {
		return err
	}

	if o.RevokeRequestBody != nil {
		if err := r.SetBodyParam(o.RevokeRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}