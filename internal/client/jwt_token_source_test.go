package client

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetermineAlg(t *testing.T) {
	testCases := []struct {
		name          string
		algorithm     string
		expectedAlg   jwa.SignatureAlgorithm
		expectedError bool
	}{
		{"RS256", "RS256", jwa.RS256, false},
		{"RS384", "RS384", jwa.RS384, false},
		{"RS512", "RS512", jwa.RS512, false},
		{"PS256", "PS256", jwa.PS256, false},
		{"PS384", "PS384", jwa.PS384, false},
		{"PS512", "PS512", jwa.PS512, false},
		{"ES256", "ES256", jwa.ES256, false},
		{"ES384", "ES384", jwa.ES384, false},
		{"ES512", "ES512", jwa.ES512, false},
		{"Invalid", "INVALID", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			alg, err := DetermineSigningAlgorithm(tc.algorithm)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedAlg, alg)
			}
		})
	}
}

func TestClientAssertion(t *testing.T) {
	// Generate a test RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Convert private key to PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	clientID := "test-client-id"
	audience := "https://example.com/"
	lifetime := 2 * time.Minute

	// Create a privateKeyJwtTokenSource with our test RSA key
	ts := &privateKeyJwtTokenSource{
		uri:                       "https://example.com",
		clientID:                  clientID,
		clientAssertionSigningAlg: "RS256",
		clientAssertionSigningKey: string(privateKeyPEM),
		audience:                  audience,
	}

	// Get the signed assertion
	alg, err := DetermineSigningAlgorithm(ts.clientAssertionSigningAlg)
	require.NoError(t, err)

	baseURL, err := url.Parse(ts.uri)
	require.NoError(t, err)

	// Call the private function directly
	assertion, err := CreateClientAssertion(
		alg,
		ts.clientAssertionSigningKey,
		ts.clientID,
		baseURL.JoinPath("/").String(),
	)
	require.NoError(t, err)
	assert.NotEmpty(t, assertion)

	// Parse and validate the token
	parsedToken, err := jwt.Parse([]byte(assertion), jwt.WithVerify(false))
	require.NoError(t, err)

	// Validate claims
	assert.Equal(t, clientID, parsedToken.Subject())
	assert.Equal(t, clientID, parsedToken.Issuer())

	audiences := parsedToken.Audience()
	require.Len(t, audiences, 1)
	assert.Equal(t, ts.audience, audiences[0])

	// Check expiration
	now := time.Now()
	assert.True(t, parsedToken.IssuedAt().Before(now) || parsedToken.IssuedAt().Equal(now))
	assert.True(t, parsedToken.Expiration().After(now))

	// Check that the expiration is roughly lifetime away from now
	assert.WithinDuration(t, now.Add(lifetime), parsedToken.Expiration(), 5*time.Second)
}

func TestECClientAssertion(t *testing.T) {
	// Generate a test EC key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	// Convert EC key to PEM
	ecPrivateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	require.NoError(t, err)

	ecPrivateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: ecPrivateKeyBytes,
	})

	clientID := "test-client-id"
	audience := "https://example.com/"
	lifetime := 2 * time.Minute

	// Create a privateKeyJwtTokenSource with our test EC key
	ts := &privateKeyJwtTokenSource{
		uri:                       "https://example.com",
		clientID:                  clientID,
		clientAssertionSigningAlg: "ES256",
		clientAssertionSigningKey: string(ecPrivateKeyPEM),
		audience:                  audience,
	}

	// Get the signed assertion
	alg, err := DetermineSigningAlgorithm(ts.clientAssertionSigningAlg)
	require.NoError(t, err)

	baseURL, err := url.Parse(ts.uri)
	require.NoError(t, err)

	// Call the private function directly
	assertion, err := CreateClientAssertion(
		alg,
		ts.clientAssertionSigningKey,
		ts.clientID,
		baseURL.JoinPath("/").String(),
	)
	require.NoError(t, err)
	assert.NotEmpty(t, assertion)

	// Parse and validate the token
	parsedToken, err := jwt.Parse([]byte(assertion), jwt.WithVerify(false))
	require.NoError(t, err)

	// Validate claims
	assert.Equal(t, clientID, parsedToken.Subject())
	assert.Equal(t, clientID, parsedToken.Issuer())

	audiences := parsedToken.Audience()
	require.Len(t, audiences, 1)
	assert.Equal(t, ts.audience, audiences[0])

	// Check expiration
	now := time.Now()
	assert.True(t, parsedToken.IssuedAt().Before(now) || parsedToken.IssuedAt().Equal(now))
	assert.True(t, parsedToken.Expiration().After(now))

	// Check that the expiration is roughly lifetime away from now
	assert.WithinDuration(t, now.Add(lifetime), parsedToken.Expiration(), 5*time.Second)

	// Verify the algorithm header from the JWT directly
	// Split the JWT into its parts
	parts := strings.Split(assertion, ".")
	require.GreaterOrEqual(t, len(parts), 1, "JWT should have at least a header part")

	// Decode the header part (part[0])
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	require.NoError(t, err)

	// Parse the header as JSON
	var header map[string]interface{}
	err = json.Unmarshal(headerBytes, &header)
	require.NoError(t, err)

	// Check the algorithm
	algHeader, ok := header["alg"]
	require.True(t, ok, "alg header should be present")
	assert.Equal(t, "ES256", algHeader)
}

