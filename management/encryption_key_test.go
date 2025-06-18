package management

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/auth0/go-auth0"
)

// Constants for wrapping sizes and parameters.
const (
	minWrapSize = 16
	maxWrapSize = 8192
	roundCount  = 6
	ivPrefix    = uint32(0xA65959A6)
)

// kwpImpl is a Key Wrapping with Padding implementation.
type kwpImpl struct {
	block cipher.Block
}

func TestEncryptionKeyManager_Create(t *testing.T) {
	configureHTTPTestRecordings(t)

	givenEncryptionKey := &EncryptionKey{
		Type: auth0.String("customer-provided-root-key"),
	}
	err := api.EncryptionKey.Create(context.Background(), givenEncryptionKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, givenEncryptionKey.GetKID())
	t.Cleanup(func() {
		cleanUpEncryptionKey(t, givenEncryptionKey.GetKID())
	})
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
			break
		}
	}

	assert.NotNil(t, oldKey)

	err = api.EncryptionKey.Rekey(context.Background())
	assert.NoError(t, err)

	keyList, err := api.EncryptionKey.List(context.Background(), PerPage(50), Page(0))
	assert.NoError(t, err)
	assert.NotEmpty(t, keyList.Keys)

	for _, key := range keyList.Keys {
		if key.GetState() == "active" && key.GetType() == "tenant-master-key" {
			newKey = key
			break
		}
	}

	assert.NotNil(t, newKey)

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

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %w", err)
	}

	publicRSAKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("public key is not of type *rsa.PublicKey")
	}

	ephemeralKey := make([]byte, 32)
	if _, err := rand.Read(ephemeralKey); err != nil {
		return "", fmt.Errorf("failed to generate ephemeral key: %w", err)
	}

	plaintextKey := make([]byte, 32)
	if _, err := rand.Read(plaintextKey); err != nil {
		return "", fmt.Errorf("failed to generate plaintext key: %w", err)
	}

	wrappedEphemeralKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicRSAKey, ephemeralKey, nil)
	if err != nil {
		return "", fmt.Errorf("failed to wrap ephemeral key: %w", err)
	}

	kwp, err := newKWP(ephemeralKey)
	if err != nil {
		return "", fmt.Errorf("failed to create KWP instance: %w", err)
	}

	wrappedTargetKey, err := kwp.wrap(plaintextKey)
	if err != nil {
		return "", fmt.Errorf("failed to wrap target key using KWP: %w", err)
	}

	wrappedEphemeralKey = append(wrappedEphemeralKey, wrappedTargetKey...)

	return base64.StdEncoding.EncodeToString(wrappedEphemeralKey), nil
}

func newKWP(wrappingKey []byte) (*kwpImpl, error) {
	switch len(wrappingKey) {
	case 16, 32:
		block, err := aes.NewCipher(wrappingKey)
		if err != nil {
			return nil, fmt.Errorf("kwp: error building AES cipher: %v", err)
		}

		return &kwpImpl{block: block}, nil
	default:
		return nil, fmt.Errorf("kwp: invalid AES key size; want 16 or 32, got %d", len(wrappingKey))
	}
}

func wrappingSize(inputSize int) int {
	paddingSize := 7 - (inputSize+7)%8
	return inputSize + paddingSize + 8
}

func (kwp *kwpImpl) computeW(iv, key []byte) ([]byte, error) {
	if len(key) <= 8 || len(key) > math.MaxInt32-16 || len(iv) != 8 {
		return nil, fmt.Errorf("kwp: computeW called with invalid parameters")
	}

	data := make([]byte, wrappingSize(len(key)))
	copy(data, iv)
	copy(data[8:], key)
	blockCount := len(data)/8 - 1

	buf := make([]byte, 16)
	copy(buf, data[:8])

	for i := 0; i < roundCount; i++ {
		for j := 0; j < blockCount; j++ {
			copy(buf[8:], data[8*(j+1):])
			kwp.block.Encrypt(buf, buf)

			roundConst := uint(i*blockCount + j + 1)
			for b := 0; b < 4; b++ {
				buf[7-b] ^= byte(roundConst & 0xFF)
				roundConst >>= 8
			}

			copy(data[8*(j+1):], buf[8:])
		}
	}

	copy(data[:8], buf)

	return data, nil
}

func (kwp *kwpImpl) wrap(data []byte) ([]byte, error) {
	if len(data) < minWrapSize {
		return nil, fmt.Errorf("kwp: key size to wrap too small")
	}

	if len(data) > maxWrapSize {
		return nil, fmt.Errorf("kwp: key size to wrap too large")
	}

	iv := make([]byte, 8)
	binary.BigEndian.PutUint32(iv, ivPrefix)
	binary.BigEndian.PutUint32(iv[4:], uint32(len(data)))

	return kwp.computeW(iv, data)
}
