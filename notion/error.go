package notion

import "fmt"

// These constants represent the error codes one can get from the Notion API.
const (
	ErrorCodeInvalidJSON        = "invalid_json"
	ErrorCodeInvalidRequestURL  = "invalid_request_url"
	ErrorCodeInvalidRequest     = "invalid_request"
	ErrorCodeValidationError    = "validation_error"
	ErrorCodeUnauthorized       = "unauthorized"
	ErrorCodeRestrictedResource = "restricted_resource"
	ErrorCodeObjectNotFound     = "object_not_found"
	ErrorCodeConflict           = "conflict_error"
	ErrorCodeRateLimited        = "rate_limited"
	ErrorCodeInternalError      = "internal_server_error"
	ErrorCodeServiceUnavailable = "service_unavailable"
)

// APIError is an error from the Notion API
type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for APIError
func (a APIError) Error() string {
	return fmt.Sprintf("Notion API Error: %s - %s", a.Code, a.Message)
}

// IsAPIErrorWithCode returns true if the error is a Notion APIError and has the given code
func IsAPIErrorWithCode(err error, code string) bool {
	apiErr, ok := err.(APIError)
	return ok && apiErr.Code == code
}

// IsNotFound returns true if the error is an APIError and has the "object_not_found" code
func IsNotFound(err error) bool {
	return IsAPIErrorWithCode(err, ErrorCodeObjectNotFound)
}

// IsConflictError returns true if the error is an APIError and has the "conflict_error" code
func IsConflictError(err error) bool {
	return IsAPIErrorWithCode(err, ErrorCodeConflict)
}

// IsRateLimitError returns true if the error is an APIError and has the "rate_limited" code
func IsRateLimitError(err error) bool {
	return IsAPIErrorWithCode(err, ErrorCodeRateLimited)
}
