package idtokenvalidator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

// ValidationOptions allows validating optional claims that might not always be in the ID token.
type ValidationOptions struct {
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
	jwks           *jwk.Cache
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

	alg, err := determineSigningAlgorithm(idTokenSigningAlg)
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

	if alg == jwa.RS256() {
		i.jwksURL = i.issuer + ".well-known/jwks.json"

		cache, err := jwk.NewCache(ctx, httprc.NewClient())
		if err != nil {
			return nil, err
		}
		i.jwks = cache

		registerOpts := []jwk.RegisterOption{}

		if i.httpClient != nil {
			registerOpts = append(registerOpts, jwk.WithHTTPClient(i.httpClient))
		}

		err = i.jwks.Register(ctx, i.jwksURL, registerOpts...)
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
func (i *IDTokenValidator) Validate(idToken string, optional ValidationOptions) error {
	validator := jwt.ValidatorFunc(func(_ context.Context, t jwt.Token) error {
		sub, hasSub := t.Subject()
		if !hasSub || sub == "" {
			return errors.New("sub claim must be a string present in the ID token")
		}

		if optional.Organization != "" {
			if strings.HasPrefix(optional.Organization, "org_") {
				var orgID string
				if err := t.Get("org_id", &orgID); err != nil {
					return errors.New("org_id claim must be a string present in the ID token")
				}

				if orgID != optional.Organization {
					return fmt.Errorf("org_id claim value mismatch in the ID token; expected \"%s\", found \"%s\"", optional.Organization, orgID)
				}
			} else {
				var orgName string
				if err := t.Get("org_name", &orgName); err != nil {
					return errors.New("org_name claim must be a string present in the ID token")
				}

				if orgName != strings.ToLower(optional.Organization) {
					return fmt.Errorf("org_name claim value mismatch in the ID token; expected \"%s\", found \"%s\"", optional.Organization, orgName)
				}
			}
		}

		now := time.Now()

		if optional.Nonce != "" {
			var nonce string
			if err := t.Get("nonce", &nonce); err != nil {
				return errors.New("nonce claim must be a string present in the ID token")
			}

			if nonce != optional.Nonce {
				return fmt.Errorf("nonce claim value mismatch in the ID token; expected \"%s\", found \"%s\"", optional.Nonce, nonce)
			}
		}

		aud, hasAud := t.Audience()
		if hasAud && len(aud) > 1 {
			var azp string
			if err := t.Get("azp", &azp); err != nil {
				return errors.New("azp claim must be a string present in the ID token when Audience (aud) claim has multiple values")
			}

			if azp != i.audience {
				return fmt.Errorf("azp claim mismatch in the ID token; expected \"%s\", found \"%s\"", i.audience, azp)
			}
		}

		if optional.MaxAge != 0 {
			var authTimeRaw any
			if err := t.Get("auth_time", &authTimeRaw); err != nil {
				return errors.New("auth_time claim must be a number present in the ID token when MaxAge is specified")
			}

			authTimeFloat, ok := authTimeRaw.(float64)
			if !ok {
				return errors.New("auth_time claim must be a number present in the ID token when MaxAge is specified")
			}

			at := time.Unix(int64(authTimeFloat), 0)
			validUntil := at.Add(optional.MaxAge).Add(i.clockTolerance)

			if now.After(validUntil) {
				return fmt.Errorf("auth_time claim in the ID token indicates that too much time has passed since the last end-user authentication. Current time %s is after last auth at %s", now, validUntil)
			}
		}

		return nil
	})

	decodedToken, err := jws.Parse([]byte(idToken))
	if err != nil {
		return err
	}

	headers := decodedToken.Signatures()[0].ProtectedHeaders()

	headerAlg, _ := headers.Algorithm()
	if headerAlg != jwa.HS256() && headerAlg != jwa.RS256() {
		return fmt.Errorf("signature algorithm \"%s\" is not supported. Expected the ID token to be signed with \"HS256\" or \"RS256\"", headerAlg)
	}

	if headerAlg != i.alg {
		return fmt.Errorf("unexpected signature algorithm; found \"%s\" but expected \"%s\"", headerAlg, i.alg)
	}

	// These options run in the order specified, so changing the order may change the errors returned.
	// Our own validator func should always be ran last.
	keyOpts := []jwt.ParseOption{
		jwt.WithValidate(true),
		jwt.WithAcceptableSkew(i.clockTolerance),
		jwt.WithRequiredClaim("aud"),
		jwt.WithRequiredClaim("sub"),
		jwt.WithRequiredClaim("iss"),
		jwt.WithRequiredClaim("iat"),
		jwt.WithAudience(i.audience),
		jwt.WithIssuer(i.issuer),
		jwt.WithValidator(validator),
	}

	if i.alg == jwa.HS256() {
		keyOpts = append(keyOpts, jwt.WithKey(i.alg, i.clientSecret))
	} else {
		cachedSet, err := i.jwks.CachedSet(i.jwksURL)
		if err != nil {
			return err
		}
		keyOpts = append(keyOpts, jwt.WithKeySet(cachedSet))
	}

	_, err = jwt.Parse([]byte(idToken), keyOpts...)

	return err
}

func determineSigningAlgorithm(alg string) (jwa.SignatureAlgorithm, error) {
	switch alg {
	case "HS256":
		return jwa.HS256(), nil
	case "RS256":
		return jwa.RS256(), nil
	default:
		var zero jwa.SignatureAlgorithm
		return zero, fmt.Errorf("unsupported algorithm provided: %q", alg)
	}
}
