package management

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tink-crypto/tink-go/v2/kwp/subtle"

	"github.com/auth0/go-auth0"
)

func TestEncryptionKeyManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)
	givenEncryptionKey := &EncryptionKey{
		Type: auth0.String("customer-provided-root-key"),
	}
	err := api.EncryptionKey.Create(context.Background(), givenEncryptionKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEncryptionKey.GetKID())
	cleanUpEncryptionKey(t, givenEncryptionKey.GetKID())
}

func TestEncryptionKeyManager_List(t *testing.T) {
	configureHTTPTestRecordings(t)
	key := givenEncryptionKey(t)
	keyList, err := api.EncryptionKey.List(context.Background(), PerPage(50), Page(0))
	assert.NoError(t, err)
	assert.Contains(t, keyList.Keys, key)
}

func TestEncryptionKeyManager_Read(t *testing.T) {
	configureHTTPTestRecordings(t)
	key := givenEncryptionKey(t)
	readKey, err := api.EncryptionKey.Read(context.Background(), key.GetKID())
	assert.NoError(t, err)
	assert.Equal(t, key, readKey)
}

func TestEncryptionKeyManager_Rekey(t *testing.T) {
	configureHTTPTestRecordings(t)
	oldKeyList, err := api.EncryptionKey.List(context.Background(), PerPage(50), Page(0))
	assert.NoError(t, err)
	assert.NotEmpty(t, oldKeyList.Keys)
	var oldKey, newKey *EncryptionKey
	for _, key := range oldKeyList.Keys {
		if key.GetState() == "active" && key.GetType() == "tenant-master-key" {
			oldKey = key
		}
	}
	err = api.EncryptionKey.Rekey(context.Background())
	assert.NoError(t, err)
	keyList, err := api.EncryptionKey.List(context.Background(), PerPage(50), Page(0))
	assert.NoError(t, err)
	assert.NotEmpty(t, keyList.Keys)
	for _, key := range keyList.Keys {
		if key.GetState() == "active" && key.GetType() == "tenant-master-key" {
			newKey = key
		}
	}
	assert.NotEqual(t, oldKey.GetKID(), newKey.GetKID())
	assert.NotEqual(t, keyList.Keys, oldKeyList.Keys)
}

func TestEncryptionKeyManager_Delete(t *testing.T) {
	configureHTTPTestRecordings(t)
	key := givenEncryptionKey(t)
	err := api.EncryptionKey.Delete(context.Background(), key.GetKID())
	assert.NoError(t, err)
	keyRead, err := api.EncryptionKey.Read(context.Background(), key.GetKID())
	assert.NoError(t, err)
	assert.Equal(t, keyRead.GetState(), "destroyed")
}

func TestEncryptionKeyManager_CreatePublicWrappingKey(t *testing.T) {
	configureHTTPTestRecordings(t)
	key := givenEncryptionKey(t)
	wrappingKey, err := api.EncryptionKey.CreatePublicWrappingKey(context.Background(), key.GetKID())
	assert.NoError(t, err)
	assert.NotEmpty(t, wrappingKey.GetPublicKey())
}

func TestEncryptionKeyManager_ImportWrappedKey(t *testing.T) {
	configureHTTPTestRecordings(t)
	key := givenEncryptionKey(t)
	wrappingKey, err := api.EncryptionKey.CreatePublicWrappingKey(context.Background(), key.GetKID())
	assert.NoError(t, err)
	assert.NotEmpty(t, wrappingKey.GetPublicKey())
	wrappedKeyStr, err := createAWSWrappedCiphertext(wrappingKey.GetPublicKey())
	assert.NoError(t, err)

	key.WrappedKey = &wrappedKeyStr

	err = api.EncryptionKey.ImportWrappedKey(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, key.GetType(), "customer-provided-root-key")
	assert.Equal(t, key.GetState(), "active")
}

func givenEncryptionKey(t *testing.T) *EncryptionKey {
	t.Helper()
	givenEncryptionKey := &EncryptionKey{
		Type: auth0.String("customer-provided-root-key"),
	}
	err := api.EncryptionKey.Create(context.Background(), givenEncryptionKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEncryptionKey.GetKID())
	t.Cleanup(func() {
		cleanUpEncryptionKey(t, givenEncryptionKey.GetKID())
	})
	return givenEncryptionKey
}

func cleanUpEncryptionKey(t *testing.T, kid string) {
	t.Helper()
	err := api.EncryptionKey.Delete(context.Background(), kid)
	assert.NoError(t, err)
}

func createAWSWrappedCiphertext(publicKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to decode public key PEM")
	}

	// Parse the public key
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %w", err)
	}

	// Ensure the public key is of type *rsa.PublicKey
	publicRSAKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("public key is not of type *rsa.PublicKey")
	}

	// Generate a 256-bit (32-byte) ephemeral key
	ephemeralKey := make([]byte, 32)
	if _, err := rand.Read(ephemeralKey); err != nil {
		return "", fmt.Errorf("failed to generate ephemeral key: %w", err)
	}

	// Generate a 256-bit (32-byte) plaintext key
	plaintextKey := make([]byte, 32)
	if _, err := rand.Read(plaintextKey); err != nil {
		return "", fmt.Errorf("failed to generate plaintext key: %w", err)
	}

	// Wrap the ephemeral key using RSA-OAEP with SHA-256
	wrappedEphemeralKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicRSAKey, ephemeralKey, nil)
	if err != nil {
		return "", fmt.Errorf("failed to wrap ephemeral key: %w", err)
	}

	// Create a KWP (Key Wrapping with Padding) instance using the ephemeral key
	kwp, err := subtle.NewKWP(ephemeralKey)
	if err != nil {
		return "", fmt.Errorf("failed to create KWP instance: %w", err)
	}

	// Wrap the plaintext key using KWP
	wrappedTargetKey, err := kwp.Wrap(plaintextKey)
	if err != nil {
		return "", fmt.Errorf("failed to wrap target key using KWP: %w", err)
	}

	// Return the concatenation of the wrapped ephemeral key and the wrapped plaintext key
	wrappedEphemeralKey = append(wrappedEphemeralKey, wrappedTargetKey...)
	cipherBytes := wrappedEphemeralKey

	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}
