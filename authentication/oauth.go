package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/auth0/go-auth0/authentication/oauth"
)

// OAuth exposes logging in using OAuth based APIs.
type OAuth manager

// LoginWithGrant allows logging in with an OAuth 2.0 grant. This should only be needed if a grant
// type is not supported byt this SDK.
func (o *OAuth) LoginWithGrant(ctx context.Context, grantType string, body url.Values, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	body.Add("grant_type", grantType)

	err = o.authentication.Request(ctx, "POST", o.authentication.URI("oauth", "token"), body, &t, opts...)
	return
}

// LoginWithPassword performs the Password OAuth 2.0 grant that highly-trusted apps use to access an API.
// In this flow, the end-user is asked to fill in credentials (username/password), typically using
// an interactive form in the user-agent (browser). This information is sent to the backend and
// from there to Auth0. It is therefore imperative that the application is absolutely trusted with
// this information. For single-page applications and native/mobile apps, we recommend using web flows
// instead.
//
// See: https://auth0.com/docs/api/authentication#resource-owner-password
//
// Use the `Header` RequestOption to set the `auth0-forwarded-for` header to an end-user's IP if you
// you want brute force protection to work in server-side scenarios.
// See https://auth0.com/docs/get-started/authentication-and-authorization-flow/avoid-common-issues-with-resource-owner-password-flow-and-attack-protection
func (o *OAuth) LoginWithPassword(ctx context.Context, body oauth.LoginWithPasswordRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	grantType := "password"
	data := url.Values{
		"username": []string{body.Username},
		"password": []string{body.Password},
	}

	if body.Realm != "" {
		grantType = "http://auth0.com/oauth/grant-type/password-realm"
		data.Set("realm", body.Realm)
	}

	if body.Scope != "" {
		data.Set("scope", body.Scope)
	}

	for k, v := range body.ExtraParameters {
		data.Set(k, v)
	}

	err = o.addClientAuthentication(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, grantType, data, opts...)
	return
}

// LoginWithAuthCode performs the Authorization Code grant type OAuth 2.0 grant.
//
// This is the flow that regular web apps use to access an API. Use this endpoint to exchange an
// Authorization Code for a Token.
//
// See: https://auth0.com/docs/api/authentication?http#authorization-code-flow44
func (o *OAuth) LoginWithAuthCode(ctx context.Context, body oauth.LoginWithAuthCodeRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"code": []string{body.Code},
	}

	if body.RedirectURI != "" {
		data.Set("redirect_uri", body.RedirectURI)
	}

	err = o.addClientAuthentication(body.ClientAuthentication, data, true)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "authorization_code", data, opts...)
	return
}

// LoginWithAuthCodeWithPKCE performs the Authorization Code with Proof Key for Code Exchange OAuth 2.0 grant type.
//
// This flow was originally designed to protect the authorization code flow in mobile apps but its ability
// to prevent authorization code injection makes it useful for every type of OAuth client, even web apps
// that use client authentication.
//
// See: https://auth0.com/docs/api/authentication?http#authorization-code-flow-with-pkce45
func (o *OAuth) LoginWithAuthCodeWithPKCE(ctx context.Context, body oauth.LoginWithAuthCodeWithPKCERequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"code":          []string{body.Code},
		"code_verifier": []string{body.CodeVerifier},
	}

	if body.RedirectURI != "" {
		data.Set("redirect_uri", body.RedirectURI)
	}

	err = o.addClientAuthentication(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "authorization_code", data, opts...)
	return
}

// LoginWithClientCredentials performs the Client Credentials OAuth 2.0 grant type.
//
// This flow was originally designed to protect the authorization code flow in mobile apps but its ability
// to prevent authorization code injection makes it useful for every type of OAuth client, even web apps
// that use client authentication.
//
// See: https://auth0.com/docs/api/authentication?http#client-credentials-flow
func (o *OAuth) LoginWithClientCredentials(ctx context.Context, body oauth.LoginWithClientCredentialsRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"code": []string{body.Audience},
	}

	err = o.addClientAuthentication(body.ClientAuthentication, data, true)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "client_credentials", data, opts...)
	return
}

// RefreshToken is used to refresh and Access Token using the Refresh Token you got during authorization.
//
// See: https://auth0.com/docs/api/authentication?http#refresh-token
func (o *OAuth) RefreshToken(ctx context.Context, body oauth.RefreshTokenRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"refresh_token": []string{body.RefreshToken},
	}

	if body.Scope != "" {
		data.Set("scope", body.Scope)
	}

	err = o.addClientAuthentication(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "refresh_token", data, opts...)
	return
}

// RevokeRefreshToken is used to invalidate an Refresh Token if it has been compromised.
//
// The behaviour of this endpoint depends on the state of the Refresh Token Revocation Deletes Grant toggle.
// If this toggle is enabled, then each revocation request invalidates not only the specific token, but all
// other tokens based on the same authorization grant. This means that all Refresh Tokens that have been
// issued for the same user, application, and audience will be revoked. If this toggle is disabled, then only
// the refresh token is revoked, while the grant is left intact.
//
// See: https://auth0.com/docs/api/authentication?http#revoke-refresh-token
func (o *OAuth) RevokeRefreshToken(ctx context.Context, body oauth.RevokeRefreshTokenRequest, opts ...RequestOption) error {
	if body.ClientID == "" {
		body.ClientID = o.authentication.clientID
	}

	if body.ClientSecret != "" && o.authentication.clientSecret != "" {
		body.ClientSecret = o.authentication.clientSecret
	}

	return o.authentication.Request(ctx, "POST", o.authentication.URI("oauth", "revoke"), body, nil, opts...)
}

func (o *OAuth) addClientAuthentication(params oauth.ClientAuthentication, body url.Values, required bool) error {
	if params.ClientID != "" {
		body.Set("client_id", params.ClientID)
	} else {
		body.Set("client_id", o.authentication.clientID)
	}

	if params.ClientSecret != "" && body.Get("client_secret") == "" {
		body.Set("client_secret", params.ClientSecret)
	} else if o.authentication.clientSecret != "" {
		body.Set("client_secret", o.authentication.clientSecret)
	}

	if required && body.Get("client_secret") == "" {
		return errors.New("client_secret is required but not provided")
	}

	return nil
}