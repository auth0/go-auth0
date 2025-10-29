package core

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"testing"

	gowiremock "github.com/wiremock/go-wiremock"
	wiremocktestcontainersgo "github.com/wiremock/wiremock-testcontainers-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Global test fixtures
var (
	WireMockContainer *wiremocktestcontainersgo.WireMockContainer
	WireMockBaseURL   string
	WireMockClient    *gowiremock.Client
)

// TestMain sets up shared test fixtures for all tests in this package
func TestMain(m *testing.M) {
	// Setup shared WireMock container
	ctx := context.Background()
	container, err := wiremocktestcontainersgo.RunContainerAndStopOnCleanup(
		ctx,
		&testing.T{},
		wiremocktestcontainersgo.WithImage("docker.io/wiremock/wiremock:3.9.1"),
	)
	if err != nil {
		fmt.Printf("Failed to start WireMock container: %v\n", err)
		os.Exit(1)
	}

	// Store global references
	WireMockContainer = container

	// Try to get the base URL using the standard method first
	baseURL, err := container.Endpoint(ctx, "")
	if err == nil {
		// Standard method worked (running outside DinD)
		// This uses the mapped port (e.g., localhost:59553)
		WireMockBaseURL = "http://" + baseURL
		WireMockClient = container.Client
	} else {
		// Standard method failed, use internal IP fallback (DinD environment)
		fmt.Printf("Standard endpoint resolution failed, using internal IP fallback: %v\n", err)

		inspect, err := container.Inspect(ctx)
		if err != nil {
			fmt.Printf("Failed to inspect WireMock container: %v\n", err)
			os.Exit(1)
		}

		// Find the IP address from the container's networks
		var containerIP string
		for _, network := range inspect.NetworkSettings.Networks {
			if network.IPAddress != "" {
				containerIP = network.IPAddress
				break
			}
		}

		if containerIP == "" {
			fmt.Printf("Failed to get WireMock container IP address\n")
			os.Exit(1)
		}

		// In DinD, use the internal port directly (8080 for WireMock HTTP)
		// Don't use the mapped port since it doesn't exist in this environment
		WireMockBaseURL = fmt.Sprintf("http://%s:8080", containerIP)

		// The container.Client was created with a bad URL, so we need a new one
		WireMockClient = gowiremock.NewClient(WireMockBaseURL)
	}

	fmt.Printf("WireMock available at: %s\n", WireMockBaseURL)

	// Run all tests
	code := m.Run()

	// Cleanup
	if WireMockContainer != nil {
		WireMockContainer.Terminate(ctx)
	}

	// Exit with the same code as the tests
	os.Exit(code)
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
	defer WireMockClient.Reset()

	// Mock the OAuth token endpoint
	stub := gowiremock.Post(gowiremock.URLPathEqualTo("/oauth/token")).WillReturnResponse(
		gowiremock.NewResponse().WithJSONBody(
			map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		).WithStatus(http.StatusOK),
	)
	err := WireMockClient.StubFor(stub)
	require.NoError(t, err, "Failed to create WireMock stub")

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
	defer WireMockClient.Reset()

	// Mock the OAuth token endpoint
	stub := gowiremock.Post(gowiremock.URLPathEqualTo("/oauth/token")).WillReturnResponse(
		gowiremock.NewResponse().WithJSONBody(
			map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		).WithStatus(http.StatusOK),
	)
	err := WireMockClient.StubFor(stub)
	require.NoError(t, err, "Failed to create WireMock stub")

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
	defer WireMockClient.Reset()

	// Mock the OAuth token endpoint
	stub := gowiremock.Post(gowiremock.URLPathEqualTo("/oauth/token")).WillReturnResponse(
		gowiremock.NewResponse().WithJSONBody(
			map[string]interface{}{
				"access_token": "test-token",
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		).WithStatus(http.StatusOK),
	)
	err := WireMockClient.StubFor(stub)
	require.NoError(t, err, "Failed to create WireMock stub")

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
	defer WireMockClient.Reset()

	// Mock the OAuth token endpoint
	stub := gowiremock.Post(gowiremock.URLPathEqualTo("/oauth/token")).WillReturnResponse(
		gowiremock.NewResponse().WithJSONBody(
			map[string]interface{}{
				"access_token": "test-token-jwt",
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		).WithStatus(http.StatusOK),
	)
	err := WireMockClient.StubFor(stub)
	require.NoError(t, err, "Failed to create WireMock stub")

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
	defer WireMockClient.Reset()

	// Mock the OAuth token endpoint
	stub := gowiremock.Post(gowiremock.URLPathEqualTo("/oauth/token")).WillReturnResponse(
		gowiremock.NewResponse().WithJSONBody(
			map[string]interface{}{
				"access_token": "test-token-jwt-audience",
				"token_type":   "Bearer",
				"expires_in":   86400,
			},
		).WithStatus(http.StatusOK),
	)
	err := WireMockClient.StubFor(stub)
	require.NoError(t, err, "Failed to create WireMock stub")

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
