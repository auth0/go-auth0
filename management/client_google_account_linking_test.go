package management

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestClientGoogleAccountLinking_Marshal(t *testing.T) {
	t.Run("allow_streamlined_linking true is preserved", func(t *testing.T) {
		b, err := json.Marshal(&Client{
			GoogleAccountLinking: &ClientGoogleAccountLinking{
				AllowStreamlinedLinking: auth0.Bool(true),
			},
		})
		require.NoError(t, err)
		assert.JSONEq(t, `{"google_account_linking":{"allow_streamlined_linking":true}}`, string(b))
	})

	t.Run("allow_streamlined_linking false is preserved", func(t *testing.T) {
		b, err := json.Marshal(&Client{
			GoogleAccountLinking: &ClientGoogleAccountLinking{
				AllowStreamlinedLinking: auth0.Bool(false),
			},
		})
		require.NoError(t, err)
		assert.JSONEq(t, `{"google_account_linking":{"allow_streamlined_linking":false}}`, string(b))
	})

	t.Run("nil GoogleAccountLinking omits the key", func(t *testing.T) {
		b, err := json.Marshal(&Client{Name: auth0.String("x")})
		require.NoError(t, err)
		assert.NotContains(t, string(b), "google_account_linking")
	})
}

func TestClientGoogleAccountLinking_Unmarshal(t *testing.T) {
	t.Run("allow_streamlined_linking true is read", func(t *testing.T) {
		var c Client
		err := json.Unmarshal([]byte(`{"google_account_linking":{"allow_streamlined_linking":true}}`), &c)
		require.NoError(t, err)

		require.NotNil(t, c.GoogleAccountLinking)
		require.NotNil(t, c.GoogleAccountLinking.AllowStreamlinedLinking)
		assert.True(t, *c.GoogleAccountLinking.AllowStreamlinedLinking)
	})

	t.Run("allow_streamlined_linking false is read", func(t *testing.T) {
		var c Client
		err := json.Unmarshal([]byte(`{"google_account_linking":{"allow_streamlined_linking":false}}`), &c)
		require.NoError(t, err)

		require.NotNil(t, c.GoogleAccountLinking)
		require.NotNil(t, c.GoogleAccountLinking.AllowStreamlinedLinking)
		assert.False(t, *c.GoogleAccountLinking.AllowStreamlinedLinking)
	})

	t.Run("absent google_account_linking is nil", func(t *testing.T) {
		var c Client
		err := json.Unmarshal([]byte(`{"name":"x"}`), &c)
		require.NoError(t, err)

		assert.Nil(t, c.GoogleAccountLinking)
	})
}
