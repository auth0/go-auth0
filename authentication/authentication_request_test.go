package authentication

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthentication_NewRequest(t *testing.T) {
	testURL, _ := url.Parse("https://example.com")
	auth := &Authentication{
		url: testURL,
	}

	t.Run("creates request successfully", func(t *testing.T) {
		payload := map[string]interface{}{
			"key": "value",
		}

		req, err := auth.NewRequest(context.Background(), "POST", "/test", payload)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.String(), "/test")
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	})

	t.Run("handles nil payload", func(t *testing.T) {
		req, err := auth.NewRequest(context.Background(), "GET", "/test", nil)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "GET", req.Method)
	})

	t.Run("applies valid options", func(t *testing.T) {
		validOption := Header("X-Test", "test-value")

		req, err := auth.NewRequest(context.Background(), "GET", "/test", nil, validOption)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "test-value", req.Header.Get("X-Test"))
	})

	t.Run("handles nil options correctly", func(t *testing.T) {
		validOption := Header("X-Test", "test-value")

		// Test with a mix of nil and valid options to exercise the nil check
		req, err := auth.NewRequest(context.Background(), "GET", "/test", nil, validOption, nil, validOption)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "test-value", req.Header.Get("X-Test"))
	})

	t.Run("handles all nil options", func(t *testing.T) {
		req, err := auth.NewRequest(context.Background(), "GET", "/test", nil, nil, nil)

		assert.NoError(t, err)
		assert.NotNil(t, req)
	})
}

func TestAuthentication_NewFormRequest(t *testing.T) {
	testURL, _ := url.Parse("https://example.com")
	auth := &Authentication{
		url: testURL,
	}

	t.Run("creates form request successfully", func(t *testing.T) {
		payload := url.Values{}
		payload.Set("key", "value")
		payload.Set("another", "test")

		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.String(), "/form")
		assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
		assert.Greater(t, req.ContentLength, int64(0))
	})

	t.Run("handles empty payload", func(t *testing.T) {
		payload := url.Values{}

		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, int64(0), req.ContentLength)
	})

	t.Run("applies valid options", func(t *testing.T) {
		payload := url.Values{}
		payload.Set("key", "value")

		validOption := Header("X-Custom", "custom-value")

		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload, validOption)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "custom-value", req.Header.Get("X-Custom"))
	})

	t.Run("handles nil options correctly", func(t *testing.T) {
		payload := url.Values{}
		payload.Set("key", "value")

		validOption := Header("X-Custom", "custom-value")

		// Test with a mix of nil and valid options to exercise the nil check
		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload, validOption, nil, validOption)

		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "custom-value", req.Header.Get("X-Custom"))
	})

	t.Run("handles all nil options", func(t *testing.T) {
		payload := url.Values{}
		payload.Set("key", "value")

		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload, nil, nil)

		assert.NoError(t, err)
		assert.NotNil(t, req)
	})

	t.Run("option can modify form payload", func(t *testing.T) {
		payload := url.Values{}
		payload.Set("original", "value")

		// Create an option that modifies the form payload
		modifyFormOption := newRequestOption(func(_ *http.Request, payload url.Values) {
			if payload != nil {
				payload.Set("added", "by-option")
			}
		})

		req, err := auth.NewFormRequest(context.Background(), "POST", "/form", payload, modifyFormOption)

		assert.NoError(t, err)
		assert.NotNil(t, req)

		// Verify the form data includes both original and added values
		body := req.Body
		require.NotNil(t, body)

		// Read the body to check the form data
		bodyBytes := make([]byte, req.ContentLength)
		_, err = body.Read(bodyBytes)
		assert.NoError(t, err)

		bodyString := string(bodyBytes)
		assert.Contains(t, bodyString, "original=value")
		assert.Contains(t, bodyString, "added=by-option")
	})
}

func TestAuthentication_RequestOption_Header(t *testing.T) {
	t.Run("sets header correctly", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		option := Header("Authorization", "Bearer token123")
		option.apply(req, nil)

		assert.Equal(t, "Bearer token123", req.Header.Get("Authorization"))
	})

	t.Run("sets multiple headers", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		option1 := Header("X-Custom-1", "value1")
		option2 := Header("X-Custom-2", "value2")

		option1.apply(req, nil)
		option2.apply(req, nil)

		assert.Equal(t, "value1", req.Header.Get("X-Custom-1"))
		assert.Equal(t, "value2", req.Header.Get("X-Custom-2"))
	})
}

func TestAuthentication_RequestOption_WithForm(t *testing.T) {
	t.Run("option works with form payload", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", strings.NewReader(""))
		require.NoError(t, err)

		payload := url.Values{}
		payload.Set("test", "value")

		option := Header("Content-Type", "application/x-www-form-urlencoded")
		option.apply(req, payload)

		assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
	})
}
