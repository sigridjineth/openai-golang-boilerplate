package errors

import "fmt"

// APIError represents an error that occured on an API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("[%d:%s] %s", e.StatusCode, e.Type, e.Message)
}

type APIErrorResponse struct {
	Error APIError `json:"error"`
}
