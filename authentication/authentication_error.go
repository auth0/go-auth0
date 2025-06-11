package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error represents errors returned from the Authentication API. The `Err` property can
// be used to check the error code returned from the API.
type Error struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"error_description"`
	MFAToken   string `json:"mfa_token,omitempty"`
}

func newError(response *http.Response) error {
	apiError := &Error{}

	if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
		return &Error{
			StatusCode: response.StatusCode,
			Err:        http.StatusText(response.StatusCode),
			Message:    fmt.Errorf("failed to decode json error response payload: %w", err).Error(),
		}
	}

	// This can happen in case the error message structure changes.
	// If that happens we still want to display the correct code.
	if apiError.Status() == 0 {
		apiError.StatusCode = response.StatusCode
		if apiError.Err == "" {
			apiError.Err = http.StatusText(response.StatusCode)
		}
	}

	return apiError
}

// Error formats the error into a string representation.
func (a *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", a.StatusCode, a.Err, a.Message)
}

// GetMFAToken returns the MFA token associated with the error, if any.
func (a *Error) GetMFAToken() string {
	if a == nil || a.MFAToken == "" {
		return ""
	}

	return a.MFAToken
}

// Status returns the status code of the error.
func (a *Error) Status() int {
	return a.StatusCode
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the differences between error responses between the APIs.
func (a *Error) UnmarshalJSON(b []byte) error {
	type authError Error

	type authErrorWrapper struct {
		*authError
		Code        string          `json:"code"`
		Description json.RawMessage `json:"description"` // Can be string or object
	}

	alias := &authErrorWrapper{(*authError)(a), "", nil}

	err := json.Unmarshal(b, alias)
	if err != nil {
		return err
	}

	if alias.Code != "" {
		a.Err = alias.Code
	}

	if len(alias.Description) > 0 {
		var descText string

		err := json.Unmarshal(alias.Description, &descText)
		if err == nil {
			a.Message = descText
		} else {
			a.Message = string(alias.Description)
		}
	}

	return nil
}
