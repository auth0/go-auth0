package core

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// WireMock configuration - expects WireMock to be running via docker compose
var (
	WireMockBaseURL = getWireMockBaseURL()
)

func getWireMockBaseURL() string {
	if url := os.Getenv("WIREMOCK_URL"); url != "" {
		return url
	}
	return "http://localhost:8080"
}

// setupWireMockStub creates a stub for the OAuth token endpoint
func setupWireMockStub(t *testing.T, accessToken string) {
	t.Helper()

	stub := map[string]interface{}{
		"request": map[string]interface{}{
			"method":  "POST",
			"urlPath": "/oauth/token",
		},
		"response": map[string]interface{}{
			"status": http.StatusOK,
			"headers": map[string]string{
				"Content-Type": "application/json",
			},
			"jsonBody": map[string]interface{}{
				"access_token": accessToken,
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		},
	}

	body, err := json.Marshal(stub)
	require.NoError(t, err)

	resp, err := http.Post(WireMockBaseURL+"/__admin/mappings", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode, "Failed to create WireMock stub")
}

// resetWireMock clears all stubs
func resetWireMock(t *testing.T) {
	t.Helper()

	req, err := http.NewRequest(http.MethodDelete, WireMockBaseURL+"/__admin/mappings", nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
}

// generateTestRSAPrivateKey generates a test RSA private key in PEM format
func generateTestRSAPrivateKey(t *testing.T) string {
	t.Helper()

	// Generate a test RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Convert private key to PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return string(privateKeyPEM)
}

func TestClientCredentialsAndAudienceOption_TokenURL(t *testing.T) {
	// Reset WireMock state before each test
	defer resetWireMock(t)

	// Mock the OAuth token endpoint
	setupWireMockStub(t, "test-token")

	tests := []struct {
		name     string
		baseURL  string
		audience string
	}{
		{
			name:     "BaseURL with /api/v2 suffix",
			baseURL:  WireMockBaseURL + "/api/v2",
			audience: "https://test.auth0.com/api/v2/",
		},
		{
			name:     "BaseURL without /api/v2 suffix",
			baseURL:  WireMockBaseURL,
			audience: "https://test.auth0.com/api/v2/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			option := &ClientCredentialsAndAudienceOption{
				Ctx:          &ctx,
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				Audience:     tt.audience,
			}

			options := &RequestOptions{
				BaseURL: tt.baseURL,
			}

			// Apply the option which should initialize the token source
			option.applyRequestOptions(options)

			// Verify that token source was created
			require.NotNil(t, options.TokenSource, "TokenSource should be initialized")

			// Trigger a token fetch to verify the URL is correct
			token, err := options.TokenSource.Token()
			require.NoError(t, err, "Token fetch should succeed - this means TokenURL was correctly constructed as /oauth/token, not /api/v2/oauth/token")
			assert.NotEmpty(t, token.AccessToken, "Access token should not be empty")
			assert.Equal(t, "test-token", token.AccessToken)
		})
	}
}

func TestClientCredentialsOption_TokenURL(t *testing.T) {
	// Reset WireMock state before each test
	defer resetWireMock(t)

	// Mock the OAuth token endpoint
	setupWireMockStub(t, "test-token")

	tests := []struct {
		name    string
		baseURL string
	}{
		{
			name:    "BaseURL with /api/v2 suffix",
			baseURL: WireMockBaseURL + "/api/v2",
		},
		{
			name:    "BaseURL without /api/v2 suffix",
			baseURL: WireMockBaseURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			option := &ClientCredentialsOption{
				Ctx:          &ctx,
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
			}

			options := &RequestOptions{
				BaseURL: tt.baseURL,
			}

			// Apply the option which should initialize the token source
			option.applyRequestOptions(options)

			// Verify that token source was created
			require.NotNil(t, options.TokenSource, "TokenSource should be initialized")

			// Trigger a token fetch to verify the URL is correct
			token, err := options.TokenSource.Token()
			require.NoError(t, err, "Token fetch should succeed")
			assert.NotEmpty(t, token.AccessToken, "Access token should not be empty")
			assert.Equal(t, "test-token", token.AccessToken)
		})
	}
}

