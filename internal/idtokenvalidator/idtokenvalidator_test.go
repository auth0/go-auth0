package idtokenvalidator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/assert"
)

var (
	jwtDomain       = "not-a-tenant.com"
	jwtURL          = "https://" + jwtDomain + "/"
	jwtClientID     = "test-client-id"
	jwtClientSecret = "test-client-secret"
	jwtPublicKey    = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8foXPIpkeLKAVfVg/W0X
steFas2XwrxAGG0lnLS3mc/cYc/pD/plsR779O8It/2YmHFWIDmCIcW57boDae/K
AhVBLHUa3ato7h5agbY2mKSDUEjqjWilAbdyUZDz8US8ocAmehyVWMuVqeGxunPH
opm4JQ2OGcE31MbtcJN07zCa/R/LUi8KMeuujQ6cceIGupCdOsK6JkoUB2wkvFpU
CiOwqTG51Eq4DSTukDr7tDfe0s5e1MBxVUxLkrw7zBrlDxPgZ+M260FUlRqOKKsk
3IUke/vcQac6t+js1zOs0mapqLybkszGwl2wY0JaOLOtcL5zi4U9w/GeOHOVROdn
6wIDAQAB
-----END PUBLIC KEY-----`
	jwtPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDx+hc8imR4soBV
9WD9bRey14VqzZfCvEAYbSWctLeZz9xhz+kP+mWxHvv07wi3/ZiYcVYgOYIhxbnt
ugNp78oCFUEsdRrdq2juHlqBtjaYpINQSOqNaKUBt3JRkPPxRLyhwCZ6HJVYy5Wp
4bG6c8eimbglDY4ZwTfUxu1wk3TvMJr9H8tSLwox666NDpxx4ga6kJ06wromShQH
bCS8WlQKI7CpMbnUSrgNJO6QOvu0N97Szl7UwHFVTEuSvDvMGuUPE+Bn4zbrQVSV
Go4oqyTchSR7+9xBpzq36OzXM6zSZqmovJuSzMbCXbBjQlo4s61wvnOLhT3D8Z44
c5VE52frAgMBAAECggEALCx3qXmqNc6AVzDgb+NGfEOT+5dkqQwst0jVoPHswouL
s998sIoJnngFjwVEFjKZdNrb2i4lb3zlIFzg2qoHurGeoDsQmH7+PNoVs7BL7zm5
LyLgjsgXt2SB3hoULmtZ9D1byNcG/JrNy6GEDIGuZCSj1T/QPStkwdc+6VpB8pgW
E8D7jCt40Tik2neYQkDnY775kGAHGWEqpdPCwm+KOnuE1fHx/jk38lmUgYNjKq0h
JK6Ncjen1X+ZsYfGx4dALWG4cqo3lE0YXXuHuvjJV3aVfzH8t7W4fuZ4+8xvdhhV
F4br5FimWLbTe2qT4lSpadkbLm3aBlSUR7eAP0BlwQKBgQD5ayZpP5OMp1zfa4hA
fM8nVUEaVLkRwFK5NChfjHGiaye2RjrnIorXMsFxXjEscgTn2Ux9CgcBhp1fTBhy
6cmhkp1talAIqLBivNQJT0YTfA+uHrHTTyMfEUgsMzPiiAg7FV7BCG6xd/nsk3yg
ZUfoXefrhq9LIHsJx7cK12VViQKBgQD4XKvwYmX5t7fZFBPd7dv5ZrcMHQnBMHd7
is3QhgyKuEgVDzKQ9SA004I9iSvcI3dE/npj31P39N5bbuvYTh4WR/SR4VvXavNG
AqUR7wm8jTlbiWEPgF9MxC24zaa07Kbxs+P8XT/7wWuijf6+baSFgxQMb80fUArv
7guKikCo0wKBgCUn3DIDoZRrfj9eQo7wyN9gKPGmO2e0kd47MeSCBI+gjOrvbWjv
UWWbjwu3b3Xiim6LhYR/EOoeRqViraW4xCvIrqEVHFUd5CDhZmj4oUTXz3It6mnD
OUUwiuLiwdD2WNuMZHA3NF5FtDqVAhTW4a5xBtKkXsq/TPT5BoCb8+GZAoGAUWAD
0gpbgTuJ2G10qPWDaq8V8Lke9haMP4VWNCmHuHfy3juRhN9cAxL+DG2CWmmgbZG3
xjtpRsgLhwfL7J6DyyceYiHltqpLNTgun7ajiQz4qx5TGAImt39bv75aDdOwS2d2
nrxq93EDdEp0Gi7QhhJRolWLbuQKAV0MmQL9dpMCgYEA5+ug3CDI/jyTHG4ZEVoG
qmIg7QoHrVEmZrvCMiFw8bbuBvoMnvu1o1zfvAkNrDfibZyxYKHzSqgeVPQShvLa
P6JCu67ieCGP8C8CMFiQhJ9n4sYGnkzkz67NpkHSzDPA6DfvG4pYuvBQRIefnhYh
IDGpghhKHMV2DAyzeM4cDU8=
-----END PRIVATE KEY-----`
	defaultJWTArgs = jwtArgs{
		issuer: jwtURL,
		sub:    "me",
		exp:    time.Now().Add(3 * time.Second),
		iat:    time.Now(),
		aud:    []string{jwtClientID},
	}
)

