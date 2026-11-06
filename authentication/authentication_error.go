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
	// MFARequirements describes the factors the user can challenge or enroll when MFA is required.
	MFARequirements *MFARequirements `json:"mfa_requirements,omitempty"`
	// ValidationErrors lists the request fields that failed validation, if any.
	ValidationErrors *[]ValidationError `json:"validation_errors,omitempty"`
}

// MFARequirements describes the MFA factors that can be challenged or enrolled
// when the API responds with an mfa_required error.
type MFARequirements struct {
	// Challenge lists the factors the user can be challenged with.
	Challenge []MFAFactor `json:"challenge,omitempty"`
	// Enroll lists the factors the user can enroll.
	Enroll []MFAFactor `json:"enroll,omitempty"`
}

// MFAFactor represents a single MFA factor.
type MFAFactor struct {
	// Type is the factor type, such as "otp", "oob", "email", "phone", "push-notification", "webauthn-roaming" or "webauthn-platform".
	Type string `json:"type"`
}

// ValidationError describes a single request field that failed validation.
type ValidationError struct {
	// Field is the name of the request field that failed validation.
	Field string `json:"field"`
	// Message describes why the field failed validation.
	Message string `json:"message"`
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

// GetMFARequirements returns the MFA requirements associated with the error, if any.
func (a *Error) GetMFARequirements() *MFARequirements {
	if a == nil {
		return nil
	}

	return a.MFARequirements
}

// GetValidationErrors returns the field level validation errors associated with the error, if any.
func (a *Error) GetValidationErrors() []ValidationError {
	if a == nil || a.ValidationErrors == nil {
		return nil
	}

	return *a.ValidationErrors
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
