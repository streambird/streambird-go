package magiclink

import (
	"fmt"
	"net/http"

	streambird "github.com/streambird/streambird-go/v1"
)

// Client describes your balance information.
type Client struct {
	C *streambird.Client
}

const path = "auth/magic_links"

// Create a magic link token for the given UserID
func (c *Client) Create(createReq *streambird.MagicLinkCreateReq) (*streambird.MagicLinkCreateResp, error) {

	resp := &streambird.MagicLinkCreateResp{}
	if err := c.C.Request(resp, http.MethodPost, fmt.Sprintf("%s/create", path), createReq); err != nil {
		return nil, err
	}

	return resp, nil
}
