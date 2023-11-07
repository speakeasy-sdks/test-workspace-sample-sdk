// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AccountCreationRequestBody struct {
	Email              *string `json:"email,omitempty"`
	MerchantID         *string `json:"merchantId,omitempty"`
	Name               *string `json:"name,omitempty"`
	PartnerReferenceNo *string `json:"partnerReferenceNo,omitempty"`
	PhoneNo            *string `json:"phoneNo,omitempty"`
}

func (o *AccountCreationRequestBody) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AccountCreationRequestBody) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}

func (o *AccountCreationRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountCreationRequestBody) GetPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerReferenceNo
}

func (o *AccountCreationRequestBody) GetPhoneNo() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNo
}

type AccountCreationRequest struct {
	ChannelID   *string                     `header:"style=simple,explode=false,name=Channel-ID"`
	RequestBody *AccountCreationRequestBody `request:"mediaType=application/json"`
	XClientKey  *string                     `header:"style=simple,explode=false,name=X-Client-Key"`
	XExternalID *string                     `header:"style=simple,explode=false,name=X-EXTERNAL-ID"`
	XSignature  *string                     `header:"style=simple,explode=false,name=X-Signature"`
	XTimestamp  *string                     `header:"style=simple,explode=false,name=X-Timestamp"`
}

func (o *AccountCreationRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *AccountCreationRequest) GetRequestBody() *AccountCreationRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AccountCreationRequest) GetXClientKey() *string {
	if o == nil {
		return nil
	}
	return o.XClientKey
}

func (o *AccountCreationRequest) GetXExternalID() *string {
	if o == nil {
		return nil
	}
	return o.XExternalID
}

func (o *AccountCreationRequest) GetXSignature() *string {
	if o == nil {
		return nil
	}
	return o.XSignature
}

func (o *AccountCreationRequest) GetXTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.XTimestamp
}

// AccountCreationResponseBody - Account Creation response
type AccountCreationResponseBody struct {
	PartnerReferenceNo *string `json:"partnerReferenceNo,omitempty"`
	ReferenceNo        *string `json:"referenceNo,omitempty"`
	RegistrationToken  *string `json:"registrationToken,omitempty"`
	ResponseCode       *string `json:"responseCode,omitempty"`
	ResponseMessage    *string `json:"responseMessage,omitempty"`
	ResponseTime       *string `json:"responseTime,omitempty"`
}

func (o *AccountCreationResponseBody) GetPartnerReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.PartnerReferenceNo
}

func (o *AccountCreationResponseBody) GetReferenceNo() *string {
	if o == nil {
		return nil
	}
	return o.ReferenceNo
}

func (o *AccountCreationResponseBody) GetRegistrationToken() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationToken
}

func (o *AccountCreationResponseBody) GetResponseCode() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCode
}

func (o *AccountCreationResponseBody) GetResponseMessage() *string {
	if o == nil {
		return nil
	}
	return o.ResponseMessage
}

func (o *AccountCreationResponseBody) GetResponseTime() *string {
	if o == nil {
		return nil
	}
	return o.ResponseTime
}

type AccountCreationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Account Creation response
	Object *AccountCreationResponseBody
}

func (o *AccountCreationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountCreationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *AccountCreationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountCreationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountCreationResponse) GetObject() *AccountCreationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
