package email

import (
	"fmt"
	"net/http"

	streambird "github.com/streambird/streambird-go/v1"
)

// Client describes your balance information.
type Client struct {
	C *streambird.Client
}

const path = "auth/otps/email"

// LoginOrCreate by Email OTP
func (c *Client) LoginOrCreate(req *streambird.OTPLoginOrCreateByEmailRequest) (*streambird.OTPLoginOrCreateByEmailResponse, error) {

	resp := &streambird.OTPLoginOrCreateByEmailResponse{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/login_or_create", path), req); err != nil {
		return nil, err
	}

	return resp, nil
}

// Send OTP by Email
func (c *Client) Send(req *streambird.OTPSendByEmailRequest) (*streambird.OTPSendByEmailResponse, error) {

	resp := &streambird.OTPSendByEmailResponse{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/send", path), req); err != nil {
		return nil, err
	}

	return resp, nil
}
