package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/auth0/go-auth0/internal/client"
	"github.com/auth0/go-auth0/internal/idtokenvalidator"
)

// UserInfoResponse defines the response from the user info API.
type UserInfoResponse struct {
	// Unknown claims in the response object that are not defined in this struct will be stored here.
	AdditionalClaims map[string]interface{} `json:"-"`
	// The user's preferred postal address.
	Address *UserAddress `json:"address,omitempty"`
	// The user's birthday, represented as an ISO 8601:2004 [ISO8601‑2004] YYYY-MM-DD format.
	Birthdate string `json:"birthdate,omitempty"`
	// The user's preferred email address.
	Email string `json:"email,omitempty"`
	// Whether the user's email address has been verified or not.
	EmailVerified bool `json:"email_verified,omitempty"`
	// Surname(s) or last name(s) of the user.
	FamilyName string `json:"family_name,omitempty"`
	// The user's gender.
	Gender string `json:"gender,omitempty"`
	// Given name(s) or first name(s) of the user.
	GivenName string `json:"given_name,omitempty"`
	// The user's locale, represented as a BCP47 language tag.
	Locale string `json:"locale,omitempty"`
	// Middle name(s) of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// Full name of the user in displayable form including all name parts, possibly including titles and
	// suffixes, ordered according to the user's locale and preferences.
	Name string `json:"name,omitempty"`
	// Casual name of the user that may or may not be the same as FirstName.
	Nickname string `json:"nickname,omitempty"`
	// The user's preferred telephone number.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Whether the user's phone number has been verified or not.
	PhoneNumberVerified bool `json:"phone_number_verified,omitempty"`
	// URL of the user's profile picture.
	Picture string `json:"picture,omitempty"`
	// Shorthand name by which the user wishes to be referred by, such as janedoe or j.doe.
	PreferredUsername string `json:"preferred_username,omitempty"`
	// URL of the user's profile page.
	Profile string `json:"profile,omitempty"`
	// The Auth0 user identifier. This is unique to each user.
	Sub string `json:"sub,omitempty"`
	// Time and date the user's information was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// URL of the user's web page or blog.
	Website string `json:"website,omitempty"`
	// User's time zone as a "tz database name".
	ZoneInformation string `json:"zoneinfo,omitempty"`
}