func TestClientCredentialsAndAudienceOption_Consistency(t *testing.T) {
	// This test ensures ClientCredentialsAndAudienceOption behaves the same as ClientCredentialsOption
	// in terms of URL construction
	defer resetWireMock(t)

	// Mock the OAuth token endpoint
	setupWireMockStub(t, "test-token")

	ctx := context.Background()
	baseURL := WireMockBaseURL + "/api/v2"

	// Test ClientCredentialsOption
	option1 := &ClientCredentialsOption{
		Ctx:          &ctx,
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
	}
	options1 := &RequestOptions{BaseURL: baseURL}
	option1.applyRequestOptions(options1)

	// Test ClientCredentialsAndAudienceOption
	option2 := &ClientCredentialsAndAudienceOption{
		Ctx:          &ctx,
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Audience:     "https://test.auth0.com/api/v2/",
	}
	options2 := &RequestOptions{BaseURL: baseURL}
	option2.applyRequestOptions(options2)

	// Both should successfully create token sources
	require.NotNil(t, options1.TokenSource)
	require.NotNil(t, options2.TokenSource)

	// Both should successfully fetch tokens
	token1, err1 := options1.TokenSource.Token()
	require.NoError(t, err1)
	assert.NotEmpty(t, token1.AccessToken)

	token2, err2 := options2.TokenSource.Token()
	require.NoError(t, err2)
	assert.NotEmpty(t, token2.AccessToken)
}

func TestClientCredentialsPrivateKeyJwtOption_TokenURL(t *testing.T) {
	// Reset WireMock state before each test
	defer resetWireMock(t)

	// Mock the OAuth token endpoint
	setupWireMockStub(t, "test-token-jwt")

	// Generate a test RSA private key
	privateKeyPEM := generateTestRSAPrivateKey(t)

	tests := []struct {
		name    string
		baseURL string
	}{
		{
			name:    "BaseURL with /api/v2 suffix",
			baseURL: WireMockBaseURL + "/api/v2",
		},
		{
			name:    "BaseURL without /api/v2 suffix",
			baseURL: WireMockBaseURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			option := &ClientCredentialsPrivateKeyJwtOption{
				Ctx:        &ctx,
				ClientID:   "test-client-id",
				PrivateKey: privateKeyPEM,
				Algorithm:  "RS256",
			}

			options := &RequestOptions{
				BaseURL: tt.baseURL,
			}

			// Apply the option which should initialize the token source
			option.applyRequestOptions(options)

			// Verify that token source was created
			require.NotNil(t, options.TokenSource, "TokenSource should be initialized")

			// Trigger a token fetch to verify the URL is correct
			token, err := options.TokenSource.Token()
			require.NoError(t, err, "Token fetch should succeed - this means TokenURL was correctly constructed as /oauth/token")
			assert.NotEmpty(t, token.AccessToken, "Access token should not be empty")
			assert.Equal(t, "test-token-jwt", token.AccessToken)
		})
	}
}

func TestClientCredentialsPrivateKeyJwtAndAudienceOption_TokenURL(t *testing.T) {
	// Reset WireMock state before each test
	defer resetWireMock(t)

	// Mock the OAuth token endpoint
	setupWireMockStub(t, "test-token-jwt-audience")

	// Generate a test RSA private key
	privateKeyPEM := generateTestRSAPrivateKey(t)

	tests := []struct {
		name     string
		baseURL  string
		audience string
	}{
		{
			name:     "BaseURL with /api/v2 suffix",
			baseURL:  WireMockBaseURL + "/api/v2",
			audience: "https://test.auth0.com/api/v2/",
		},
		{
			name:     "BaseURL without /api/v2 suffix",
			baseURL:  WireMockBaseURL,
			audience: "https://test.auth0.com/api/v2/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			option := &ClientCredentialsPrivateKeyJwtAndAudienceOption{
				Ctx:        &ctx,
				ClientID:   "test-client-id",
				PrivateKey: privateKeyPEM,
				Algorithm:  "RS256",
				Audience:   tt.audience,
			}

			options := &RequestOptions{
				BaseURL: tt.baseURL,
			}

			// Apply the option which should initialize the token source
			option.applyRequestOptions(options)

			// Verify that token source was created
			require.NotNil(t, options.TokenSource, "TokenSource should be initialized")

			// Trigger a token fetch to verify the URL is correct
			token, err := options.TokenSource.Token()
			require.NoError(t, err, "Token fetch should succeed - this means TokenURL was correctly constructed as /oauth/token")
			assert.NotEmpty(t, token.AccessToken, "Access token should not be empty")
			assert.Equal(t, "test-token-jwt-audience", token.AccessToken)
		})
	}
}
