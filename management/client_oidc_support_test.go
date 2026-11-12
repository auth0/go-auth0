package management

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestClientOIDCSupport_Marshal(t *testing.T) {
	t.Run("explicit is_allowed false is preserved", func(t *testing.T) {
		b, err := json.Marshal(&ClientOIDCSupport{IsAllowed: auth0.Bool(false)})
		require.NoError(t, err)
		assert.JSONEq(t, `{"is_allowed":false}`, string(b))
	})

	t.Run("empty allowed_scopes is preserved", func(t *testing.T) {
		b, err := json.Marshal(&ClientOIDCSupport{
			IsAllowed:     auth0.Bool(true),
			AllowedScopes: &[]string{},
		})
		require.NoError(t, err)
		assert.JSONEq(t, `{"is_allowed":true,"allowed_scopes":[]}`, string(b))
	})

	t.Run("allow_all_scopes false alongside scopes is preserved", func(t *testing.T) {
		b, err := json.Marshal(&ClientOIDCSupport{
			IsAllowed:      auth0.Bool(true),
			AllowedScopes:  &[]string{OIDCScopeProfile, OIDCScopeEmail},
			AllowAllScopes: auth0.Bool(false),
		})
		require.NoError(t, err)
		assert.JSONEq(t, `{"is_allowed":true,"allowed_scopes":["profile","email"],"allow_all_scopes":false}`, string(b))
	})

	t.Run("nil OIDCSupport omits the key", func(t *testing.T) {
		b, err := json.Marshal(&Client{Name: auth0.String("x")})
		require.NoError(t, err)
		assert.NotContains(t, string(b), "oidc_support")
	})
}

func TestClientOIDCSupport_Unmarshal(t *testing.T) {
	var c Client
	err := json.Unmarshal([]byte(`{"oidc_support":{"is_allowed":true,"allow_all_scopes":true}}`), &c)
	require.NoError(t, err)

	assert.True(t, c.GetOIDCSupport().GetIsAllowed())
	assert.True(t, c.GetOIDCSupport().GetAllowAllScopes())
	assert.Empty(t, c.GetOIDCSupport().GetAllowedScopes())
}
