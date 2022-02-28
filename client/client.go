package client

import (
	"log"
	"os"

	streambird "github.com/streambird/streambird-go/v1"
	magiclink "github.com/streambird/streambird-go/v1/magiclink"
)

type Client struct {
	client     *streambird.Client
	MagicLinks *magiclink.Client
}

func (c *Client) SetDebugLog(debugLog *log.Logger) {
	if debugLog == nil {
		c.client.DebugLog = log.New(os.Stdout, "DEBUG", log.Lshortfile)
	} else {
		c.client.DebugLog = debugLog
	}
}

// New creates a new Streambird client object.
func New(apiKey string) *Client {
	client := streambird.New(apiKey)

	magiclink := &magiclink.Client{
		C: client,
	}
	return &Client{
		client:     client,
		MagicLinks: magiclink,
	}
}
