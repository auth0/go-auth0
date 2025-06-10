package management

import (
	"context"
	"encoding/json"
)

// NetworkACL : Network ACL.
type NetworkACL struct {
	// The id of the Network ACL
	ID *string `json:"id,omitempty"`
	// The description of the Network ACL
	Description *string `json:"description,omitempty"`
	// Whether the Network ACL is active
	Active *bool `json:"active,omitempty"`
	// The priority of the Network ACL
	Priority *int `json:"priority,omitempty"`
	// The rule of the Network ACL
	Rule *NetworkACLRule `json:"rule,omitempty"`
	// The timestamp when the Network ACL Configuration was created
	CreatedAt *string `json:"created_at,omitempty"`
	// The timestamp when the Network ACL Configuration was last updated
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NetworkACLRule : Network ACL Rule.
type NetworkACLRule struct {
	// The action of the Network ACL Rule
	Action *NetworkACLRuleAction `json:"action,omitempty"`
	// The match of the Network ACL Rule
	Match *NetworkACLRuleMatch `json:"match,omitempty"`
	// The not match of the Network ACL Rule
	NotMatch *NetworkACLRuleMatch `json:"not_match,omitempty"`
	// The scope of the Network ACL Rule
	Scope *string `json:"scope,omitempty"`
}

// NetworkACLRuleAction : Network ACL Rule Action.
type NetworkACLRuleAction struct {
	// The block of the Network ACL Rule Action
	Block *bool `json:"block,omitempty"`
	// The allow of the Network ACL Rule Action
	Allow *bool `json:"allow,omitempty"`
	// The log of the Network ACL Rule Action
	Log *bool `json:"log,omitempty"`
	// The redirect of the Network ACL Rule Action
	Redirect *bool `json:"redirect,omitempty"`
	// The redirect_uri of the Network ACL Rule Action
	RedirectURI *string `json:"redirect_uri,omitempty"`
}

// NetworkACLRuleMatch : Network ACL Rule Match.
type NetworkACLRuleMatch struct {
	// Anonymous Proxy as reported by GeoIP
	AnonymousProxy *bool `json:"anonymous_proxy,omitempty"`
	// ASNs
	Asns []int `json:"asns,omitempty"`
	// Geo Country Codes
	GeoCountryCodes *[]string `json:"geo_country_codes,omitempty"`
	// Geo Subdivision Codes
	GeoSubdivisionCodes *[]string `json:"geo_subdivision_codes,omitempty"`
	// IPv4 CIDRs
	IPv4Cidrs *[]string `json:"ipv4_cidrs,omitempty"`
	// IPv6 CIDRs
	IPv6Cidrs *[]string `json:"ipv6_cidrs,omitempty"`
	// JA3 Fingerprints
	Ja3Fingerprints *[]string `json:"ja3_fingerprints,omitempty"`
	// JA4 Fingerprints
	Ja4Fingerprints *[]string `json:"ja4_fingerprints,omitempty"`
	// User Agents
	UserAgents *[]string `json:"user_agents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
//
// It is required to handle the json field page_background, which can either
// be a hex color string, or an object describing a gradient.
func (n *NetworkACL) MarshalJSON() ([]byte, error) {
	type networkACLSubset struct {
		Description *string         `json:"description,omitempty"`
		Active      *bool           `json:"active,omitempty"`
		Priority    *int            `json:"priority,omitempty"`
		Rule        *NetworkACLRule `json:"rule,omitempty"`
	}

	return json.Marshal(&networkACLSubset{
		Description: n.Description,
		Active:      n.Active,
		Priority:    n.Priority,
		Rule:        n.Rule,
	})
}

// NetworkACLManager manages Network ACL resources.
type NetworkACLManager manager

// Create a new Network ACL.
func (m *NetworkACLManager) Create(ctx context.Context, n *NetworkACL, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("network-acls"), n, opts...)
}

// Read a Network ACL.
func (m *NetworkACLManager) Read(ctx context.Context, id string, opts ...RequestOption) (n *NetworkACL, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("network-acls", id), &n, opts...)
	return
}

// Update a Network ACL.
func (m *NetworkACLManager) Update(ctx context.Context, id string, n *NetworkACL, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PUT", m.management.URI("network-acls", id), n, opts...)
}

// Patch a Network ACL.
func (m *NetworkACLManager) Patch(ctx context.Context, id string, n *NetworkACL, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("network-acls", id), n, opts...)
}

// Delete a Network ACL.
func (m *NetworkACLManager) Delete(ctx context.Context, id string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("network-acls", id), nil, opts...)
}

// List all Network ACLs.
func (m *NetworkACLManager) List(ctx context.Context, opts ...RequestOption) (nl []*NetworkACL, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("network-acls"), &nl, opts...)
	return
}
