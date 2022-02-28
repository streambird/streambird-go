//
// Copyright (c) 2014 Streambird B.V.
// All rights reserved.
//
// Author: Maurice Nonnekes <maurice@streambird.com>

// Package streambird is an official library for interacting with Streambird.com API.
// The Streambird API connects your website or application to operators around the world. With our API you can integrate SMS, Chat & Voice.
// More documentation you can find on the Streambird developers portal: https://developers.streambird.com/
package streambird

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "1.0.1"

	// Endpoint points you to Streambird REST API.
	Endpoint = "http://localhost:11019/v1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

var (
	// ErrUnexpectedResponse is used when there was an internal server error and nothing can be done at this point.
	ErrUnexpectedResponse = errors.New("The Streambird API is currently unavailable")
)

// Client is used to access API with a given key.
// Uses standard lib HTTP client internally, so should be reused instead of created as needed and it is safe for concurrent use.
type Client struct {
	ApiKey     string       // The API access key.
	HTTPClient *http.Client // The HTTP client to send requests on.
	DebugLog   *log.Logger  // Optional logger for debugging purposes.
}

type contentType string

const (
	contentTypeEmpty          contentType = ""
	contentTypeJSON           contentType = "application/json"
	contentTypeFormURLEncoded contentType = "application/x-www-form-urlencoded"
)

// errorReader reads the provided byte slice into an appropriate error.
type errorReader func([]byte) error

// New creates a new Streambird client object.
func New(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: httpClientTimeout,
		},
	}
}

// Request is for internal use only and unstable.
func (c *Client) Request(v interface{}, method, path string, data interface{}) error {
	if !strings.HasPrefix(path, "https://") && !strings.HasPrefix(path, "http://") {
		path = fmt.Sprintf("%s/%s", Endpoint, path)
	}
	uri, err := url.Parse(path)
	if err != nil {
		return err
	}

	body, contentType, err := prepareRequestBody(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.ApiKey)
	request.Header.Set("User-Agent", "Streambird/ApiClient/"+ClientVersion+" Go/"+runtime.Version())
	if contentType != contentTypeEmpty {
		request.Header.Set("Content-Type", string(contentType))
	}

	if c.DebugLog != nil {
		if data != nil {
			c.DebugLog.Printf("HTTP REQUEST: %s %s %s", method, uri.String(), body)
		} else {
			c.DebugLog.Printf("HTTP REQUEST: %s %s", method, uri.String())
		}
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if c.DebugLog != nil {
		c.DebugLog.Printf("HTTP RESPONSE: %s", string(responseBody))
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// Status codes 200 and 201 are indicative of being able to convert the
		// response body to the struct that was specified.
		if err := json.Unmarshal(responseBody, &v); err != nil {
			return fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
		}

		return nil
	case http.StatusNoContent:
		// Status code 204 is returned for successful DELETE requests. Don't try to
		// unmarshal the body: that would return errors.
		return nil
	case http.StatusInternalServerError:
		// Status code 500 is a server error and means nothing can be done at this
		// point.
		return ErrUnexpectedResponse
	default:
		var errorResponse ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return err
		}

		return errorResponse
	}
}

// prepareRequestBody takes untyped data and attempts constructing a meaningful
// request body from it. It also returns the appropriate Content-Type.
func prepareRequestBody(data interface{}) ([]byte, contentType, error) {
	switch data := data.(type) {
	case nil:
		// Nil bodies are accepted by `net/http`, so this is not an error.
		return nil, contentTypeEmpty, nil
	case string:
		return []byte(data), contentTypeFormURLEncoded, nil
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, "", err
		}

		return b, contentTypeJSON, nil
	}
}
