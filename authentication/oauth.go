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
