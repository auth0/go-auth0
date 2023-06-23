package authentication

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

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
	OAuth        *OAuth
	Passwordless *Passwordless

	auth0ClientInfo       *client.Auth0ClientInfo
	basePath              string
	common                manager
	clientID              string
	clientSecret          string
	idTokenClockTolerance time.Duration
	debug                 bool
	http                  *http.Client
	idTokenSigningAlg     string
	idTokenValidator      *idtokenvalidator.IDTokenValidator
	url                   *url.URL
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
	}

	for _, option := range options {
		option(a)
	}

	a.common.authentication = a
	a.Database = (*Database)(&a.common)
	a.OAuth = (*OAuth)(&a.common)
	a.Passwordless = (*Passwordless)(&a.common)

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
// may return less information than a a scope of `openid profile email`.
//
// See: https://auth0.com/docs/api/authentication?http#get-user-info
func (a *Authentication) UserInfo(ctx context.Context, accessToken string, opts ...RequestOption) (user *UserInfoResponse, err error) {
	opts = append(opts, Header("Authorization", "Bearer "+accessToken))
	err = a.Request(ctx, "GET", a.URI("userinfo"), nil, &user, opts...)
	return
}
