package management

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
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
			uri:           "http://example.com",
			payload:       nil,
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid POST request with payload",
			method:        "POST",
			uri:           "http://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Invalid payload encoding",
			method:        "POST",
			uri:           "http://example.com",
			payload:       make(chan int), // Unsupported type
			options:       nil,
			expectedError: fmt.Errorf("encoding request payload failed"),
		},
		{
			name:          "Valid DELETE request",
			method:        "DELETE",
			uri:           "http://example.com",
			payload:       nil,
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid PUT request with payload",
			method:        "PUT",
			uri:           "http://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
		{
			name:          "Valid PATCH request with payload",
			method:        "PATCH",
			uri:           "http://example.com",
			payload:       map[string]string{"key": "value"},
			options:       nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Management{}
			request, err := m.NewRequest(context.Background(), tt.method, tt.uri, tt.payload, tt.options...)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Nil(t, request)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, request)
				assert.Equal(t, "application/json", request.Header.Get("Content-Type"))

				if tt.method == "GET" || tt.method == "DELETE" {
					requestBody, _ := io.ReadAll(request.Body)
					assert.Empty(t, string(requestBody))
				}
			}
		})
	}
}
