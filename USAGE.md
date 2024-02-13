<!-- Start SDK Example Usage [usage] -->
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
		ChannelID:   testworkspacesamplesdk.String("95221"),
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