package main

import (
	"context"
	"log"

	"github.com/auth0/go-auth0/v2/authentication"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
)

func main() {
	ctx := context.Background()
	authClient, err := authentication.New(
		ctx,
		"auth-staging.shipt.com",
	)

	if err != nil {
		panic(err)
	}

	tokenSet, err := authClient.OAuth.LoginWithCustomTokenExchange(
		ctx,
		oauth.LoginWithCustomTokenExchangeRequest{
			ClientAuthentication: oauth.ClientAuthentication{
				ClientID:     "6OOhdCBIBAOnXHGQyVCqymm347wyADhX",
				ClientSecret: "2WanPSL35F8tX7STZn5PaY9HAozcpxnhjd7XFjBqRiEOFeck5a_0AMbDWocAgxvq",
			},
			SubjectToken:     "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ikd0NmppNUlHUVBkb0FxSW45cU5xaCJ9.eyJodHRwczovL3d3dy5zaGlwdC5jb20vc2hpcHRfdXNlcl9lbWFpbCI6ImpvaG4ucm9lc2xlcittYXIxOTI2QHNoaXB0LmNvbSIsImh0dHBzOi8vd3d3LnNoaXB0LmNvbS9zaGlwdF91c2VyX2lkIjoiODAwNjkyOCIsImlzcyI6Imh0dHBzOi8vYXV0aC1zdGFnaW5nLnNoaXB0LmNvbS8iLCJzdWIiOiJhdXRoMHw2OWJjNTFkYmYxNjIxM2Y4MjgzMDBlZmYiLCJhdWQiOlsiaHR0cHM6Ly9zdGFnaW5nLW1lbWJlci1hcGkuc2hpcHQuY29tLyIsImh0dHBzOi8vc3RnLWZvaC5zaGlwdC1zdGFnaW5nLmF1dGgwYXBwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE3Nzk4MTQ4NDYsImV4cCI6MTc3OTgxNTE0Niwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCBpczpDdXN0b21lciBvZmZsaW5lX2FjY2VzcyIsImF6cCI6IjU3cUdyY0hlZ0R5RlpKOFd1VHQ1aVFXZjc1UUI1V2ZJIn0.2DLkIVFcWR7ExKi90UA_oGYmTOZtg6wQn0YFXzTcxWfkW04cbNjBqITibXMVA1CHf_xrGfMhLu3OJ7gsYyBTpSr0tcGE1qcDGXmRhxMNmHcFVCrORm9gP6Wt1l5IyU1V-wjHBe-GA1suONvkHgsqoYn18PzRRLu3Ig0iLpp0T0pLNpU-eu2xGM0xIHXN8PT1zoGajeszNQw6qtCdy1u32wa-GGrZvq4EintB7H2LTV1hk_B3JM8qsdDN3DMvSkuuMLIcUT9A36TCfA6-fT7Mpj7Bi1yhylsQjlb7BhZ9xtcOES7wCFrDbhIPj6YNUi-OczdoGmLLXm4wj1XEctFUEQ",
			SubjectTokenType: "https://staging-member-api.shipt.com/",
			Audience:         "https://staging-member-api.shipt.com/cart-builder/",
			Scope:            "openid email profile is:Customer",
		},
		oauth.IDTokenValidationOptions{},
	)
	if err != nil {
		log.Fatalf("failed to exchange token: %v", err)
	}
	log.Printf("got token: %+v", tokenSet)
}
