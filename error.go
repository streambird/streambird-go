package streambird

import (
	"fmt"
)

// Error holds details including error code, human readable description and optional parameter that is related to the error.
type Error struct {
	Code        int
	Description string
	Parameter   string
}

// Error implements error interface.
func (e Error) Error() string {
	return e.Description
}

// ErrorResponse represents errored API response.
type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
	StatusCode   int64  `json:"status_code"`
}

// Error implements error interface.
func (r ErrorResponse) Error() string {
	return fmt.Sprintf("API error: %s", r.ErrorMessage)
}
