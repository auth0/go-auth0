package management

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestNewRequest(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		endpoint      string
		payload       interface{}
		options       []RequestOption
		expectedError string
		expectedBody  string
	}{
		{
			name:          "Create client request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       &Client{Name: auth0.String("TestClient"), Description: auth0.String("Test description")},
			options:       nil,
			expectedError: "",
			expectedBody:  `{"name":"TestClient","description":"Test description"}` + "\n",
		},
		{
			name:          "Create client request empty object body",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       &Client{},
			options:       nil,
			expectedError: "",
			expectedBody:  "{}\n",
		},
		{
			name:          "Read client request",
			method:        http.MethodGet,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       nil,
			options:       nil,
			expectedError: "",
			expectedBody:  "",
		},
		{
			name:          "List clients request",
			method:        http.MethodGet,
			endpoint:      api.URI("clients"),
			payload:       nil,
			options:       nil,
			expectedError: "",
			expectedBody:  "",
		},
		{
			name:          "List clients request empty body",
			method:        http.MethodGet,
			endpoint:      api.URI("clients"),
			payload:       &Client{},
			options:       nil,
			expectedError: "",
			expectedBody:  "",
		},
		{
			name:          "Update client request",
			method:        http.MethodPatch,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       &Client{Name: auth0.String("UpdatedTestClient")},
			options:       nil,
			expectedError: "",
			expectedBody:  `{"name":"UpdatedTestClient"}` + "\n",
		},
		{
			name:          "Delete client request",
			method:        http.MethodDelete,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       nil,
			options:       nil,
			expectedError: "",
			expectedBody:  "",
		},
		{
			name:          "Invalid payload request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       make(chan int),
			options:       nil,
			expectedError: "encoding request payload failed: json: unsupported type: chan int",
			expectedBody:  "",
		},
		{
			name:          "Map payload request",
			method:        http.MethodPost,
			endpoint:      api.URI("clients"),
			payload:       map[string]interface{}{},
			options:       nil,
			expectedError: "",
			expectedBody:  "{}" + "\n",
		},
		{
			name:     "Request with options",
			method:   http.MethodPost,
			endpoint: api.URI("clients"),
			payload:  nil,
			options: []RequestOption{
				Body([]byte(`{"custom":"data"}`)),
			},
			expectedError: "",
			expectedBody:  `{"custom":"data"}`,
		},
		{
			name:          "Request with Delete with Body",
			method:        http.MethodDelete,
			endpoint:      api.URI("clients", "c4vFzE4qeMgIEzRryyCmHcxGBZqswlbX"),
			payload:       &Client{Name: auth0.String("TestClient"), Description: auth0.String("Test description")},
			options:       nil,
			expectedError: "",
			expectedBody:  `{"name":"TestClient","description":"Test description"}` + "\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request, err := api.NewRequest(context.Background(), testCase.method, testCase.endpoint, testCase.payload, testCase.options...)

			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
				assert.Nil(t, request)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, request)

			requestBody, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			assert.Equal(t, testCase.expectedBody, string(requestBody))

			if testCase.expectedBody != "" {
				assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
			} else {
				assert.Empty(t, request.Header.Get("Content-Type"))
			}
		})
	}
}
