package authentication

import (
	"context"
	"fmt"
	"github.com/auth0/go-auth0/authentication/ciba"
	"net/url"
	"strings"
)

// CIBA manager.
type CIBA manager

// Initiate given a CIBA request params initiates a CIBA authentication request.
//
// This sends a POST request to the /bc-authorize endpoint and returns an auth_req_id for polling.
func (c *CIBA) Initiate(ctx context.Context, body ciba.Request, opts ...RequestOption) (r *ciba.Response, err error) {
	if body.ClientID == "" {
		body.ClientID = c.authentication.clientID
	}

	if body.ClientSecret == "" {
		body.ClientSecret = c.authentication.clientSecret
	}

	var missing []string
	check(&missing, "ClientID", body.ClientID != "" || c.authentication.clientID != "")
	check(&missing, "ClientSecret", body.ClientSecret != "" || c.authentication.clientSecret != "")
	check(&missing, "LoginHint", body.LoginHint != "")
	check(&missing, "Scope", body.Scope != "")
	check(&missing, "BindingMessage", body.BindingMessage != "")

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	data := url.Values{
		"client_id":       []string{body.ClientID},
		"client_secret":   []string{body.ClientSecret},
		"login_hint":      []string{body.LoginHint},
		"scope":           []string{body.Scope},
		"binding_message": []string{body.BindingMessage},
	}

	// Perform the request
	err = c.authentication.Request(ctx, "POST", c.authentication.URI("bc-authorize"), data, &r, opts...)
	if err != nil {
		return nil, err
	}

	return
}
