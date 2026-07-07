package management

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectionOptionsOAuth2UnmarshalCustomHeaders(t *testing.T) {
	t.Run("map format", func(t *testing.T) {
		payload := `{
			"authorizationURL": "https://example.com/oauth/authorize",
			"tokenURL": "https://example.com/oauth/token",
			"customHeaders": {"X-Custom-Header": "custom-value"}
		}`

		var opts ConnectionOptionsOAuth2
		require.NoError(t, json.Unmarshal([]byte(payload), &opts))
		assert.Equal(t, ConnectionCustomHeadersOAuth2{"X-Custom-Header": "custom-value"}, opts.GetCustomHeaders())
	})

	t.Run("legacy string format", func(t *testing.T) {
		payload := `{
			"authorizationURL": "https://example.com/oauth/authorize",
			"tokenURL": "https://example.com/oauth/token",
			"customHeaders": "{\"X-Custom-Header\":\"custom-value\"}"
		}`

		var opts ConnectionOptionsOAuth2
		require.NoError(t, json.Unmarshal([]byte(payload), &opts))
		assert.Nil(t, opts.CustomHeaders)
		assert.Equal(t, ConnectionAuthorizationEndpoint("https://example.com/oauth/authorize"), opts.GetAuthorizationURL())
	})

	t.Run("null", func(t *testing.T) {
		payload := `{"customHeaders": null}`

		var opts ConnectionOptionsOAuth2
		require.NoError(t, json.Unmarshal([]byte(payload), &opts))
		assert.Nil(t, opts.CustomHeaders)
	})

	t.Run("absent", func(t *testing.T) {
		payload := `{"tokenURL": "https://example.com/oauth/token"}`

		var opts ConnectionOptionsOAuth2
		require.NoError(t, json.Unmarshal([]byte(payload), &opts))
		assert.Nil(t, opts.CustomHeaders)
	})

	t.Run("invalid type", func(t *testing.T) {
		payload := `{"customHeaders": 42}`

		var opts ConnectionOptionsOAuth2
		assert.Error(t, json.Unmarshal([]byte(payload), &opts))
	})

	t.Run("map format round trip", func(t *testing.T) {
		payload := `{"customHeaders": {"X-Custom-Header": "custom-value"}}`

		var opts ConnectionOptionsOAuth2
		require.NoError(t, json.Unmarshal([]byte(payload), &opts))

		remarshaled, err := json.Marshal(&opts)
		require.NoError(t, err)
		assert.JSONEq(t, payload, string(remarshaled))
	})
}
