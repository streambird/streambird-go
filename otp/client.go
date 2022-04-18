package otp

import (
	"fmt"
	"net/http"

	streambird "github.com/streambird/streambird-go/v1"
	"github.com/streambird/streambird-go/v1/otp/email"
	"github.com/streambird/streambird-go/v1/otp/sms"
)

// Client describes your balance information.
type Client struct {
	C     *streambird.Client
	Email *email.Client
	SMS   *sms.Client
}

const path = "auth/otps"

// Verify a OTP
func (c *Client) Verify(verifyReq *streambird.OTPVerifyRequest) (*streambird.OTPVerifyResponse, error) {

	resp := &streambird.OTPVerifyResponse{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/verify", path), verifyReq); err != nil {
		return nil, err
	}

	return resp, nil
}
