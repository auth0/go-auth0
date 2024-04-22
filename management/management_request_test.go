package management

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

func TestNewRequest(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		endpoint      string
		payload       interface{}
		options       []RequestOption
		expectedError error
		expectedBody  string
	}{
		{
			name:          "Create client request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       &Client{Name: auth0.String("TestClient"), Description: auth0.String("Test description")},
			options:       nil,
			expectedError: nil,
			expectedBody:  `{"name":"TestClient","description":"Test description"}` + "\n",
		},
		{
			name:          "Read client request",
			method:        http.MethodGet,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       nil,
			options:       nil,
			expectedError: nil,
			expectedBody:  "",
		},
		{
			name:          "List clients request",
			method:        http.MethodGet,
			endpoint:      api.URI("clients"),
			payload:       nil,
			options:       nil,
			expectedError: nil,
			expectedBody:  "",
		},
		{
			name:          "Update client request",
			method:        http.MethodPatch,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       &Client{Name: auth0.String("UpdatedTestClient")},
			options:       nil,
			expectedError: nil,
			expectedBody:  `{"name":"UpdatedTestClient"}` + "\n",
		},
		{
			name:          "Delete client request",
			method:        http.MethodDelete,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       nil,
			options:       nil,
			expectedError: nil,
			expectedBody:  "",
		},
		{
			name:          "Invalid payload request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       make(chan int),
			options:       nil,
			expectedError: fmt.Errorf("encoding request payload failed: json: unsupported type: chan int"),
			expectedBody:  "",
		},
		{
			name:          "Map payload request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       map[string]interface{}{},
			options:       nil,
			expectedError: nil,
			expectedBody:  "{}\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request, err := api.NewRequest(context.Background(), testCase.method, testCase.endpoint, testCase.payload, testCase.options...)

			if testCase.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, testCase.expectedError.Error(), err.Error())
				assert.Nil(t, request)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, request)
			assert.Equal(t, "application/json", request.Header.Get("Content-Type"))

			requestBody, _ := io.ReadAll(request.Body)
			assert.Equal(t, testCase.expectedBody, (string(requestBody)))
		})
	}
}
