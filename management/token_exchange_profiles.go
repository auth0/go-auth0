package management

import (
	"context"
	"encoding/json"
	"time"
)

// TokenExchangeProfile represents a token exchange profile.
type TokenExchangeProfile struct {
	// ID is the unique identifier of the token exchange profile.
	ID *string `json:"id,omitempty"`
	// Name is the name of the token exchange profile.
	Name *string `json:"name,omitempty"`
	// SubjectTokenType is the type of the subject token.
	SubjectTokenType *string `json:"subject_token_type,omitempty"`
	// ActionID is the identifier of the action.
	ActionID *string `json:"action_id,omitempty"`
	// Type is the type of the token exchange profile.
	Type *string `json:"type,omitempty"`
	// CreatedAt is the date and time when the token exchange profile was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt is the date and time when the token exchange profile was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// TokenExchangeProfileList is a list of TokenExchangeProfiles.
type TokenExchangeProfileList struct {
	List
	TokenExchangeProfiles []*TokenExchangeProfile `json:"token_exchange_profiles"`
}

// MarshalJSON implements the json.Marshaler interface for the TokenExchangeProfile type.
func (t *TokenExchangeProfile) MarshalJSON() ([]byte, error) {
	type TokenExchangeProfileSubset struct {
		Name             *string `json:"name,omitempty"`
		SubjectTokenType *string `json:"subject_token_type,omitempty"`
		ActionID         *string `json:"action_id,omitempty"`
		Type             *string `json:"type,omitempty"`
	}

	return json.Marshal(&TokenExchangeProfileSubset{
		Name:             t.Name,
		SubjectTokenType: t.SubjectTokenType,
		ActionID:         t.ActionID,
		Type:             t.Type,
	})
}

func (t *TokenExchangeProfile) cleanForPatch() {
	*t = TokenExchangeProfile{
		Name:             t.Name,
		SubjectTokenType: t.SubjectTokenType,
	}
}

// TokenExchangeProfileManager manages Auth0 Token Exchange Profile resources.
type TokenExchangeProfileManager manager

// Create a new token exchange profile.
//
// See: https://auth0.com/docs/api/management/v2#!/token-exchange-profiles/post_token_exchange_profile
func (m *TokenExchangeProfileManager) Create(ctx context.Context, t *TokenExchangeProfile, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("token-exchange-profiles"), t, opts...)
	return
}

// List all token exchange profiles.
//
// See: https://auth0.com/docs/api/management/v2#!/token-exchange-profiles/get_token_exchange_profiles
func (m *TokenExchangeProfileManager) List(ctx context.Context, opts ...RequestOption) (t *TokenExchangeProfileList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("token-exchange-profiles"), &t, applyListCheckpointDefaults(opts))
	return
}

// Read a single token exchange profile against the ID.
//
// See: https://auth0.com/docs/api/management/v2#!/token-exchange-profiles/get_token_exchange_profile
func (m *TokenExchangeProfileManager) Read(ctx context.Context, id string, opts ...RequestOption) (t *TokenExchangeProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("token-exchange-profiles", id), &t, opts...)
	return
}

// Update an existing token exchange profile against the ID.
//
// See: https://auth0.com/docs/api/management/v2#!/token-exchange-profiles/patch_token_exchange_profile
func (m *TokenExchangeProfileManager) Update(ctx context.Context, id string, t *TokenExchangeProfile, opts ...RequestOption) (err error) {
	if t != nil {
		t.cleanForPatch()
	}

	err = m.management.Request(ctx, "PATCH", m.management.URI("token-exchange-profiles", id), t, opts...)

	return
}

// Delete an existing token exchange profile against the ID.
//
// See: https://auth0.com/docs/api/management/v2#!/token-exchange-profiles/delete_token_exchange_profile
func (m *TokenExchangeProfileManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("token-exchange-profiles", id), nil, opts...)
	return
}
