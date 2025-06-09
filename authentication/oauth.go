package authentication

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/auth0/go-auth0/internal/idtokenvalidator"
)

// OAuth exposes logging in using OAuth based APIs.
type OAuth manager

// LoginWithGrant allows logging in with an OAuth 2.0 grant. This should only be needed if a grant
// type is not supported byt this SDK.
func (o *OAuth) LoginWithGrant(ctx context.Context, grantType string, body url.Values, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	body.Add("grant_type", grantType)

	err = o.authentication.Request(ctx, "POST", o.authentication.URI("oauth", "token"), body, &t, opts...)

	if t != nil && t.IDToken != "" {
		err = o.authentication.idTokenValidator.Validate(t.IDToken, idtokenvalidator.ValidationOptions{
			MaxAge:       validationOptions.MaxAge,
			Nonce:        validationOptions.Nonce,
			Organization: validationOptions.Organization,
		})

		if err != nil {
			return nil, err
		}
	}

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
// want brute force protection to work in server-side scenarios.
// See https://auth0.com/docs/get-started/authentication-and-authorization-flow/avoid-common-issues-with-resource-owner-password-flow-and-attack-protection
func (o *OAuth) LoginWithPassword(ctx context.Context, body oauth.LoginWithPasswordRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
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

	if body.Audience != "" {
		data.Set("audience", body.Audience)
	}

	for k, v := range body.ExtraParameters {
		data.Set(k, v)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, grantType, data, validationOptions, opts...)

	return
}

// LoginWithAuthCode performs the Authorization Code grant type OAuth 2.0 grant.
//
// This is the flow that regular web apps use to access an API. Use this endpoint to exchange an
// Authorization Code for a token.
//
// See: https://auth0.com/docs/api/authentication?http#authorization-code-flow44
func (o *OAuth) LoginWithAuthCode(ctx context.Context, body oauth.LoginWithAuthCodeRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"code": []string{body.Code},
	}

	if body.RedirectURI != "" {
		data.Set("redirect_uri", body.RedirectURI)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "authorization_code", data, validationOptions, opts...)

	return
}

// LoginWithAuthCodeWithPKCE performs the Authorization Code with Proof Key for Code Exchange OAuth 2.0 grant type.
//
// This flow was originally designed to protect the authorization code flow in mobile apps but its ability
// to prevent authorization code injection makes it useful for every type of OAuth client, even web apps
// that use client authentication.
//
// See: https://auth0.com/docs/api/authentication?http#authorization-code-flow-with-pkce45
func (o *OAuth) LoginWithAuthCodeWithPKCE(ctx context.Context, body oauth.LoginWithAuthCodeWithPKCERequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"code":          []string{body.Code},
		"code_verifier": []string{body.CodeVerifier},
	}

	if body.RedirectURI != "" {
		data.Set("redirect_uri", body.RedirectURI)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "authorization_code", data, validationOptions, opts...)

	return
}

// LoginWithClientCredentials performs the Client Credentials OAuth 2.0 grant type.
//
// Use this endpoint to directly request an access token by using the Client's credentials (a Client ID and
// a Client Secret).
//
// See: https://auth0.com/docs/api/authentication?http#client-credentials-flow
func (o *OAuth) LoginWithClientCredentials(ctx context.Context, body oauth.LoginWithClientCredentialsRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"audience": []string{body.Audience},
	}

	if body.Organization != "" {
		data.Set("organization", body.Organization)
	}

	for k, v := range body.ExtraParameters {
		data.Set(k, v)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "client_credentials", data, validationOptions, opts...)

	return
}

// RefreshToken is used to refresh and access token using the refresh token you got during authorization.
//
// See: https://auth0.com/docs/api/authentication?http#refresh-token
func (o *OAuth) RefreshToken(ctx context.Context, body oauth.RefreshTokenRequest, validationOptions oauth.IDTokenValidationOptions, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	data := url.Values{
		"refresh_token": []string{body.RefreshToken},
	}

	if body.Scope != "" {
		data.Set("scope", body.Scope)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, false)

	if err != nil {
		return
	}

	t, err = o.LoginWithGrant(ctx, "refresh_token", data, validationOptions, opts...)

	return
}

// RevokeRefreshToken is used to invalidate a refresh token if it has been compromised.
//
// The behaviour of this endpoint depends on the state of the **Refresh Token Revocation Deletes Grant** toggle.
// If this toggle is enabled, then each revocation request invalidates not only the specific token, but all
// other tokens based on the same authorization grant. This means that all refresh tokens that have been
// issued for the same user, application, and audience will be revoked. If this toggle is disabled, then only
// the refresh token is revoked, while the grant is left intact.
//
// See: https://auth0.com/docs/api/authentication?http#revoke-refresh-token
func (o *OAuth) RevokeRefreshToken(ctx context.Context, body oauth.RevokeRefreshTokenRequest, opts ...RequestOption) error {
	if body.ClientID == "" {
		body.ClientID = o.authentication.clientID
	}

	if body.ClientSecret == "" && o.authentication.clientSecret != "" {
		body.ClientSecret = o.authentication.clientSecret
	}

	return o.authentication.Request(ctx, "POST", o.authentication.URI("oauth", "revoke"), body, nil, opts...)
}

// PushedAuthorization performs a Pushed Authorization Request that can be used to initiate an OAuth flow from
// the backchannel instead of building a URL.
//
// See: https://www.rfc-editor.org/rfc/rfc9126.html
func (o *OAuth) PushedAuthorization(ctx context.Context, body oauth.PushedAuthorizationRequest, opts ...RequestOption) (p *oauth.PushedAuthorizationRequestResponse, err error) {
	missing := []string{}
	check(&missing, "ClientID", (body.ClientID != "" || o.authentication.clientID != ""))
	check(&missing, "ResponseType", body.ResponseType != "")
	check(&missing, "RedirectURI", body.RedirectURI != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"response_type": []string{body.ResponseType},
		"redirect_uri":  []string{body.RedirectURI},
	}

	addIfNotEmpty("scope", body.Scope, data)
	addIfNotEmpty("audience", body.Audience, data)
	addIfNotEmpty("nonce", body.Nonce, data)
	addIfNotEmpty("response_mode", body.ResponseMode, data)
	addIfNotEmpty("organization", body.Organization, data)
	addIfNotEmpty("invitation", body.Invitation, data)
	addIfNotEmpty("connection", body.Connection, data)
	addIfNotEmpty("code_challenge", body.CodeChallenge, data)

	for key, value := range body.ExtraParameters {
		data.Set(key, value)
	}

	err = o.authentication.addClientAuthenticationToURLValues(body.ClientAuthentication, data, true)

	if err != nil {
		return nil, err
	}

	err = o.authentication.Request(ctx, "POST", o.authentication.URI("oauth", "par"), data, &p, opts...)

	if err != nil {
		return nil, err
	}

	return
}