func TestIDTokenValidation(t *testing.T) {
	t.Run("should error if passed unsupported idTokenSigningAlg", func(t *testing.T) {
		_, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS512")
		assert.ErrorContains(t, err, "unsupported algorithm")
	})

	t.Run("validates a token signed by a secret", func(t *testing.T) {
		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.NoError(t, err)
	})

	t.Run("validates a token signed by a private key", func(t *testing.T) {
		args := defaultJWTArgs
		args.privateKey = jwtPrivateKey
		args.publicKey = jwtPublicKey

		token, server, err := givenAJWT(t, args)
		assert.NoError(t, err)

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		validator, err := New(context.Background(), URL.Host, jwtClientID, jwtClientSecret, "RS256", WithHTTPClient(server.Client()))
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.NoError(t, err)
	})

	t.Run("validates the signing alg provided is the one used", func(t *testing.T) {
		args := defaultJWTArgs
		args.privateKey = jwtPrivateKey
		args.publicKey = jwtPublicKey

		token, server, err := givenAJWT(t, args)
		assert.NoError(t, err)

		URL, err := url.Parse(server.URL)
		assert.NoError(t, err)

		validator, err := New(context.Background(), URL.Host, jwtClientID, jwtClientSecret, "HS256", WithHTTPClient(server.Client()))
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "unexpected signature algorithm; found")
	})

	t.Run("validates the jwt signing algorithm is supported", func(t *testing.T) {
		builder := jwt.NewBuilder().
			Issuer(defaultJWTArgs.issuer).
			Subject(defaultJWTArgs.sub).
			Audience(defaultJWTArgs.aud).
			Expiration(defaultJWTArgs.exp).
			IssuedAt(defaultJWTArgs.iat)

		token, err := builder.Build()
		assert.NoError(t, err)

		jwtPayload, err := jwt.Sign(token, jwt.WithKey(jwa.HS512, []byte(jwtClientSecret)))
		assert.NoError(t, err)

		validator, err := New(context.Background(), jwtDomain, jwtClientSecret, jwtClientID, "HS256")
		assert.NoError(t, err)

		err = validator.Validate(string(jwtPayload), ValidationOptions{})
		assert.ErrorContains(t, err, "signature algorithm \"HS512\" is not supported")
	})

	t.Run("rejects if azp is not present and more than one audience is provided", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.aud = []string{jwtClientID, "bar"}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "azp claim must be a string present in the ID token")
	})

	t.Run("rejects if azp is unknown", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.aud = []string{jwtClientID, "bar"}
		args.payload = map[string]interface{}{
			"azp": "invalid-value",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "azp claim mismatch in the ID token")
	})

	t.Run("verifies the audience when azp is present", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.aud = []string{jwtClientID, "bar"}
		args.payload = map[string]interface{}{
			"azp": jwtClientID,
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.NoError(t, err)
	})

	t.Run("verifies when provided a nonce", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"nonce": "test-nonce",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Nonce: "test-nonce"})
		assert.NoError(t, err)
	})

	t.Run("errors when provided a nonce and it doesn't exist", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Nonce: "test-nonce"})
		assert.ErrorContains(t, err, "nonce claim must be a string present in the ID token")
	})

	t.Run("errors when provided a nonce and it doesn't match", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"nonce": "different-nonce",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Nonce: "test-nonce"})
		assert.ErrorContains(t, err, "nonce claim value mismatch in the ID token;")
	})

	t.Run("verifies iss exists", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		// Minimal JWT with no time based values and is missing iss. This can't be achieved by building using jwx
		err = validator.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYXVkIjoidGVzdC1jbGllbnQtaWQifQ.52jn1t1Jh_HS2rg1xbyvO1HdNXXVF22e7D1g9TGeWJI", ValidationOptions{})
		assert.ErrorContains(t, err, "\"iss\" not satisfied: required claim not found")
	})

	t.Run("verifies iss is valid", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.issuer = "invalid-iss"

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "\"iss\" not satisfied:")

		args.issuer = ""

		token, _, err = givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "\"iss\" not satisfied:")
	})

	t.Run("verifies sub exists", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		// Minimal JWT with no time based values and is missing iss. This can't be achieved by building using jwx
		err = validator.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhdWQiOiJ0ZXN0LWNsaWVudC1pZCIsImlzcyI6Imh0dHBzOi8vbm90LWEtdGVuYW50LmNvbS8ifQ.LrVdUAUzsADeOz2TBt4rsiQW93knty1SVEp9eYG48SE", ValidationOptions{})
		assert.ErrorContains(t, err, "\"sub\" not satisfied: required claim not found")
	})

	t.Run("verifies sub is valid", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"sub": "",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "sub claim must be a string present in the ID token")
	})

	t.Run("verifies aud exists", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.aud = nil

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "\"aud\" not satisfied: required claim not found")
	})

	t.Run("verifies aud is valid", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.aud = []string{}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "\"aud\" not satisfied")
	})

	t.Run("allows clock skew on iat", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256", WithClockTolerance(10*time.Second))
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"iat": time.Now().Add(5 * time.Second),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.NoError(t, err)
	})

	t.Run("verifies exp is in the future", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256", WithClockTolerance(time.Second))
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"exp": time.Now().Add(-10 * time.Second),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.ErrorContains(t, err, "\"exp\" not satisfied")
	})

	t.Run("allows clock skew on exp", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256", WithClockTolerance(10*time.Second))
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"exp": time.Now().Add(-5 * time.Second),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{})
		assert.NoError(t, err)
	})

	t.Run("passes when auth_time is within max_age", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"auth_time": time.Now().Add(-100 * time.Second).Unix(),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{MaxAge: 200 * time.Second})
		assert.NoError(t, err)
	})

	t.Run("errors when auth_time is beyond max_age", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"auth_time": time.Now().Add(-500 * time.Second).Unix(),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{MaxAge: 200 * time.Second})
		assert.ErrorContains(t, err, "auth_time claim in the ID token indicates that too much time has passed since")
	})

	t.Run("allows clock skew on auth_time", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256", WithClockTolerance(100*time.Second))
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"auth_time": time.Now().Add(-250 * time.Second).Unix(),
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{MaxAge: 200 * time.Second})
		assert.NoError(t, err)
	})

	t.Run("verifies auth_time is a number that exists when max_age is passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		// Test when value doesn't exist
		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{MaxAge: 200 * time.Second})
		assert.ErrorContains(t, err, "auth_time claim must be a number present in the ID token")

		// Test invalid type
		args.payload = map[string]interface{}{
			"auth_time": "test-value",
		}
		token, _, err = givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{MaxAge: 200 * time.Second})
		assert.ErrorContains(t, err, "auth_time claim must be a number present in the ID token")
	})

	t.Run("passes if org_id is valid when organization passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"org_id": "org_12345",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "org_12345"})
		assert.NoError(t, err)
	})

	t.Run("verifies org_id exists when organization is passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "org_12345"})
		assert.ErrorContains(t, err, "org_id claim must be a string present in the ID token")
	})

	t.Run("verifies org_id matches when organization is passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"org_id": "invalid-value",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "org_12345"})
		assert.ErrorContains(t, err, "org_id claim value mismatch in the ID token")
	})

	t.Run("passes if org_name is valid when organization passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"org_name": "my-org",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "my-org"})
		assert.NoError(t, err)
	})

	t.Run("verifies org_name exists when organization is passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "my-org"})
		assert.ErrorContains(t, err, "org_name claim must be a string present in the ID token")
	})

	t.Run("verifies org_name matches when organization is passed", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"org_name": "invalid-value",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "my-org"})
		assert.ErrorContains(t, err, "org_name claim value mismatch in the ID token")
	})

	t.Run("verifies org_name matches when organization is passed, case insensitive", func(t *testing.T) {
		validator, err := New(context.Background(), jwtDomain, jwtClientID, jwtClientSecret, "HS256")
		assert.NoError(t, err)

		args := defaultJWTArgs
		args.clientSecret = jwtClientSecret
		args.payload = map[string]interface{}{
			"org_name": "my-org",
		}

		token, _, err := givenAJWT(t, args)
		assert.NoError(t, err)

		err = validator.Validate(token, ValidationOptions{Organization: "my-Org"})
		assert.NoError(t, err)
	})
}

