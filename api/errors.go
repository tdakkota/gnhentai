package api

import "fmt"

// APIError represents an error returned by the API.
type APIError struct {
	// StatusCode is a status code of the error.
	StatusCode int
	// Message is a message of the error.
	Message string
}

var _ error = (*APIError)(nil)

// Error implements [error].
func (e *APIError) Error() string {
	return fmt.Sprintf("code %d: %s", e.StatusCode, e.Message)
}
