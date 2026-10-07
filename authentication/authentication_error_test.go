package authentication

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newError(t *testing.T) {
	var testCases = []struct {
		name             string
		givenResponse    http.Response
		expectedError    Error
		expectedMFAToken string

		expectedMFARequirements  *MFARequirements
		expectedValidationErrors []ValidationError
	}{
		{
			name: "it fails to decode if body is not json",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader("Hello, I'm not JSON.")),
			},
			expectedError: Error{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "failed to decode json error response payload: invalid character 'H' looking for beginning of value",
			},
			expectedMFAToken: "",
		},
		{
			name: "it correctly decodes the error response payload",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":400,"error":"invalid_scope","error_description":"Scope must be an array or a string"}`)),
			},
			expectedError: Error{
				StatusCode: 400,
				Err:        "invalid_scope",
				Message:    "Scope must be an array or a string",
			},
			expectedMFAToken: "",
		},
		{
			name: "it will still post the correct status code if the body doesn't have the correct structure",
			givenResponse: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader(`{"errorMessage":"wrongStruct"}`)),
			},
			expectedError: Error{
				StatusCode: 500,
				Err:        "Internal Server Error",
				Message:    "",
			},
			expectedMFAToken: "",
		},
		{
			name: "it will handle an invalid sign up response",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"name":"BadRequestError","code":"invalid_signup","description":"Invalid sign up","statusCode":400}`)),
			},
			expectedError: Error{
				StatusCode: 400,
				Err:        "invalid_signup",
				Message:    "Invalid sign up",
			},
			expectedMFAToken: "",
		},
		{
			name: "it will handle invalid password response",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser((strings.NewReader(`{"name":"PasswordStrengthError","message":"Password is too weak","code":"invalid_password","description":{"rules":[{"message":"At least %d characters in length","format":[8],"code":"lengthAtLeast","verified":true},{"message":"Contain at least %d of the following %d types of characters:","code":"containsAtLeast","format":[3,4],"items":[{"message":"lower case letters (a-z)","code":"lowerCase","verified":true},{"message":"upper case letters (A-Z)","code":"upperCase","verified":false},{"message":"numbers (i.e. 0-9)","code":"numbers","verified":false},{"message":"special characters (e.g. !@#$%^&*)","code":"specialCharacters","verified":true}],"verified":false}],"verified":false},"policy":"* At least 8 characters in length\n* Contain at least 3 of the following 4 types of characters:\n * lower case letters (a-z)\n * upper case letters (A-Z)\n * numbers (i.e. 0-9)\n * special characters (e.g. !@#$%^&*)","statusCode":400}`))),
			},
			expectedError: Error{
				StatusCode: 400,
				Err:        "invalid_password",
				Message:    `{"rules":[{"message":"At least %d characters in length","format":[8],"code":"lengthAtLeast","verified":true},{"message":"Contain at least %d of the following %d types of characters:","code":"containsAtLeast","format":[3,4],"items":[{"message":"lower case letters (a-z)","code":"lowerCase","verified":true},{"message":"upper case letters (A-Z)","code":"upperCase","verified":false},{"message":"numbers (i.e. 0-9)","code":"numbers","verified":false},{"message":"special characters (e.g. !@#$%^&*)","code":"specialCharacters","verified":true}],"verified":false}],"verified":false}`,
			},
			expectedMFAToken: "",
		},
		{
			name: "it will handle invalid password response with MFA token",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser((strings.NewReader(`{"name":"PasswordStrengthError","message":"Password is too weak","code":"invalid_password","description":{"rules":[{"message":"At least %d characters in length","format":[8],"code":"lengthAtLeast","verified":true},{"message":"Contain at least %d of the following %d types of characters:","code":"containsAtLeast","format":[3,4],"items":[{"message":"lower case letters (a-z)","code":"lowerCase","verified":true},{"message":"upper case letters (A-Z)","code":"upperCase","verified":false},{"message":"numbers (i.e. 0-9)","code":"numbers","verified":false},{"message":"special characters (e.g. !@#$%^&*)","code":"specialCharacters","verified":true}],"verified":false}],"verified":false},"policy":"* At least 8 characters in length\n* Contain at least 3 of the following 4 types of characters:\n * lower case letters (a-z)\n * upper case letters (A-Z)\n * numbers (i.e. 0-9)\n * special characters (e.g. !@#$%^&*)","mfa_token":"123456","statusCode":400}`))),
			},
			expectedError: Error{
				StatusCode: 400,
				Err:        "invalid_password",
				Message:    `{"rules":[{"message":"At least %d characters in length","format":[8],"code":"lengthAtLeast","verified":true},{"message":"Contain at least %d of the following %d types of characters:","code":"containsAtLeast","format":[3,4],"items":[{"message":"lower case letters (a-z)","code":"lowerCase","verified":true},{"message":"upper case letters (A-Z)","code":"upperCase","verified":false},{"message":"numbers (i.e. 0-9)","code":"numbers","verified":false},{"message":"special characters (e.g. !@#$%^&*)","code":"specialCharacters","verified":true}],"verified":false}],"verified":false}`,
				MFAToken:   "123456",
			},
			expectedMFAToken: "123456",
		},
		{
			name: "it will decode mfa requirements with the MFA token",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(`{"error":"mfa_required","error_description":"Multifactor authentication required","mfa_token":"abc","mfa_requirements":{"challenge":[{"type":"otp"}],"enroll":[{"type":"phone"}]}}`)),
			},
			expectedError: Error{
				StatusCode: 403,
				Err:        "mfa_required",
				Message:    "Multifactor authentication required",
				MFAToken:   "abc",
				MFARequirements: &MFARequirements{
					Challenge: []MFAFactor{{Type: "otp"}},
					Enroll:    []MFAFactor{{Type: "phone"}},
				},
			},
			expectedMFAToken: "abc",
			expectedMFARequirements: &MFARequirements{
				Challenge: []MFAFactor{{Type: "otp"}},
				Enroll:    []MFAFactor{{Type: "phone"}},
			},
		},
		{
			name: "it will decode validation errors",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"error":"invalid_request","error_description":"Validation failed","validation_errors":[{"field":"connection","message":"email_otp is not enabled"}]}`)),
			},
			expectedError: Error{
				StatusCode: 400,
				Err:        "invalid_request",
				Message:    "Validation failed",
				ValidationErrors: &[]ValidationError{
					{Field: "connection", Message: "email_otp is not enabled"},
				},
			},
			expectedMFAToken: "",
			expectedValidationErrors: []ValidationError{
				{Field: "connection", Message: "email_otp is not enabled"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := newError(&testCase.givenResponse)

			var actualError *Error

			ok := errors.As(err, &actualError)
			assert.True(t, ok, "newError should return an *Error")
			assert.Equal(t, testCase.expectedError, *actualError)
			assert.Equal(t, testCase.expectedMFAToken, actualError.GetMFAToken())
			assert.Equal(t, testCase.expectedMFARequirements, actualError.GetMFARequirements())
			assert.Equal(t, testCase.expectedValidationErrors, actualError.GetValidationErrors())
		})
	}
}

func TestError_NilSafeGetters(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		var err *Error

		assert.Nil(t, err.GetMFARequirements())
		assert.Nil(t, err.GetValidationErrors())
	})

	t.Run("fields absent", func(t *testing.T) {
		err := &Error{StatusCode: 400, Err: "invalid_request"}

		assert.Nil(t, err.GetMFARequirements())
		assert.Nil(t, err.GetValidationErrors())
	})
}
