package management

import (
	"context"
	"encoding/json"
)

// UserAttributeProfile represents a User Attribute Profile object.
type UserAttributeProfile struct {
	// ID is the unique User Attribute Profile identifier.
	ID *string `json:"id,omitempty"`
	// Name is the name of the User Attribute Profile.
	Name *string `json:"name,omitempty"`
	// UserID is the User ID mapping configuration
	UserID *UserAttributeProfileUserID `json:"user_id,omitempty"`
	// UserAttributes User attributes configuration map. Keys are attribute names, values are the mapping configuration for each attribute.
	UserAttributes map[string]*UserAttributeProfileUserAttributes `json:"user_attributes,omitempty"`
}

// UserAttributeProfileUserID represents the configuration for the user ID mapping.
type UserAttributeProfileUserID struct {
	// OIDCMapping is the OIDC mapping for the user ID.
	OIDCMapping *string `json:"oidc_mapping,omitempty"`
	// SAMLMapping is the SAML mapping for the user ID.
	SAMLMapping *[]string `json:"saml_mapping,omitempty"`
	// SCIMMapping is the SCIM mapping for the user ID.
	SCIMMapping *string `json:"scim_mapping,omitempty"`
	// StrategyOverrides is a map of strategy names to their respective overrides for the User ID.
	// Supported strategies are: "ad", "adfs", "google-apps", "oidc", "okta", "pingfederate", "samlp", "waad"
	StrategyOverrides map[string]*UserAttributeProfileStrategyOverrides `json:"strategy_overrides,omitempty"`
}

// UserAttributeProfileOIDCMapping represents the OIDC mapping configuration.
type UserAttributeProfileOIDCMapping struct {
	// Mapping is the OIDC mapping field.
	Mapping *string `json:"mapping,omitempty"`
	// DisplayName is the display name for the OIDC mapping.
	DisplayName *string `json:"display_name,omitempty"`
}

// UserAttributeProfileStrategyOverrides represents the strategy-specific overrides for user attribute mappings.
type UserAttributeProfileStrategyOverrides struct {
	// OIDCMapping is the OIDC mapping override.
	OIDCMapping *string `json:"oidc_mapping,omitempty"`
	// SAMLMapping is the SAML mapping override.
	SAMLMapping *[]string `json:"saml_mapping,omitempty"`
	// SCIMMapping is the SCIM mapping override.
	SCIMMapping *string `json:"scim_mapping,omitempty"`
}

// UserAttributeProfileUserAttributes represents the configuration for a user attribute.
type UserAttributeProfileUserAttributes struct {
	// Description is a description of the user attribute.
	Description *string `json:"description,omitempty"`
	// Label is the display label for the user attribute.
	Label *string `json:"label,omitempty"`
	// ProfileRequired indicates whether the attribute is required in the profile.
	ProfileRequired *bool `json:"profile_required,omitempty"`
	// Auth0Mapping is the Auth0 mapping for the user attribute.
	Auth0Mapping *string `json:"auth0_mapping,omitempty"`
	// OIDCMapping is the OIDC mapping for the user attribute.
	OIDCMapping *UserAttributeProfileOIDCMapping `json:"oidc_mapping,omitempty"`
	// SAMLMapping is the SAML mapping for the user attribute.
	SAMLMapping *[]string `json:"saml_mapping,omitempty"`
	// SCIMMapping is the SCIM mapping for the user attribute.
	SCIMMapping *string `json:"scim_mapping,omitempty"`
	// StrategyOverrides is a map of strategy names to their respective overrides for the user attribute.
	StrategyOverrides map[string]*UserAttributesStrategyOverride `json:"strategy_overrides,omitempty"`
}

// UserAttributesStrategyOverride represents strategy-specific overrides for user attribute mappings.
type UserAttributesStrategyOverride struct {
	OIDCMapping *UserAttributeProfileOIDCMapping `json:"oidc_mapping,omitempty"`
	SAMLMapping *[]string                        `json:"saml_mapping,omitempty"`
	SCIMMapping *string                          `json:"scim_mapping,omitempty"`
}