func TestIncompatibleKeyTypeForAlgorithm(t *testing.T) {
	// Generate an RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	ts := &privateKeyJwtTokenSource{
		uri:                       "https://example.com",
		clientID:                  "client-id",
		clientAssertionSigningAlg: "ES256", // Incompatible with RSA key
		clientAssertionSigningKey: string(privateKeyPEM),
	}

	// Get the signed assertion
	alg, err := DetermineSigningAlgorithm(ts.clientAssertionSigningAlg)
	require.NoError(t, err)

	baseURL, err := url.Parse(ts.uri)
	require.NoError(t, err)

	// Try to use RSA key with ES256 (elliptic curve algorithm)
	_, err = CreateClientAssertion(
		alg,
		ts.clientAssertionSigningKey,
		ts.clientID,
		baseURL.JoinPath("/").String(),
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ES256 algorithm requires an EC key")
}

func TestVerifyKeyCompatibility(t *testing.T) {
	// Generate an RSA key
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	rsaKeyBytes := x509.MarshalPKCS1PrivateKey(rsaKey)
	rsaKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: rsaKeyBytes,
	})

	rsaJWK, err := jwk.ParseKey(rsaKeyPEM, jwk.WithPEM(true))
	require.NoError(t, err)

	// Test compatible combinations
	assert.NoError(t, verifyKeyCompatibility(jwa.RS256, rsaJWK))
	assert.NoError(t, verifyKeyCompatibility(jwa.PS256, rsaJWK))

	// Test incompatible combinations
	assert.Error(t, verifyKeyCompatibility(jwa.ES256, rsaJWK))
}

func TestPrivateKeyJwtTokenSource(t *testing.T) {
	// Generate a test RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Convert private key to PEM
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for client_assertion_type and grant_type
		err := r.ParseForm()
		require.NoError(t, err)

		assert.Equal(t, "urn:ietf:params:oauth:client-assertion-type:jwt-bearer", r.Form.Get("client_assertion_type"))
		assert.Equal(t, "client_credentials", r.Form.Get("grant_type"))
		assert.NotEmpty(t, r.Form.Get("client_assertion"))

		// Return a mock token
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"access_token": "mock-access-token",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	}))
	defer server.Close()

	// Create token source
	tokenSource := newPrivateKeyJwtTokenSource(
		context.Background(),
		server.URL,
		"RS256",
		string(privateKeyPEM),
		"test-client-id",
		"test-audience",
	)

	// Get token
	token, err := tokenSource.Token()
	require.NoError(t, err)
	assert.Equal(t, "mock-access-token", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestPrivateKeyJwtTokenSourceRefresh(t *testing.T) {
	// Generate a test RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// Track token request count
	requestCount := 0

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount++

		// Return a token with short expiration
		w.Header().Set("Content-Type", "application/json")
		w.Write(fmt.Appendf(nil, `{
            "access_token": "mock-token-%d",
            "token_type": "Bearer",
            "expires_in": 2
        }`, requestCount))
	}))
	defer server.Close()

	// Create token source
	tokenSource := newPrivateKeyJwtTokenSource(
		context.Background(),
		server.URL,
		"RS256",
		string(privateKeyPEM),
		"test-client-id",
		"test-audience",
	)

	// Get first token
	token1, err := tokenSource.Token()
	require.NoError(t, err)
	assert.Equal(t, "mock-token-1", token1.AccessToken)

	// Wait for token to expire (just over 2 seconds)
	time.Sleep(3 * time.Second)

	// Get second token - should trigger a refresh
	token2, err := tokenSource.Token()
	require.NoError(t, err)
	assert.Equal(t, "mock-token-2", token2.AccessToken)
	assert.NotEqual(t, token1.AccessToken, token2.AccessToken)

	// Verify server received two requests
	assert.Equal(t, 2, requestCount)
}

