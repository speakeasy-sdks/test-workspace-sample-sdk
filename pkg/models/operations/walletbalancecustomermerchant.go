// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type WalletbalanceCustomerMerchantRequestBody struct {
	AccountToken       *string `json:"accountToken,omitempty"`
	MerchantID         *string `json:"merchantId,omitempty"`
	PartnerReferenceNo *string `json:"partnerReferenceNo,omitempty"`
	PublicUserID       *string `json:"publicUserId,omitempty"`
}

func (o *WalletbalanceCustomerMerchantRequestBody) GetAccountToken() *string {
	if o == nil {
		return nil
	}
	return o.AccountToken
}

func (o *WalletbalanceCustomerMerchantRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *WalletbalanceCustomerMerchantRequestBody) GetPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerReferenceNo
}

func (o *WalletbalanceCustomerMerchantRequestBody) GetPublicUserID() *string {
	if o == nil {
		return nil
	}
	return o.PublicUserID
}

type WalletbalanceCustomerMerchantRequest struct {
	AuthorizationCustomer *string                                   `header:"style=simple,explode=false,name=Authorization-Customer"`
	ChannelID             *string                                   `header:"style=simple,explode=false,name=CHANNEL-ID"`
	RequestBody           *WalletbalanceCustomerMerchantRequestBody `request:"mediaType=application/json"`
	XClientKey            *string                                   `header:"style=simple,explode=false,name=X-CLIENT-KEY"`
	XExternalID           *string                                   `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature            *string                                   `header:"style=simple,explode=false,name=X-SIGNATURE"`
	XTimestamp            *string                                   `header:"style=simple,explode=false,name=X-TIMESTAMP"`
}

func (o *WalletbalanceCustomerMerchantRequest) GetAuthorizationCustomer() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationCustomer
}

func (o *WalletbalanceCustomerMerchantRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *WalletbalanceCustomerMerchantRequest) GetRequestBody() *WalletbalanceCustomerMerchantRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *WalletbalanceCustomerMerchantRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *WalletbalanceCustomerMerchantRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *WalletbalanceCustomerMerchantRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *WalletbalanceCustomerMerchantRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

type WalletbalanceCustomerMerchantResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *WalletbalanceCustomerMerchantResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WalletbalanceCustomerMerchantResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WalletbalanceCustomerMerchantResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
