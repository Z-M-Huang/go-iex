package iex

import (
	"fmt"
	"net/http"
)

//APIError internal api error
type APIError struct {
	StatusCode int
	Message    string
}

//Error implements the error interface
func (e APIError) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
}