func TestPrivateKeyJwtTokenSourceErrors(t *testing.T) {
	// Generate a test RSA key for valid cases
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	t.Run("invalid algorithm", func(t *testing.T) {
		ts := &privateKeyJwtTokenSource{
			ctx:                       context.Background(),
			uri:                       "https://example.com",
			clientID:                  "client-id",
			clientAssertionSigningAlg: "INVALID_ALG", // Invalid algorithm
			clientAssertionSigningKey: string(privateKeyPEM),
			audience:                  "audience",
		}

		_, err := ts.Token()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid algorithm")
	})

	t.Run("invalid URI", func(t *testing.T) {
		ts := &privateKeyJwtTokenSource{
			ctx:                       context.Background(),
			uri:                       "://invalid-uri", // Invalid URI format
			clientID:                  "client-id",
			clientAssertionSigningAlg: "RS256",
			clientAssertionSigningKey: string(privateKeyPEM),
			audience:                  "audience",
		}

		_, err := ts.Token()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid URI")
	})

	t.Run("invalid signing key", func(t *testing.T) {
		ts := &privateKeyJwtTokenSource{
			ctx:                       context.Background(),
			uri:                       "https://example.com",
			clientID:                  "client-id",
			clientAssertionSigningAlg: "RS256",
			clientAssertionSigningKey: "not-a-valid-key", // Invalid key
			audience:                  "audience",
		}

		_, err := ts.Token()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create client assertion")
	})

	t.Run("token request failure", func(t *testing.T) {
		// Create a server that returns an error
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "server_error"}`))
		}))
		defer server.Close()

		ts := &privateKeyJwtTokenSource{
			ctx:                       context.Background(),
			uri:                       server.URL,
			clientID:                  "client-id",
			clientAssertionSigningAlg: "RS256",
			clientAssertionSigningKey: string(privateKeyPEM),
			audience:                  "audience",
		}

		_, err := ts.Token()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token request failed")
	})

	t.Run("connection refused error", func(t *testing.T) {
		// Create a server, then immediately close it to simulate a connection refused error
		server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
		serverURL := server.URL
		server.Close() // Close immediately to simulate connection failure

		ts := &privateKeyJwtTokenSource{
			ctx:                       context.Background(),
			uri:                       serverURL,
			clientID:                  "client-id",
			clientAssertionSigningAlg: "RS256",
			clientAssertionSigningKey: string(privateKeyPEM),
			audience:                  "audience",
		}

		_, err := ts.Token()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token request failed")
	})
}

func TestCreateClientAssertionErrors(t *testing.T) {
	// Test with invalid key
	t.Run("invalid key", func(t *testing.T) {
		_, err := CreateClientAssertion(
			jwa.RS256,
			"not-a-valid-key",
			"client-id",
			"https://example.com/",
		)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse signing key")
	})

	// Test with incompatible key type
	t.Run("incompatible key type", func(t *testing.T) {
		// Generate EC key
		ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)

		ecKeyBytes, err := x509.MarshalECPrivateKey(ecKey)
		require.NoError(t, err)

		ecKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: ecKeyBytes,
		})

		// Try to use EC key with RS256
		_, err = CreateClientAssertion(
			jwa.RS256,
			string(ecKeyPEM),
			"client-id",
			"https://example.com/",
		)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "requires an RSA key")
	})

	// Test unsupported algorithm
	t.Run("unsupported algorithm", func(t *testing.T) {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		require.NoError(t, err)

		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privateKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		})

		key, err := jwk.ParseKey(privateKeyPEM, jwk.WithPEM(true))
		require.NoError(t, err)

		// Using an unsupported algorithm
		err = verifyKeyCompatibility("unsupported" /* using string instead of jwa.SignatureAlgorithm */, key)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported algorithm")
	})
}
