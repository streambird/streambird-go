package main

import (
	"fmt"

	"github.com/streambird/streambird-go/v1"
	client "github.com/streambird/streambird-go/v1/client"
)

func main() {
	cl := client.New("sk_live_bMv4mqNkcEzYwNSy6JtIPzmatZDQrrULEQwgfa25utwSvwsI")
	// createReq := &streambird.MagicLinkCreateRequest{
	// 	UserID: "db34c9d2-80cb-4898-bfd4-d0e42fd12877",
	// }
	// client.SetDebugLog(nil)
	// resp, err := client.MagicLinks.Create(createReq)
	// fmt.Println("response", resp)
	// fmt.Println("err", err)

	verifyReq := &streambird.OTPVerifyRequest{
		OTP: "123456",
	}

	resp, err := cl.OTPs.Verify(verifyReq)

	fmt.Println("response", resp)
	fmt.Println("err", err)
}