// UserAddress defines a user's address.
type UserAddress struct {
	// Country component of the address.
	Country string `json:"country,omitempty"`
	// Full mailing address, formatted for display or use on a mailing label.
	Formatted string `json:"formatted,omitempty"`
	// City or locality component of the address.
	Locality string `json:"locality,omitempty"`
	// Zip or postal code component of the address.
	PostalCode string `json:"postal_code,omitempty"`
	// State, province, prefecture, or region component of the address.
	Region string `json:"region,omitempty"`
	// Full street address component, which may include house number, street name, Post Office Box,
	// and multi-line extended street address information.
	StreetAddress string `json:"street_address,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// It is required to handle the mapping of unknown claims from the main response body into the `AdditionalClaims`
// field on the struct.
func (u *UserInfoResponse) UnmarshalJSON(b []byte) error {
	type userInfoWrapper UserInfoResponse

	ui := userInfoWrapper{}

	err := json.Unmarshal(b, &ui)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &(ui.AdditionalClaims))
	if err != nil {
		return err
	}

	typ := reflect.TypeOf(ui)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			delete(ui.AdditionalClaims, jsonTag)
		}
	}

	*u = UserInfoResponse(ui)

	return nil
}

// Authentication is the auth client.
type Authentication struct {
	Database     *Database
	MFA          *MFA
	OAuth        *OAuth
	Passwordless *Passwordless
	CIBA         *CIBA

	auth0ClientInfo           *client.Auth0ClientInfo
	basePath                  string
	common                    manager
	clientID                  string
	clientSecret              string
	clientAssertionSigningKey string
	clientAssertionSigningAlg string
	idTokenClockTolerance     time.Duration
	debug                     bool
	http                      *http.Client
	idTokenSigningAlg         string
	idTokenValidator          *idtokenvalidator.IDTokenValidator
	url                       *url.URL
	retryStrategy             client.RetryOptions
}

type manager struct {
	authentication *Authentication
}

// New creates an auth client.
func New(ctx context.Context, domain string, options ...Option) (*Authentication, error) {
	// Ignore the scheme if it was defined in the domain variable, then prefix
	// with https as it's the only scheme supported by the Auth0 API.
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}

	domain = "https://" + domain

	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}

	a := &Authentication{
		url:               u,
		basePath:          "",
		debug:             false,
		http:              http.DefaultClient,
		auth0ClientInfo:   client.DefaultAuth0ClientInfo,
		idTokenSigningAlg: "RS256",
		retryStrategy:     client.DefaultRetryOptions,
	}

	for _, option := range options {
		option(a)
	}

	a.http = client.Wrap(
		a.http,
		client.WithDebug(a.debug),
		client.WithAuth0ClientInfo(a.auth0ClientInfo),
		client.WithRetries(a.retryStrategy),
	)

	a.common.authentication = a
	a.Database = (*Database)(&a.common)
	a.MFA = (*MFA)(&a.common)
	a.OAuth = (*OAuth)(&a.common)
	a.Passwordless = (*Passwordless)(&a.common)
	a.CIBA = (*CIBA)(&a.common)

	validatorOpts := []idtokenvalidator.Option{}

	if a.http != http.DefaultClient {
		validatorOpts = append(validatorOpts, idtokenvalidator.WithHTTPClient(a.http))
	}

	if a.idTokenClockTolerance.Seconds() != 0 {
		validatorOpts = append(validatorOpts, idtokenvalidator.WithClockTolerance(a.idTokenClockTolerance))
	}

	validator, err := idtokenvalidator.New(
		ctx,
		domain,
		a.clientID,
		a.clientSecret,
		a.idTokenSigningAlg,
		validatorOpts...,
	)

	if err != nil {
		return nil, err
	}

	a.idTokenValidator = validator

	return a, nil
}

// UserInfo returns a user's profile using the access token obtained during login.
//
// This endpoint will work only if `openid` was granted as a scope for the access token. The user profile
// information included in the response depends on the scopes requested. For example, a scope of just openid
// may return less information than a scope of `openid profile email`.
//
// See: https://auth0.com/docs/api/authentication?http#get-user-info
func (a *Authentication) UserInfo(ctx context.Context, accessToken string, opts ...RequestOption) (user *UserInfoResponse, err error) {
	opts = append(opts, Header("Authorization", "Bearer "+accessToken))
	err = a.Request(ctx, "GET", a.URI("userinfo"), nil, &user, opts...)

	return
}

// Helper for adding values to a url.Values instance if they are not empty.
func addIfNotEmpty(key string, value string, qs url.Values) {
	if value != "" {
		qs.Set(key, value)
	}
}

// Helper for enforcing that required values are set.
func check(errors *[]string, key string, c bool) {
	if !c {
		*errors = append(*errors, key)
	}
}

// Helper for adding client authentication into a url.Values instance.
func (a *Authentication) addClientAuthenticationToURLValues(params oauth.ClientAuthentication, body url.Values, required bool) error {
	clientID := params.ClientID
	if params.ClientID == "" {
		clientID = a.clientID
	}

	body.Set("client_id", clientID)

	clientSecret := params.ClientSecret
	if params.ClientSecret == "" {
		clientSecret = a.clientSecret
	}

	switch {
	case a.clientAssertionSigningKey != "" && a.clientAssertionSigningAlg != "":
		alg, err := client.DetermineSigningAlgorithm(a.clientAssertionSigningAlg)
		if err != nil {
			return err
		}

		// Using the improved createClientAssertion with a standard lifetime
		clientAssertion, err := client.CreateClientAssertion(
			alg,
			a.clientAssertionSigningKey,
			clientID,
			a.url.JoinPath("/").String(),
		)
		if err != nil {
			return err
		}

		body.Set("client_assertion", clientAssertion)
		body.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	case params.ClientAssertion != "" && params.ClientAssertionType != "":
		body.Set("client_assertion", params.ClientAssertion)
		body.Set("client_assertion_type", params.ClientAssertionType)
	case clientSecret != "":
		body.Set("client_secret", clientSecret)
	}

	if required && (body.Get("client_secret") == "" && body.Get("client_assertion") == "") {
		return errors.New("client_secret or client_assertion is required but not provided")
	}

	return nil
}

// Helper for adding client authentication to an oauth.ClientAuthentication struct.
func (a *Authentication) addClientAuthenticationToClientAuthStruct(params *oauth.ClientAuthentication, required bool) error {
	if params.ClientID == "" {
		params.ClientID = a.clientID
	}

	if a.clientAssertionSigningKey != "" && a.clientAssertionSigningAlg != "" {
		alg, err := client.DetermineSigningAlgorithm(a.clientAssertionSigningAlg)
		if err != nil {
			return err
		}

		// Using the improved createClientAssertion with a standard lifetime
		clientAssertion, err := client.CreateClientAssertion(
			alg,
			a.clientAssertionSigningKey,
			params.ClientID,
			a.url.JoinPath("/").String(),
		)
		if err != nil {
			return err
		}

		params.ClientAssertion = clientAssertion
		params.ClientAssertionType = "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"
	} else if params.ClientSecret == "" && a.clientSecret != "" {
		params.ClientSecret = a.clientSecret
	}

	if required && (params.ClientSecret == "" && params.ClientAssertion == "") {
		return errors.New("client_secret or client_assertion is required but not provided")
	}

	return nil
}
