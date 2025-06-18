package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/auth0/go-auth0/authentication/ciba"
	"github.com/auth0/go-auth0/internal/client"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0/authentication/oauth"
)

func TestLoginWithGrant(t *testing.T) {
	t.Run("Should return token for CIBA", func(t *testing.T) {
		// This test required approval on Guardian MFA application.
		// Hence, it cannot be recorded and is only for manual testing.
		t.Skip("Skipped as cannot be test in E2E scenario")

		// Call the Initiate method of the CIBA manager
		resp, err := authAPI.CIBA.Initiate(context.Background(), ciba.Request{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scope:        "openid",
			LoginHint: map[string]string{
				"format": "iss_sub",
				"iss":    "https://witty-silver-sailfish-sus1-staging-20240704.sus.auth0.com/",
				"sub":    "auth0|6707939cad3d8bec47ecfa2e",
			},
			BindingMessage: "TEST-BINDING-MESSAGE",
		})
		assert.Empty(t, err)
		token, err := authAPI.OAuth.LoginWithGrant(context.Background(),
			"urn:openid:params:grant-type:ciba",
			url.Values{
				"auth_req_id":   []string{resp.AuthReqID},
				"client_id":     []string{clientID},
				"client_secret": []string{clientSecret},
			},
			oauth.IDTokenValidationOptions{})
		assert.Empty(t, err)
		assert.NotEmpty(t, token.AccessToken)
	})
}
func TestOAuthLoginWithPassword(t *testing.T) {
	auth, err := New(
		context.Background(),
		domain,
		WithClientID(clientID),
	)
	require.NoError(t, err)
	t.Run("Should return tokens", func(t *testing.T) {
		configureHTTPTestRecordings(t, auth)
		user := givenAUser(t)

		tokenSet, err := auth.OAuth.LoginWithPassword(context.Background(), oauth.LoginWithPasswordRequest{
			Username: user.username,
			Password: user.password,
		}, oauth.IDTokenValidationOptions{})
		require.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should support passing extra options", func(t *testing.T) {
		configureHTTPTestRecordings(t, auth)
		user := givenAUser(t)

		tokenSet, err := auth.OAuth.LoginWithPassword(context.Background(), oauth.LoginWithPasswordRequest{
			Username: user.username,
			Password: user.password,
			Scope:    "extra-scope",
			ExtraParameters: map[string]string{
				"extra": "value",
			},
		}, oauth.IDTokenValidationOptions{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestLoginWithAuthCode(t *testing.T) {
	t.Run("Should require client_secret or client assertion", func(t *testing.T) {
		_, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{})

		assert.ErrorContains(t, err, "client_secret or client_assertion is required but not provided")
	})

	t.Run("Should throw for an invalid code", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Code: "my-invalid-code",
		}, oauth.IDTokenValidationOptions{})

		assert.Error(t, err, "Invalid authorization code")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should support setting a redirect uri", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Code:        "test-code",
			RedirectURI: "http://localhost:3000",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestLoginWithAuthCodeWithPKCE(t *testing.T) {
	t.Run("Should throw for an invalid code", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-invalid-code",
			CodeVerifier: "test-code-verifier",
		}, oauth.IDTokenValidationOptions{})

		assert.Error(t, err, "Invalid authorization code")
	})

	t.Run("Should throw for an invalid code verifier", func(t *testing.T) {
		configureHTTPTestRecordings(t, authAPI)

		_, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-code",
			CodeVerifier: "test-invalid-code-verifier",
		}, oauth.IDTokenValidationOptions{})

		assert.Error(t, err, "Failed to verify code verifier")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-code",
			CodeVerifier: "test-code-verifier",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should support setting a redirect uri", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
			Code:         "test-code",
			CodeVerifier: "test-code-verifier",
			RedirectURI:  "http://localhost:3000",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestLoginWithClientCredentials(t *testing.T) {
	t.Run("Should require client_secret or client assertion", func(t *testing.T) {
		_, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.ErrorContains(t, err, "client_secret or client_assertion is required but not provided")
	})

	t.Run("Should return tokens", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should allow overriding clientid", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
				ClientID:     "test-other-clientid",
			},
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should allow sending extra parameters", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
				ClientID:     "test-other-clientid",
			},
			Audience: "test-audience",
			ExtraParameters: map[string]string{
				"test": "value",
			},
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should support using private key jwt auth", func(t *testing.T) {
		skipE2E(t)

		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "RS256"),
		)
		require.NoError(t, err)
		configureHTTPTestRecordings(t, api)

		tokenSet, err := api.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should support passing private key jwt auth", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		auth, err := client.CreateClientAssertion("RS256", jwtPrivateKey, clientID, "https://"+domain+"/")
		require.NoError(t, err)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientAssertion:     auth,
				ClientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			},
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should error if provided an invalid algorithm", func(t *testing.T) {
		api, err := New(
			context.Background(),
			domain,
			WithIDTokenSigningAlg("HS256"),
			WithClientID(clientID),
			WithClientAssertion(jwtPrivateKey, "invalid-alg"),
		)

		require.NoError(t, err)

		_, err = api.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			Audience: "test-audience",
		}, oauth.IDTokenValidationOptions{})

		assert.ErrorContains(t, err, "unsupported client assertion algorithm \"invalid-alg\"")
	})

	t.Run("Should support passing an organization", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.LoginWithClientCredentials(context.Background(), oauth.LoginWithClientCredentialsRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			Audience:     "my-api",
			Organization: "org_test",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("Should return tokens", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.RefreshToken(context.Background(), oauth.RefreshTokenRequest{
			RefreshToken: "test-refresh-token",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.RefreshToken)
		assert.NotEmpty(t, tokenSet.Scope)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})

	t.Run("Should return tokens with reduced scopes", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		tokenSet, err := authAPI.OAuth.RefreshToken(context.Background(), oauth.RefreshTokenRequest{
			RefreshToken: "test-refresh-token",
			Scope:        "openid profile offline_access",
		}, oauth.IDTokenValidationOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenSet.AccessToken)
		assert.NotEmpty(t, tokenSet.RefreshToken)
		assert.Equal(t, "openid profile offline_access", tokenSet.Scope)
		assert.Equal(t, "Bearer", tokenSet.TokenType)
	})
}

