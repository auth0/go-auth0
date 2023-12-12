package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"error_description"`
}

func newError(response *http.Response) error {
	apiError := &AuthenticationError{}

	if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
		return &AuthenticationError{
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
func (a *AuthenticationError) Error() string {
	return fmt.Sprintf("%d %s: %s", a.StatusCode, a.Err, a.Message)
}

// Status returns the status code of the error.
func (a *AuthenticationError) Status() int {
	return a.StatusCode
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the differences between error responses between the APIs.
func (a *AuthenticationError) UnmarshalJSON(b []byte) error {
	type authError AuthenticationError
	type authErrorWrapper struct {
		*authError
		Code        string `json:"code"`
		Description string `json:"description"`
	}

	alias := &authErrorWrapper{(*authError)(a), "", ""}

	err := json.Unmarshal(b, alias)
	if err != nil {
		return err
	}

	if alias.Code != "" {
		a.Err = alias.Code
	}

	if alias.Description != "" {
		a.Message = alias.Description
	}

	return nil
}
