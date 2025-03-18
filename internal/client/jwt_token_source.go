package client

import (
	"context"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type privateKeyJwtTokenSource struct {
	ctx                       context.Context
	uri                       string
	clientID                  string
	clientAssertionSigningAlg string
	clientAssertionSigningKey string
	audience                  string
}

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

func (p privateKeyJwtTokenSource) Token() (*oauth2.Token, error) {
	alg, err := determineAlg(p.clientAssertionSigningAlg)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(p.uri)
	if err != nil {
		return nil, err
	}

	assertion, err := CreateClientAssertion(alg, p.clientAssertionSigningKey, p.clientID, baseURL.JoinPath("/").String())
	if err != nil {
		return nil, err
	}

	cfg := &clientcredentials.Config{
		TokenURL:  p.uri + "/oauth/token",
		AuthStyle: oauth2.AuthStyleInParams,
		EndpointParams: url.Values{
			"audience":              []string{p.audience},
			"client_assertion_type": []string{"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
			"client_assertion":      []string{assertion},
		},
	}

	return cfg.Token(p.ctx)
}