func TestRevokeRefreshToken(t *testing.T) {
	t.Run("Should revoke token", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		err := authAPI.OAuth.RevokeRefreshToken(context.Background(), oauth.RevokeRefreshTokenRequest{
			Token: "test-refresh-token",
		})

		assert.NoError(t, err)
	})

	t.Run("Should support passing a ClientID and ClientSecret", func(t *testing.T) {
		skipE2E(t)

		auth, err := New(
			context.Background(),
			domain,
			WithClientID(clientID),
			WithClientSecret(clientSecret),
			WithIDTokenSigningAlg("HS256"),
		)
		assert.NoError(t, err)
		configureHTTPTestRecordings(t, auth)

		err = auth.OAuth.RevokeRefreshToken(context.Background(), oauth.RevokeRefreshTokenRequest{
			Token: "test-refresh-token",
		})
		assert.NoError(t, err)
	})
}

func TestOAuthWithIDTokenVerification(t *testing.T) {
	t.Run("error for an invalid organization when using org_id", func(t *testing.T) {
		extras := map[string]interface{}{
			"org_id": "org_123",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{Organization: "org_456"})

		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})

	t.Run("error for an invalid organization when using org_name", func(t *testing.T) {
		extras := map[string]interface{}{
			"org_name": "wrong-org",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{Organization: "right-org"})

		assert.ErrorContains(t, err, "org_name claim value mismatch in the ID token")
	})

	t.Run("error for an invalid nonce", func(t *testing.T) {
		extras := map[string]interface{}{
			"nonce": "wrong-nonce",
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{Nonce: "test-nonce"})

		assert.ErrorContains(t, err, "nonce claim value mismatch in the ID token; expected")
	})

	t.Run("error for an invalid maxage", func(t *testing.T) {
		extras := map[string]interface{}{
			"auth_time": time.Now().Add(-500 * time.Second).Unix(),
		}
		api, err := withIDToken(t, extras)
		assert.NoError(t, err)

		_, err = api.OAuth.LoginWithAuthCode(context.Background(), oauth.LoginWithAuthCodeRequest{
			Code: "my-code",
		}, oauth.IDTokenValidationOptions{MaxAge: 100 * time.Second})

		assert.ErrorContains(t, err, "auth_time claim in the ID token indicates that too much time has passed")
	})
}

