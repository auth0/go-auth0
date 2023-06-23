package idtokenvalidator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// OptionalVerification allows validating optional claims that might not always be in the ID token.
type OptionalVerification struct {
	MaxAge       time.Duration
	Nonce        string
	Organization string
}

// IDTokenValidator is used to validate ID tokens retrieved from Auth0.
type IDTokenValidator struct {
	alg            jwa.SignatureAlgorithm
	audience       string
	clientSecret   []byte
	clockTolerance time.Duration
	httpClient     *http.Client
	issuer         string
	jwks           jwk.Cache
	jwksURL        string
}

// New creates and returns a new IDTokenValidator.
func New(
	ctx context.Context,
	domain string,
	clientID string,
	clientSecret string,
	idTokenSigningAlg string,
	opts ...Option,
) (*IDTokenValidator, error) {
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}
	alg, err := determineAlg(idTokenSigningAlg)
	if err != nil {
		return nil, err
	}

	i := &IDTokenValidator{
		clientSecret:   []byte(clientSecret),
		alg:            alg,
		clockTolerance: time.Minute,
		issuer:         "https://" + domain + "/",
		audience:       clientID,
	}

	for _, option := range opts {
		option(i)
	}

	if alg == jwa.RS256 {
		i.jwksURL = i.issuer + ".well-known/jwks.json"
		i.jwks = *jwk.NewCache(ctx)
		registerOpts := []jwk.RegisterOption{}
		if i.httpClient != nil {
			registerOpts = append(registerOpts, jwk.WithHTTPClient(i.httpClient))
		}

		err = i.jwks.Register(i.jwksURL, registerOpts...)
		if err != nil {
			return nil, err
		}

		_, err = i.jwks.Refresh(ctx, i.jwksURL)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}

// Validate validates the provided ID token against the values provided during the IDTokenValidator creation.
func (i *IDTokenValidator) Validate(idToken string, optional OptionalVerification) error {
	validator := jwt.ValidatorFunc(func(_ context.Context, t jwt.Token) jwt.ValidationError {
		if t.Subject() == "" {
			return jwt.NewValidationError(errors.New("sub claim must be a string present in the ID token"))
		}

		if optional.Organization != "" {
			orgID, exists := t.Get("org_id")

			if !exists {
				return jwt.NewValidationError(errors.New("org_id claim must be a string present in the ID token"))
			}

			if orgID != optional.Organization {
				return jwt.NewValidationError(fmt.Errorf("org_id claim value mismatch in the ID token; expected \"%s\", found \"%s\"", optional.Organization, orgID))
			}
		}

		now := time.Now()

		if optional.Nonce != "" {
			nonce, exists := t.Get("nonce")

			if !exists {
				return jwt.NewValidationError(errors.New("nonce claim must be a string present in the ID token"))
			}

			if nonce != optional.Nonce {
				return jwt.NewValidationError(fmt.Errorf("nonce claim value mismatch in the ID token; expected \"%s\", found \"%s\"", optional.Nonce, nonce))
			}
		}

		if len(t.Audience()) > 1 {
			azp, azpExists := t.Get("azp")

			if azpExists == false {
				return jwt.NewValidationError(errors.New("azp claim must be a string present in the ID token when Audience (aud) claim has multiple values"))
			}

			if azp != i.audience {
				return jwt.NewValidationError(fmt.Errorf("azp claim mismatch in the ID token; expected \"%s\", found \"%s\"", i.audience, azp))
			}
		}

		if optional.MaxAge != 0 {
			authTime, exists := t.Get("auth_time")

			if !exists {
				return jwt.NewValidationError(errors.New("auth_time claim must be a number present in the ID token when MaxAge is specified"))
			}

			if _, ok := authTime.(float64); !ok {
				return jwt.NewValidationError(errors.New("auth_time claim must be a number present in the ID token when MaxAge is specified"))
			}

			at := time.Unix(int64(authTime.(float64)), 0)
			validUntil := at.Add(optional.MaxAge).Add(i.clockTolerance)

			if now.After(validUntil) {
				return jwt.NewValidationError(fmt.Errorf("auth_time claim in the ID token indicates that too much time has passed since the last end-user authentication. Current time %s is after last auth at %s", now, validUntil))
			}
		}

		return nil
	})

	decodedToken, err := jws.Parse([]byte(idToken))
	if err != nil {
		return err
	}

	headers := decodedToken.Signatures()[0].ProtectedHeaders()

	if headers.Algorithm() != jwa.HS256 && headers.Algorithm() != jwa.RS256 {
		return fmt.Errorf("signature algorithm \"%s\" is not supported. Expected the ID Token to be sign with \"HS256\" or \"RS256\"", headers.Algorithm())
	}

	if headers.Algorithm() != i.alg {
		return fmt.Errorf("unexpected signature algorithm; found \"%s\" but expected \"%s\"", headers.Algorithm(), i.alg)
	}

	keyOpts := []jwt.ParseOption{
		jwt.WithValidate(true),
		jwt.WithAcceptableSkew(i.clockTolerance),
		jwt.WithAudience(i.audience),
		jwt.WithIssuer(i.issuer),
		jwt.WithValidator(validator),
	}

	if i.alg == jwa.HS256 {
		keyOpts = append(keyOpts, jwt.WithKey(i.alg, i.clientSecret))
	} else {
		keyOpts = append(keyOpts, jwt.WithKeySet(jwk.NewCachedSet(&i.jwks, i.jwksURL)))
	}

	_, err = jwt.Parse([]byte(idToken), keyOpts...)
	return err
}

func determineAlg(alg string) (jwa.SignatureAlgorithm, error) {
	switch alg {
	case "HS256":
		return jwa.HS256, nil
	case "RS256":
		return jwa.RS256, nil
	default:
		return "", fmt.Errorf("Unsupported algorithm %s provided", alg)
	}
}
