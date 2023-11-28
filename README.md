# github.com/speakeasy-sdks/test-workspace-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/test-workspace-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/test-workspace-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-workspace-sample-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	testworkspacesamplesdk "github.com/speakeasy-sdks/test-workspace-sample-sdk"
	"github.com/speakeasy-sdks/test-workspace-sample-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
	s := testworkspacesamplesdk.New()

	ctx := context.Background()
	res, err := s.AccountBinding(ctx, operations.AccountBindingRequest{
		ChannelID: testworkspacesamplesdk.String("95221"),
		RequestBody: &operations.AccountBindingRequestBody{
			AuthCode:           testworkspacesamplesdk.String("76a345deaccf47d2ac786c1a3184f987"),
			MerchantID:         testworkspacesamplesdk.String("AYOPOP"),
			PartnerReferenceNo: testworkspacesamplesdk.String("20230630A00000000000010000000203"),
		},
		XClientKey:  testworkspacesamplesdk.String("h8XiZaCHAaNIvUh60AQqwYO0hJssGfNt80Gq0LaMriOTUAH"),
		XExternalID: testworkspacesamplesdk.String("41017551351950293184162180797889"),
		XSignature:  testworkspacesamplesdk.String("57e850c5daaa6c8afb60801f9f47245b9ceef63cf76a46c1eb717e5e9174e260ce8dff1fde0a9870139840d081b4ff2c3a6a38bb2ce9df7e4115d2d61071b690957b328fa6dfb29b3305c7e596c96accc4f2515e7a5bae720062606c29b6500979bca96220e838da85c2312647ce837df49f6fa1ccf89c33aa9c46287074f1e70fc20dbada8ebee81177b18b001dabfd4464487c41d3f124178583d152339547e25b5bbbc6dfd4ec3d498e07f70dd1f91e4968c1798578c3a967be7ac0b43fb988c9a36598cba9344a9cbb4f8b0b55d533f73c6966f96f6f29945e28fbdf8a180cf51451a28ac588ba4a94f53c1c6e64977c641daac8fd195157e3fb589be45c"),
		XTimestamp:  testworkspacesamplesdk.String("2023-06-05T09:55:32+07:00"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [WhitelabelEWallet SDK](docs/sdks/whitelabelewallet/README.md)

* [AccountBinding](docs/sdks/whitelabelewallet/README.md#accountbinding) - Account Binding
* [AccountCreation](docs/sdks/whitelabelewallet/README.md#accountcreation) - Account Creation
* [AuthCaptureWithdraw](docs/sdks/whitelabelewallet/README.md#authcapturewithdraw) - Auth Capture - withdraw
* [AuthPaymentWithdraw](docs/sdks/whitelabelewallet/README.md#authpaymentwithdraw) - Auth Payment - withdraw
* [AuthQueryWithdraw](docs/sdks/whitelabelewallet/README.md#authquerywithdraw) - Auth Query - withdraw
* [AuthRefundWithdraw](docs/sdks/whitelabelewallet/README.md#authrefundwithdraw) - Auth Refund - withdraw
* [AuthVoidWithdraw](docs/sdks/whitelabelewallet/README.md#authvoidwithdraw) - Auth Void - withdraw
* [GenerateB2b2cToken](docs/sdks/whitelabelewallet/README.md#generateb2b2ctoken) - Generate B2B2C token
* [GenerateB2bToken](docs/sdks/whitelabelewallet/README.md#generateb2btoken) - Generate B2B Access token
* [GenerateWebview](docs/sdks/whitelabelewallet/README.md#generatewebview) - Generate WebView
* [OtpVerification](docs/sdks/whitelabelewallet/README.md#otpverification) - OTP Verification
* [Topup](docs/sdks/whitelabelewallet/README.md#topup) - TopUp
* [TopupInquiry](docs/sdks/whitelabelewallet/README.md#topupinquiry) - TopUp- Inquiry
* [TopupInquiryStatus](docs/sdks/whitelabelewallet/README.md#topupinquirystatus) - TopUp - Inquiry Status
* [WalletbalanceCustomerMerchant](docs/sdks/whitelabelewallet/README.md#walletbalancecustomermerchant) - WalletBalance - Customer Merchant
* [WalletbalanceMerchant](docs/sdks/whitelabelewallet/README.md#walletbalancemerchant) - WalletBalance - Merchant
<!-- End Available Resources and Operations [operations] -->

<!-- Start Special Types [types] -->
## Special Types
<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	testworkspacesamplesdk "github.com/speakeasy-sdks/test-workspace-sample-sdk"
	"github.com/speakeasy-sdks/test-workspace-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/test-workspace-sample-sdk/pkg/models/sdkerrors"
	"log"
)

func main() {
	s := testworkspacesamplesdk.New()

	ctx := context.Background()
	res, err := s.AccountBinding(ctx, operations.AccountBindingRequest{
		ChannelID: testworkspacesamplesdk.String("95221"),
		RequestBody: &operations.AccountBindingRequestBody{
			AuthCode:           testworkspacesamplesdk.String("76a345deaccf47d2ac786c1a3184f987"),
			MerchantID:         testworkspacesamplesdk.String("AYOPOP"),
			PartnerReferenceNo: testworkspacesamplesdk.String("20230630A00000000000010000000203"),
		},
		XClientKey:  testworkspacesamplesdk.String("h8XiZaCHAaNIvUh60AQqwYO0hJssGfNt80Gq0LaMriOTUAH"),
		XExternalID: testworkspacesamplesdk.String("41017551351950293184162180797889"),
		XSignature:  testworkspacesamplesdk.String("57e850c5daaa6c8afb60801f9f47245b9ceef63cf76a46c1eb717e5e9174e260ce8dff1fde0a9870139840d081b4ff2c3a6a38bb2ce9df7e4115d2d61071b690957b328fa6dfb29b3305c7e596c96accc4f2515e7a5bae720062606c29b6500979bca96220e838da85c2312647ce837df49f6fa1ccf89c33aa9c46287074f1e70fc20dbada8ebee81177b18b001dabfd4464487c41d3f124178583d152339547e25b5bbbc6dfd4ec3d498e07f70dd1f91e4968c1798578c3a967be7ac0b43fb988c9a36598cba9344a9cbb4f8b0b55d533f73c6966f96f6f29945e28fbdf8a180cf51451a28ac588ba4a94f53c1c6e64977c641daac8fd195157e3fb589be45c"),
		XTimestamp:  testworkspacesamplesdk.String("2023-06-05T09:55:32+07:00"),
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://sandbox.api.of.ayoconnect.id` | None |

#### Example

```go
package main

import (
	"context"
	testworkspacesamplesdk "github.com/speakeasy-sdks/test-workspace-sample-sdk"
	"github.com/speakeasy-sdks/test-workspace-sample-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
	s := testworkspacesamplesdk.New(
		testworkspacesamplesdk.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.AccountBinding(ctx, operations.AccountBindingRequest{
		ChannelID: testworkspacesamplesdk.String("95221"),
		RequestBody: &operations.AccountBindingRequestBody{
			AuthCode:           testworkspacesamplesdk.String("76a345deaccf47d2ac786c1a3184f987"),
			MerchantID:         testworkspacesamplesdk.String("AYOPOP"),
			PartnerReferenceNo: testworkspacesamplesdk.String("20230630A00000000000010000000203"),
		},
		XClientKey:  testworkspacesamplesdk.String("h8XiZaCHAaNIvUh60AQqwYO0hJssGfNt80Gq0LaMriOTUAH"),
		XExternalID: testworkspacesamplesdk.String("41017551351950293184162180797889"),
		XSignature:  testworkspacesamplesdk.String("57e850c5daaa6c8afb60801f9f47245b9ceef63cf76a46c1eb717e5e9174e260ce8dff1fde0a9870139840d081b4ff2c3a6a38bb2ce9df7e4115d2d61071b690957b328fa6dfb29b3305c7e596c96accc4f2515e7a5bae720062606c29b6500979bca96220e838da85c2312647ce837df49f6fa1ccf89c33aa9c46287074f1e70fc20dbada8ebee81177b18b001dabfd4464487c41d3f124178583d152339547e25b5bbbc6dfd4ec3d498e07f70dd1f91e4968c1798578c3a967be7ac0b43fb988c9a36598cba9344a9cbb4f8b0b55d533f73c6966f96f6f29945e28fbdf8a180cf51451a28ac588ba4a94f53c1c6e64977c641daac8fd195157e3fb589be45c"),
		XTimestamp:  testworkspacesamplesdk.String("2023-06-05T09:55:32+07:00"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testworkspacesamplesdk "github.com/speakeasy-sdks/test-workspace-sample-sdk"
	"github.com/speakeasy-sdks/test-workspace-sample-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
	s := testworkspacesamplesdk.New(
		testworkspacesamplesdk.WithServerURL("https://sandbox.api.of.ayoconnect.id"),
	)

	ctx := context.Background()
	res, err := s.AccountBinding(ctx, operations.AccountBindingRequest{
		ChannelID: testworkspacesamplesdk.String("95221"),
		RequestBody: &operations.AccountBindingRequestBody{
			AuthCode:           testworkspacesamplesdk.String("76a345deaccf47d2ac786c1a3184f987"),
			MerchantID:         testworkspacesamplesdk.String("AYOPOP"),
			PartnerReferenceNo: testworkspacesamplesdk.String("20230630A00000000000010000000203"),
		},
		XClientKey:  testworkspacesamplesdk.String("h8XiZaCHAaNIvUh60AQqwYO0hJssGfNt80Gq0LaMriOTUAH"),
		XExternalID: testworkspacesamplesdk.String("41017551351950293184162180797889"),
		XSignature:  testworkspacesamplesdk.String("57e850c5daaa6c8afb60801f9f47245b9ceef63cf76a46c1eb717e5e9174e260ce8dff1fde0a9870139840d081b4ff2c3a6a38bb2ce9df7e4115d2d61071b690957b328fa6dfb29b3305c7e596c96accc4f2515e7a5bae720062606c29b6500979bca96220e838da85c2312647ce837df49f6fa1ccf89c33aa9c46287074f1e70fc20dbada8ebee81177b18b001dabfd4464487c41d3f124178583d152339547e25b5bbbc6dfd4ec3d498e07f70dd1f91e4968c1798578c3a967be7ac0b43fb988c9a36598cba9344a9cbb4f8b0b55d533f73c6966f96f6f29945e28fbdf8a180cf51451a28ac588ba4a94f53c1c6e64977c641daac8fd195157e3fb589be45c"),
		XTimestamp:  testworkspacesamplesdk.String("2023-06-05T09:55:32+07:00"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