func TestPushedAuthorizationRequest(t *testing.T) {
	t.Run("Should require a client secret", func(t *testing.T) {
		_, err := authAPI.OAuth.PushedAuthorization(context.Background(), oauth.PushedAuthorizationRequest{
			ResponseType: "code",
			RedirectURI:  "http://localhost:3000/callback",
		})
		assert.ErrorContains(t, err, "client_secret or client_assertion is required but not provided")
	})

	t.Run("Should require a ClientID, ResponseType and RedirectURI", func(t *testing.T) {
		auth, err := New(
			context.Background(),
			domain,
		)
		require.NoError(t, err)
		_, err = auth.OAuth.PushedAuthorization(context.Background(), oauth.PushedAuthorizationRequest{})
		assert.ErrorContains(t, err, "missing required fields: ClientID, ResponseType, RedirectURI")
	})

	t.Run("Should make a PAR request", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		res, err := authAPI.OAuth.PushedAuthorization(context.Background(), oauth.PushedAuthorizationRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			ResponseType: "code",
			RedirectURI:  "http://localhost:3000/callback",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, res.RequestURI)
		assert.NotEmpty(t, res.ExpiresIn)
	})

	t.Run("Should support all arguments", func(t *testing.T) {
		skipE2E(t)
		configureHTTPTestRecordings(t, authAPI)

		res, err := authAPI.OAuth.PushedAuthorization(context.Background(), oauth.PushedAuthorizationRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientSecret: clientSecret,
			},
			ResponseType:  "code",
			RedirectURI:   "http://localhost:3000/callback",
			Audience:      "test-audience",
			Nonce:         "abc123",
			ResponseMode:  "form_post",
			Scope:         "openid profile email",
			Organization:  "my-org",
			Invitation:    "invite",
			Connection:    "Username-Password",
			CodeChallenge: "n4bQgYhMfWWaL-qgxVrQFaO_TxsrC4Is0V1sFbDwCgg",
			ExtraParameters: map[string]string{
				"test": "value",
			},
		})

		require.NoError(t, err)
		assert.NotEmpty(t, res.RequestURI)
		assert.NotEmpty(t, res.ExpiresIn)
	})
}

func withIDToken(t *testing.T, extras map[string]interface{}) (*Authentication, error) {
	t.Helper()

	idTokenClientSecret := "test-client-secret"
	idTokenClientid := "test-client-id"

	var idToken string

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		tokenSet := &oauth.TokenSet{
			AccessToken: "test-access-token",
			ExpiresIn:   86400,
			IDToken:     idToken,
			TokenType:   "Bearer",
		}

		b, err := json.Marshal(tokenSet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if _, err := fmt.Fprint(w, string(b)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	s := httptest.NewTLSServer(h)

	t.Cleanup(func() {
		s.Close()
	})

	URL, err := url.Parse(s.URL)
	assert.NoError(t, err)

	api, err := New(
		context.Background(),
		URL.Host,
		WithClient(s.Client()),
		WithClientID(idTokenClientid),
		WithClientSecret(idTokenClientSecret),
		WithIDTokenSigningAlg("HS256"),
	)
	assert.NoError(t, err)

	builder := jwt.NewBuilder().
		Issuer(s.URL + "/").
		Subject("me").
		Audience([]string{idTokenClientid}).
		Expiration(time.Now().Add(time.Hour)).
		IssuedAt(time.Now().Add(-time.Hour))

	for claim, value := range extras {
		builder.Claim(claim, value)
	}

	token, err := builder.Build()

	if err != nil {
		return nil, err
	}

	b, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(idTokenClientSecret)))
	if err != nil {
		return nil, err
	}

	idToken = string(b)

	return api, nil
}
