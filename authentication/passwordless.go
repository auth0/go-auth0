package authentication

import (
	"context"

	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/auth0/go-auth0/authentication/passwordless"
)

// Passwordless exposes logging in using the passwordless APIs.
type Passwordless manager

// SendEmail starts a passwordless flow by sending a link or code via email.
//
// In order to set the `x-request-language` header when sending this request, use the `Header`RequestOption
// helper.
//
// See: https://auth0.com/docs/api/authentication?http#get-code-or-link
func (p *Passwordless) SendEmail(ctx context.Context, params passwordless.SendEmailRequest, opts ...RequestOption) (r *passwordless.SendEmailResponse, err error) {
	if params.ClientID == "" {
		params.ClientID = p.authentication.clientID
	}

	if params.ClientSecret == "" && p.authentication.clientSecret != "" {
		params.ClientSecret = p.authentication.clientSecret
	}

	params.Connection = "email"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("passwordless", "start"), params, &r, opts...)
	return
}

// LoginWithEmail completes the passwordless flow started in `SendEmail` by exchanging the code for a token.
//
// See: https://auth0.com/docs/api/authentication?http#authenticate-user
func (p *Passwordless) LoginWithEmail(ctx context.Context, params passwordless.LoginWithEmailRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	if params.ClientID == "" {
		params.ClientID = p.authentication.clientID
	}

	if params.ClientSecret == "" && p.authentication.clientSecret != "" {
		params.ClientSecret = p.authentication.clientSecret
	}

	params.GrantType = "http://auth0.com/oauth/grant-type/passwordless/otp"
	params.Realm = "email"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("oauth", "token"), params, &t, opts...)
	return
}

// SendSMS starts a passwordless flow by sending a code via SMS.
//
// In order to set the `x-request-language` header when sending this request, use the `Header`RequestOption
// helper.
//
// See: https://auth0.com/docs/api/authentication?http#get-code-or-link
func (p *Passwordless) SendSMS(ctx context.Context, params passwordless.SendSMSRequest, opts ...RequestOption) (r *passwordless.SendSMSResponse, err error) {
	if params.ClientID == "" {
		params.ClientID = p.authentication.clientID
	}

	if params.ClientSecret == "" && p.authentication.clientSecret != "" {
		params.ClientSecret = p.authentication.clientSecret
	}

	params.Connection = "sms"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("passwordless", "start"), params, &r, opts...)
	return
}

// LoginWithSMS completes the passwordless flow started in `SendSMS` by exchanging the code for a token.
//
// See: https://auth0.com/docs/api/authentication?http#authenticate-user
func (p *Passwordless) LoginWithSMS(ctx context.Context, params passwordless.LoginWithSMSRequest, opts ...RequestOption) (t *oauth.TokenSet, err error) {
	if params.ClientID == "" {
		params.ClientID = p.authentication.clientID
	}

	if params.ClientSecret == "" && p.authentication.clientSecret != "" {
		params.ClientSecret = p.authentication.clientSecret
	}

	params.GrantType = "http://auth0.com/oauth/grant-type/passwordless/otp"
	params.Realm = "sms"

	err = p.authentication.Request(ctx, "POST", p.authentication.URI("oauth", "token"), params, &t, opts...)
	return
}
