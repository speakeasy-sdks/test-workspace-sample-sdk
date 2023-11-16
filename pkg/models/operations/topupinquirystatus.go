// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type TopupInquiryStatusRequestBody struct {
	MerchantID                 *string `json:"merchantId,omitempty"`
	OriginalPartnerReferenceNo *string `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	PublicUserID               *string `json:"publicUserId,omitempty"`
	ServiceCode                *string `json:"serviceCode,omitempty"`
}

func (o *TopupInquiryStatusRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *TopupInquiryStatusRequestBody) GetOriginalPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.OriginalPartnerReferenceNo
}

func (o *TopupInquiryStatusRequestBody) GetOriginalReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.OriginalReferenceNo
}

func (o *TopupInquiryStatusRequestBody) GetPublicUserID() *string {
	if o == nil {
		return nil
	}
	return o.PublicUserID
}

func (o *TopupInquiryStatusRequestBody) GetServiceCode() *string {
	if o == nil {
		return nil
	}
	return o.ServiceCode
}

type TopupInquiryStatusRequest struct {
	AuthorizationCustomer *string                        `header:"style=simple,explode=false,name=Authorization-Customer"`
	ChannelID             *string                        `header:"style=simple,explode=false,name=Channel-ID"`
	RequestBody           *TopupInquiryStatusRequestBody `request:"mediaType=application/json"`
	XClientKey            *string                        `header:"style=simple,explode=false,name=X-Client-Key"`
	XExternalID           *string                        `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature            *string                        `header:"style=simple,explode=false,name=X-Signature"`
	XTimestamp            *string                        `header:"style=simple,explode=false,name=X-Timestamp"`
}

func (o *TopupInquiryStatusRequest) GetAuthorizationCustomer() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationCustomer
}

func (o *TopupInquiryStatusRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *TopupInquiryStatusRequest) GetRequestBody() *TopupInquiryStatusRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TopupInquiryStatusRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *TopupInquiryStatusRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *TopupInquiryStatusRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *TopupInquiryStatusRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

type TopupInquiryStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TopupInquiryStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TopupInquiryStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TopupInquiryStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
