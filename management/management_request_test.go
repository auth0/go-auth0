package management

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		uri           string
		payload       interface{}
		options       []RequestOption
		expectedError error
	}{
		{
			name:          "Valid GET request",
			method:        "GET",
			uri:           "htestCasep://example.com",
			payload:       nil,
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid POST request with payload",
			method:        "POST",
			uri:           "htestCasep://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Invalid payload encoding",
			method:        "POST",
			uri:           "htestCasep://example.com",
			payload:       make(chan int), // Unsupported type
			options:       nil,
			expectedError: fmt.Errorf("encoding request payload failed"),
		},
		{
			name:          "Valid DELETE request",
			method:        "DELETE",
			uri:           "htestCasep://example.com",
			payload:       nil,
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid PUT request with payload",
			method:        "PUT",
			uri:           "htestCasep://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid PATCH request with payload",
			method:        "PATCH",
			uri:           "htestCasep://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m := &Management{}
			request, err := m.NewRequest(context.Background(), testCase.method, testCase.uri, testCase.payload, testCase.options...)
			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.expectedError.Error())
				assert.Nil(t, request)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, request)
				assert.Equal(t, "application/json", request.Header.Get("Content-Type"))

				if testCase.method == "GET" || testCase.method == "DELETE" {
					requestBody, _ := io.ReadAll(request.Body)
					assert.Empty(t, string(requestBody))
				}
			}
		})
	}
}
