package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigningKey(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		setupHTTPRecordings(t)

		signingKeys, err := m.SigningKey.List()
		assert.NoError(t, err)
		assert.NotEmpty(t, signingKeys, "expected at least one key to be returned")
	})

	t.Run("Read", func(t *testing.T) {
		setupHTTPRecordings(t)

		signingKeys, err := m.SigningKey.List()
		assert.NoError(t, err)

		signingKey, err := m.SigningKey.Read(signingKeys[0].GetKID())
		assert.NoError(t, err)
		assert.NotEmpty(t, signingKey)
	})

	t.Run("Rotate", func(t *testing.T) {
		setupHTTPRecordings(t)

		signingKey, err := m.SigningKey.Rotate()
		assert.NoError(t, err)
		assert.NotEmpty(t, signingKey)
	})

	t.Run("Revoke", func(t *testing.T) {
		setupHTTPRecordings(t)

		// Our last test revokes the key used to sign the token we're currently
		// using, so we need to re-authenticate so that subsequent tests still work.
		t.Cleanup(func() {
			initTestManagement()
		})

		signingKeys, err := m.SigningKey.List()
		assert.NoError(t, err)

		var signingKey *SigningKey
		for _, key := range signingKeys {
			if key.GetPrevious() {
				signingKey = key
				break
			}
		}

		assert.NotNil(t, signingKey, "previous key not found, nothing to revoke")

		revokedKey, err := m.SigningKey.Revoke(signingKey.GetKID())
		assert.NoError(t, err)
		assert.NotEmpty(t, revokedKey)
	})
}