// UserAttributeProfileList is a list of UserAttributeProfiles.
type UserAttributeProfileList struct {
	List
	UserAttributeProfiles []*UserAttributeProfile `json:"user_attribute_profiles"`
}

// UserAttributeProfileTemplateItem represents a User Attribute Profile Template item.
type UserAttributeProfileTemplateItem struct {
	ID          *string                       `json:"id,omitempty"`
	DisplayName *string                       `json:"display_name,omitempty"`
	Template    *UserAttributeProfileTemplate `json:"template,omitempty"`
}

// UserAttributeProfileTemplate represents a User Attribute Profile Template.
type UserAttributeProfileTemplate struct {
	Name           *string                                        `json:"name,omitempty"`
	UserID         *UserAttributeProfileUserID                    `json:"user_id,omitempty"`
	UserAttributes map[string]*UserAttributeProfileUserAttributes `json:"user_attributes,omitempty"`
}

// UserAttributeProfileTemplateList is a list of User Attribute Profile Templates.
type UserAttributeProfileTemplateList struct {
	UserAttributeProfileTemplates []*UserAttributeProfileTemplateItem `json:"user_attribute_profile_templates"`
}

// MarshalJSON ensures that we don't send the ID field when marshaling a UserAttributeProfile.
func (u *UserAttributeProfile) MarshalJSON() ([]byte, error) {
	type UserAttributeProfileSubset struct {
		Name           *string                                        `json:"name,omitempty"`
		UserID         *UserAttributeProfileUserID                    `json:"user_id,omitempty"`
		UserAttributes map[string]*UserAttributeProfileUserAttributes `json:"user_attributes,omitempty"`
	}

	return json.Marshal(&UserAttributeProfileSubset{
		Name:           u.Name,
		UserID:         u.UserID,
		UserAttributes: u.UserAttributes,
	})
}

// UserAttributeProfileManager manages Auth0 User Attribute Profile resources.
type UserAttributeProfileManager manager

// Create a new User Attribute Profile.
func (m *UserAttributeProfileManager) Create(ctx context.Context, u *UserAttributeProfile, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "POST", m.management.URI("user-attribute-profiles"), &u, opts...)
	return
}

// List all User Attribute Profiles.
func (m *UserAttributeProfileManager) List(ctx context.Context, opts ...RequestOption) (u *UserAttributeProfileList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("user-attribute-profiles"), &u, opts...)
	return
}

// Read retrieves a single User Attribute Profile by its ID.
func (m *UserAttributeProfileManager) Read(ctx context.Context, id string, opts ...RequestOption) (u *UserAttributeProfile, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("user-attribute-profiles", id), &u, opts...)
	return
}

// cleanForPatch ensure that the `UserID` field is set to an empty object if nil to allow resetting to defaults.
func (u *UserAttributeProfile) cleanForPatch() *UserAttributeProfile {
	if u.UserID == nil {
		u.UserID = &UserAttributeProfileUserID{}
	}

	return u
}

// Update an existing User Attribute Profile.
func (m *UserAttributeProfileManager) Update(ctx context.Context, id string, u *UserAttributeProfile, opts ...RequestOption) (err error) {
	if u != nil {
		u = u.cleanForPatch()
	}

	err = m.management.Request(ctx, "PATCH", m.management.URI("user-attribute-profiles", id), &u, opts...)

	return
}

// Delete a User Attribute Profile by its ID.
func (m *UserAttributeProfileManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	err = m.management.Request(ctx, "DELETE", m.management.URI("user-attribute-profiles", id), nil, opts...)
	return
}

// ListTemplates returns the templates.
func (m *UserAttributeProfileManager) ListTemplates(ctx context.Context, opts ...RequestOption) (u *UserAttributeProfileTemplateList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("user-attribute-profiles", "templates"), &u, opts...)
	return
}

// GetTemplate retrieves a single Template by its ID.
func (m *UserAttributeProfileManager) GetTemplate(ctx context.Context, id string, opts ...RequestOption) (u *UserAttributeProfileTemplateItem, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("user-attribute-profiles", "templates", id), &u, opts...)
	return
}
