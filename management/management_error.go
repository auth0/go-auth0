package management

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error is an interface describing any error which
// could be returned by the Auth0 Management API.
type Error interface {
	// Status returns the status code returned by
	// the server together with the present error.
	Status() int
	error
}

type managementError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"message"`
}

func newError(response *http.Response) error {
	apiError := &managementError{}

	if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
		return &managementError{
			StatusCode: response.StatusCode,
			Err:        http.StatusText(response.StatusCode),
			Message:    fmt.Errorf("failed to decode json error response payload: %w", err).Error(),
		}
	}

	// This can happen in case the error message structure changes.
	// If that happens we still want to display the correct code.
	if apiError.Status() == 0 {
		apiError.StatusCode = response.StatusCode
		apiError.Err = http.StatusText(response.StatusCode)
	}

	return apiError
}

// Error formats the error into a string representation.
func (m *managementError) Error() string {
	return fmt.Sprintf("%d %s: %s", m.StatusCode, m.Err, m.Message)
}

// Status returns the status code of the error.
func (m *managementError) Status() int {
	return m.StatusCode
}
