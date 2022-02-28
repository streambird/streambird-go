package main

import (
	"fmt"

	streambird "github.com/streambird/streambird-go/v1"
	client "github.com/streambird/streambird-go/v1/client"
)

func main() {
	client := client.New("sk_live_bMv4mqNkcEzYwNSy6JtIPzmatZDQrrULEQwgfa25utwSvwsI")
	createReq := &streambird.MagicLinkCreateReq{
		UserID: "db34c9d2-80cb-4898-bfd4-d0e42fd12877",
	}
	client.SetDebugLog(nil)
	resp, err := client.MagicLinks.Create(createReq)
	fmt.Println("response", resp)
	fmt.Println("err", err)
}
