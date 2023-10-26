// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type TopupInquiryRequestBodyAmount struct {
	Currency *string `json:"currency,omitempty"`
	Value    *string `json:"value,omitempty"`
}

func (o *TopupInquiryRequestBodyAmount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *TopupInquiryRequestBodyAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type TopupInquiryRequestBody struct {
	AccountToken       *string                        `json:"accountToken,omitempty"`
	Amount             *TopupInquiryRequestBodyAmount `json:"amount,omitempty"`
	MerchantID         *string                        `json:"merchantId,omitempty"`
	PartnerReferenceNo *string                        `json:"partnerReferenceNo,omitempty"`
	PublicUserID       *string                        `json:"publicUserId,omitempty"`
}

func (o *TopupInquiryRequestBody) GetAccountToken() *string {
	if o == nil {
		return nil
	}
	return o.AccountToken
}

func (o *TopupInquiryRequestBody) GetAmount() *TopupInquiryRequestBodyAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TopupInquiryRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *TopupInquiryRequestBody) GetPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerReferenceNo
}

func (o *TopupInquiryRequestBody) GetPublicUserID() *string {
	if o == nil {
		return nil
	}
	return o.PublicUserID
}

type TopupInquiryRequest struct {
	AuthorizationCustomer *string                  `header:"style=simple,explode=false,name=Authorization-Customer"`
	ChannelID             *string                  `header:"style=simple,explode=false,name=Channel-ID"`
	RequestBody           *TopupInquiryRequestBody `request:"mediaType=application/json"`
	XClientKey            *string                  `header:"style=simple,explode=false,name=X-Client-Key"`
	XExternalID           *string                  `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature            *string                  `header:"style=simple,explode=false,name=X-Signature"`
	XTimestamp            *string                  `header:"style=simple,explode=false,name=X-Timestamp"`
}

func (o *TopupInquiryRequest) GetAuthorizationCustomer() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationCustomer
}

func (o *TopupInquiryRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *TopupInquiryRequest) GetRequestBody() *TopupInquiryRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TopupInquiryRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *TopupInquiryRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *TopupInquiryRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *TopupInquiryRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

type TopupInquiryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TopupInquiryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TopupInquiryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TopupInquiryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}