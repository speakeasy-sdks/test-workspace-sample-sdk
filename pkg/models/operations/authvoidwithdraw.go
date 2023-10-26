// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AuthVoidWithdrawRequestBodyVoidAmount struct {
	Currency *string `json:"currency,omitempty"`
	Value    *string `json:"value,omitempty"`
}

func (o *AuthVoidWithdrawRequestBodyVoidAmount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AuthVoidWithdrawRequestBodyVoidAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type AuthVoidWithdrawRequestBody struct {
	AccountToken               *string                                `json:"accountToken,omitempty"`
	MerchantID                 *string                                `json:"merchantId,omitempty"`
	OriginalPartnerReferenceNo *string                                `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string                                `json:"originalReferenceNo,omitempty"`
	PartnerVoidNo              *string                                `json:"partnerVoidNo,omitempty"`
	PublicUserID               *string                                `json:"publicUserId,omitempty"`
	Reason                     *string                                `json:"reason,omitempty"`
	VoidAmount                 *AuthVoidWithdrawRequestBodyVoidAmount `json:"voidAmount,omitempty"`
}

func (o *AuthVoidWithdrawRequestBody) GetAccountToken() *string {
	if o == nil {
		return nil
	}
	return o.AccountToken
}

func (o *AuthVoidWithdrawRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *AuthVoidWithdrawRequestBody) GetOriginalPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.OriginalPartnerReferenceNo
}

func (o *AuthVoidWithdrawRequestBody) GetOriginalReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.OriginalReferenceNo
}

func (o *AuthVoidWithdrawRequestBody) GetPartnerVoidNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerVoidNo
}

func (o *AuthVoidWithdrawRequestBody) GetPublicUserID() *string {
	if o == nil {
		return nil
	}
	return o.PublicUserID
}

func (o *AuthVoidWithdrawRequestBody) GetReason() *string {
	if o == nil {
		return nil
	}
	return o.Reason
}

func (o *AuthVoidWithdrawRequestBody) GetVoidAmount() *AuthVoidWithdrawRequestBodyVoidAmount {
	if o == nil {
		return nil
	}
	return o.VoidAmount
}

type AuthVoidWithdrawRequest struct {
	AuthorizationCustomer *string                      `header:"style=simple,explode=false,name=Authorization-Customer"`
	ChannelID             *string                      `header:"style=simple,explode=false,name=Channel-ID"`
	RequestBody           *AuthVoidWithdrawRequestBody `request:"mediaType=application/json"`
	XClientKey            *string                      `header:"style=simple,explode=false,name=X-Client-Key"`
	XExternalID           *string                      `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature            *string                      `header:"style=simple,explode=false,name=X-Signature"`
	XTimestamp            *string                      `header:"style=simple,explode=false,name=X-Timestamp"`
}

func (o *AuthVoidWithdrawRequest) GetAuthorizationCustomer() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationCustomer
}

func (o *AuthVoidWithdrawRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *AuthVoidWithdrawRequest) GetRequestBody() *AuthVoidWithdrawRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AuthVoidWithdrawRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *AuthVoidWithdrawRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *AuthVoidWithdrawRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *AuthVoidWithdrawRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

type AuthVoidWithdrawResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AuthVoidWithdrawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthVoidWithdrawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthVoidWithdrawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}