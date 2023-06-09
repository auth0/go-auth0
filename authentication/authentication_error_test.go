package authentication

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newError(t *testing.T) {
	var testCases = []struct {
		name          string
		givenResponse http.Response
		expectedError authenticationError
	}{
		{
			name: "it fails to decode if body is not a json",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader("Hello, I'm not a JSON.")),
			},
			expectedError: authenticationError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "failed to decode json error response payload: invalid character 'H' looking for beginning of value",
			},
		},
		{
			name: "it correctly decodes the error response payload",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":400,"error":"invalid_scope","error_description":"Scope must be an array or a string"}`)),
			},
			expectedError: authenticationError{
				StatusCode: 400,
				Err:        "invalid_scope",
				Message:    "Scope must be an array or a string",
			},
		},
		{
			name: "it will still post the correct status code if the body doesn't have the correct structure",
			givenResponse: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader(`{"errorMessage":"wrongStruct"}`)),
			},
			expectedError: authenticationError{
				StatusCode: 500,
				Err:        "Internal Server Error",
				Message:    "",
			},
		},

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualError := newError(&testCase.givenResponse)
			assert.Equal(t, &testCase.expectedError, actualError)
		})
	}
}
