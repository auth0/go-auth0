package management

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	var testCases = []struct {
		name          string
		givenResponse http.Response
		expectedError managementError
	}{
		{
			name: "it fails to decode if body is not a json",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader("Hello, I'm not a JSON.")),
			},
			expectedError: managementError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "failed to decode json error response payload: invalid character 'H' looking for beginning of value",
			},
		},
		{
			name: "it correctly decodes the error response payload",
			givenResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":400,"error":"Bad Request","message":"One of 'client_id' or 'name' is required."}`)),
			},
			expectedError: managementError{
				StatusCode: 400,
				Err:        "Bad Request",
				Message:    "One of 'client_id' or 'name' is required.",
			},
		},
		{
			name: "it will still post the correct status code if the body doesn't have the correct structure",
			givenResponse: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader(`{"errorMessage":"wrongStruct"}`)),
			},
			expectedError: managementError{
				StatusCode: 500,
				Err:        "Internal Server Error",
				Message:    "",
			},
		},
		{
			name: "it correctly decodes errorCode insufficient_entitlement",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":403,"error":"Forbidden","message":"Please upgrade your plan.","errorCode":"insufficient_entitlement"}`)),
			},
			expectedError: managementError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "Please upgrade your plan.",
				ErrorCode:  "insufficient_entitlement",
			},
		},
		{
			name: "it correctly decodes errorCode insufficient_scope",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":403,"error":"Forbidden","message":"Insufficient scope.","errorCode":"insufficient_scope"}`)),
			},
			expectedError: managementError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "Insufficient scope.",
				ErrorCode:  "insufficient_scope",
			},
		},
		{
			name: "it returns empty errorCode when field is absent",
			givenResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader(`{"statusCode":403,"error":"Forbidden","message":"Access denied."}`)),
			},
			expectedError: managementError{
				StatusCode: 403,
				Err:        "Forbidden",
				Message:    "Access denied.",
				ErrorCode:  "",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualError := newError(&testCase.givenResponse)
			assert.Equal(t, &testCase.expectedError, actualError)
		})
	}

	t.Run("Error should format into a string", func(t *testing.T) {
		err := managementError{
			StatusCode: 403,
			Err:        "Forbidden",
			Message:    "message",
		}

		assert.Equal(t, "403 Forbidden: message", err.Error())
	})

	t.Run("Code() returns the errorCode field", func(t *testing.T) {
		err := &managementError{
			StatusCode: 403,
			Err:        "Forbidden",
			Message:    "Please upgrade your plan.",
			ErrorCode:  "insufficient_entitlement",
		}
		assert.Equal(t, "insufficient_entitlement", err.Code())
		assert.Equal(t, 403, err.Status())
	})

	t.Run("Code() returns empty string when errorCode is not set", func(t *testing.T) {
		err := &managementError{
			StatusCode: 403,
			Err:        "Forbidden",
			Message:    "Access denied.",
		}
		assert.Equal(t, "", err.Code())
	})

	t.Run("fallback error from failed decode has empty Code()", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("not json")),
		}
		actualError := newError(resp)
		mErr, ok := actualError.(*managementError)
		assert.True(t, ok)
		assert.Equal(t, "", mErr.Code())
	})
}
