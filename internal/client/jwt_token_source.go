package client

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// privateKeyJwtTokenSource implements oauth2.TokenSource for Private Key JWT client authentication.
type privateKeyJwtTokenSource struct {
	ctx                       context.Context
	uri                       string
	clientID                  string
	clientAssertionSigningAlg string
	clientAssertionSigningKey string
	audience                  string
}

// newPrivateKeyJwtTokenSource creates a new token source that uses Private Key JWT authentication.
func newPrivateKeyJwtTokenSource(
	ctx context.Context,
	uri,
	clientAssertionSigningAlg,
	clientAssertionSigningKey,
	clientID,
	audience string,
) oauth2.TokenSource {
	source := &privateKeyJwtTokenSource{
		ctx:                       ctx,
		uri:                       uri,
		clientAssertionSigningAlg: clientAssertionSigningAlg,
		clientAssertionSigningKey: clientAssertionSigningKey,
		clientID:                  clientID,
		audience:                  audience,
	}

	return oauth2.ReuseTokenSource(nil, source)
}

// Token generates a new token using Private Key JWT client authentication.
func (p privateKeyJwtTokenSource) Token() (*oauth2.Token, error) {
	alg, err := DetermineSigningAlgorithm(p.clientAssertionSigningAlg)
	if err != nil {
		return nil, fmt.Errorf("invalid algorithm: %w", err)
	}

	baseURL, err := url.Parse(p.uri)
	if err != nil {
		return nil, fmt.Errorf("invalid URI: %w", err)
	}

	assertion, err := CreateClientAssertion(
		alg,
		p.clientAssertionSigningKey,
		p.clientID,
		baseURL.JoinPath("/").String(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client assertion: %w", err)
	}

	cfg := &clientcredentials.Config{
		TokenURL:  p.uri + "/oauth/token",
		AuthStyle: oauth2.AuthStyleInParams,
		EndpointParams: url.Values{
			"audience":              []string{p.audience},
			"client_assertion_type": []string{"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
			"client_assertion":      []string{assertion},
			"grant_type":            []string{"client_credentials"},
		},
	}

	token, err := cfg.Token(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}

	return token, nil
}

// DetermineSigningAlgorithm returns the appropriate JWA signature algorithm based on the string representation.
func DetermineSigningAlgorithm(alg string) (jwa.SignatureAlgorithm, error) {
	switch alg {
	case "RS256":
		return jwa.RS256, nil
	case "RS384":
		return jwa.RS384, nil
	case "RS512":
		return jwa.RS512, nil
	case "PS256":
		return jwa.PS256, nil
	case "PS384":
		return jwa.PS384, nil
	case "PS512":
		return jwa.PS512, nil
	case "ES256":
		return jwa.ES256, nil
	case "ES384":
		return jwa.ES384, nil
	case "ES512":
		return jwa.ES512, nil
	default:
		return "", fmt.Errorf("unsupported client assertion algorithm %q", alg)
	}
}

// CreateClientAssertion creates a JWT token for client authentication with the specified lifetime.
func CreateClientAssertion(alg jwa.SignatureAlgorithm, signingKey, clientID, audience string) (string, error) {
	key, err := jwk.ParseKey([]byte(signingKey), jwk.WithPEM(true))
	if err != nil {
		return "", fmt.Errorf("failed to parse signing key: %w", err)
	}

	// Verify that the key type is compatible with the algorithm
	if err := verifyKeyCompatibility(alg, key); err != nil {
		return "", err
	}

	now := time.Now()

	token, err := jwt.NewBuilder().
		IssuedAt(now).
		NotBefore(now).
		Subject(clientID).
		JwtID(uuid.NewString()).
		Issuer(clientID).
		Audience([]string{audience}).
		Expiration(now.Add(2 * time.Minute)).
		Build()
	if err != nil {
		return "", fmt.Errorf("failed to build JWT: %w", err)
	}

	signedToken, err := jwt.Sign(token, jwt.WithKey(alg, key))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return string(signedToken), nil
}

// verifyKeyCompatibility checks if the provided key is compatible with the specified algorithm.
func verifyKeyCompatibility(alg jwa.SignatureAlgorithm, key jwk.Key) error {
	keyType := key.KeyType()

	// Check key compatibility with algorithm
	switch alg {
	case jwa.RS256, jwa.RS384, jwa.RS512, jwa.PS256, jwa.PS384, jwa.PS512:
		if keyType != "RSA" {
			return fmt.Errorf("%s algorithm requires an RSA key, but got %s", alg, keyType)
		}
	case jwa.ES256, jwa.ES384, jwa.ES512:
		if keyType != "EC" {
			return fmt.Errorf("%s algorithm requires an EC key, but got %s", alg, keyType)
		}
	default:
		return fmt.Errorf("unsupported algorithm: %s", alg)
	}

	return nil
}
