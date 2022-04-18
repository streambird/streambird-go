package sms

import (
	"fmt"
	"net/http"

	streambird "github.com/streambird/streambird-go/v1"
)

// Client describes your balance information.
type Client struct {
	C *streambird.Client
}

const path = "auth/otps/sms"

// LoginOrCreate by SMS OTP
func (c *Client) LoginOrCreate(req *streambird.OTPLoginOrCreateBySMSRequest) (*streambird.OTPLoginOrCreateBySMSResponse, error) {

	resp := &streambird.OTPLoginOrCreateBySMSResponse{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/login_or_create", path), req); err != nil {
		return nil, err
	}

	return resp, nil
}

// Send OTP by SMS
func (c *Client) Send(req *streambird.OTPSendBySMSRequest) (*streambird.OTPSendBySMSResponse, error) {

	resp := &streambird.OTPSendBySMSResponse{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/send", path), req); err != nil {
		return nil, err
	}

	return resp, nil
}