type jwtArgs struct {
	payload      map[string]interface{}
	clientSecret string
	privateKey   string
	publicKey    string
	issuer       string
	sub          string
	exp          time.Time
	iat          time.Time
	aud          []string
}

func givenAJWT(t *testing.T, args jwtArgs) (string, *httptest.Server, error) {
	t.Helper()

	alg, key, server, err := configureSigning(t, args)
	if err != nil {
		return "", nil, err
	}

	builder := jwt.NewBuilder().
		Issuer(args.issuer).
		Subject(args.sub).
		Audience(args.aud).
		Expiration(args.exp).
		IssuedAt(args.iat)

	for claim, value := range args.payload {
		builder.Claim(claim, value)
	}

	if server != nil {
		builder.Issuer(server.URL + "/")
	}

	token, err := builder.Build()

	if err != nil {
		return "", nil, err
	}

	b, err := jwt.Sign(token, jwt.WithKey(alg, key))
	if err != nil {
		return "", nil, err
	}

	return string(b), server, nil
}

func configureSigning(t *testing.T, args jwtArgs) (jwa.SignatureAlgorithm, jwk.Key, *httptest.Server, error) {
	t.Helper()

	if args.clientSecret != "" {
		raw, err := jwk.FromRaw([]byte(args.clientSecret))
		return jwa.HS256, raw, nil, err
	}

	publicKey, err := jwk.ParseKey([]byte(args.publicKey), jwk.WithPEM(true))
	if err != nil {
		return jwa.RS256, nil, nil, err
	}

	err = publicKey.Set(jwk.KeyIDKey, "1")
	if err != nil {
		return jwa.RS256, nil, nil, err
	}

	err = publicKey.Set(jwk.AlgorithmKey, jwa.RS256)
	if err != nil {
		return jwa.RS256, nil, nil, err
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		b, err := json.Marshal(publicKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if _, err := fmt.Fprintf(w, `{"keys": [%s] }`, b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	s := httptest.NewTLSServer(h)

	t.Cleanup(func() {
		defer s.Close()
	})

	privateKey, err := jwk.ParseKey([]byte(args.privateKey), jwk.WithPEM(true))
	if err != nil {
		return jwa.RS256, nil, nil, err
	}

	err = privateKey.Set(jwk.KeyIDKey, "1")
	if err != nil {
		return jwa.RS256, nil, nil, err
	}

	err = privateKey.Set(jwk.AlgorithmKey, jwa.RS256)

	return jwa.RS256, privateKey, s, err
}
