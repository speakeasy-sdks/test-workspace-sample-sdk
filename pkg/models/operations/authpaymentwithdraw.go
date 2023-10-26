// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AuthPaymentWithdrawRequestBodyAmount struct {
	Currency *string `json:"currency,omitempty"`
	Value    *string `json:"value,omitempty"`
}

func (o *AuthPaymentWithdrawRequestBodyAmount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AuthPaymentWithdrawRequestBodyAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type AuthPaymentWithdrawRequestBody struct {
	AccountToken       *string                               `json:"accountToken,omitempty"`
	Amount             *AuthPaymentWithdrawRequestBodyAmount `json:"amount,omitempty"`
	MerchantID         *string                               `json:"merchantId,omitempty"`
	PartnerReferenceNo *string                               `json:"partnerReferenceNo,omitempty"`
	PublicUserID       *string                               `json:"publicUserId,omitempty"`
	Title              *string                               `json:"title,omitempty"`
}

func (o *AuthPaymentWithdrawRequestBody) GetAccountToken() *string {
	if o == nil {
		return nil
	}
	return o.AccountToken
}

func (o *AuthPaymentWithdrawRequestBody) GetAmount() *AuthPaymentWithdrawRequestBodyAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AuthPaymentWithdrawRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *AuthPaymentWithdrawRequestBody) GetPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerReferenceNo
}

func (o *AuthPaymentWithdrawRequestBody) GetPublicUserID() *string {
	if o == nil {
		return nil
	}
	return o.PublicUserID
}

func (o *AuthPaymentWithdrawRequestBody) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type AuthPaymentWithdrawRequest struct {
	AuthorizationCustomer *string                         `header:"style=simple,explode=false,name=Authorization-Customer"`
	ChannelID             *string                         `header:"style=simple,explode=false,name=Channel-ID"`
	RequestBody           *AuthPaymentWithdrawRequestBody `request:"mediaType=application/json"`
	XClientKey            *string                         `header:"style=simple,explode=false,name=X-Client-Key"`
	XExternalID           *string                         `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature            *string                         `header:"style=simple,explode=false,name=X-Signature"`
	XTimestamp            *string                         `header:"style=simple,explode=false,name=X-Timestamp"`
}

func (o *AuthPaymentWithdrawRequest) GetAuthorizationCustomer() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationCustomer
}

func (o *AuthPaymentWithdrawRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *AuthPaymentWithdrawRequest) GetRequestBody() *AuthPaymentWithdrawRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AuthPaymentWithdrawRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *AuthPaymentWithdrawRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *AuthPaymentWithdrawRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *AuthPaymentWithdrawRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

type AuthPaymentWithdrawResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AuthPaymentWithdrawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthPaymentWithdrawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthPaymentWithdrawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}