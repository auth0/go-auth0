package management

import "context"

// Hook is a secure, self-contained function that
// allows the behavior of Auth0 to be customized.
//
// See: https://auth0.com/docs/customize/hooks
type Hook struct {
	// The hook's identifier.
	ID *string `json:"id,omitempty"`

	// The name of the hook. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name *string `json:"name,omitempty"`

	// A script that contains the hook's code.
	Script *string `json:"script,omitempty"`

	// The extensibility point name
	// Can currently be any of the following:
	// "credentials-exchange", "pre-user-registration",
	// "post-user-registration", "post-change-password"
	TriggerID *string `json:"triggerId,omitempty"`

	// Used to store additional metadata
	Dependencies *map[string]string `json:"dependencies,omitempty"`

	// Enabled should be set to true if the hook is enabled, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
}

// HookList is a list of Hooks.
type HookList struct {
	List
	Hooks []*Hook `json:"hooks"`
}

// HookSecrets are the secret keys and values associated with a Hook.
type HookSecrets map[string]string

// Keys gets the configured hook secret keys.
func (s HookSecrets) Keys() []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

// Difference returns a new map containing only keys which
// are present in s that are missing from other.
func (s HookSecrets) difference(other HookSecrets) HookSecrets {
	d := make(HookSecrets)
	for k, v := range s {
		if _, ok := other[k]; !ok {
			d[k] = v
		}
	}
	return d
}

// Intersection returns a new map containing only
// keys which are present in both s and other.
func (s HookSecrets) intersection(other HookSecrets) HookSecrets {
	i := make(HookSecrets)
	for k, v := range s {
		if _, ok := other[k]; ok {
			i[k] = v
		}
	}
	return i
}

// HookManager manages Auth0 Hook resources.
type HookManager manager

// Create a new hook.
//
// Note: Changing a hook's trigger changes the signature of the script and should be done with caution.
//
// See: https://auth0.com/docs/api/management/v2#!/Hooks/post_hooks
func (m *HookManager) Create(ctx context.Context, h *Hook, opts ...RequestOption) error {
	return m.management.Request(ctx, "POST", m.management.URI("hooks"), h, opts...)
}

// Read hook details. Accepts a list of fields to include or exclude in the result.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks_by_id
func (m *HookManager) Read(ctx context.Context, id string, opts ...RequestOption) (h *Hook, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("hooks", id), &h, opts...)
	return
}

// Update an existing hook.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/patch_hooks_by_id
func (m *HookManager) Update(ctx context.Context, id string, h *Hook, opts ...RequestOption) error {
	return m.management.Request(ctx, "PATCH", m.management.URI("hooks", id), h, opts...)
}

// Delete a hook.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/delete_hooks_by_id
func (m *HookManager) Delete(ctx context.Context, id string, opts ...RequestOption) error {
	return m.management.Request(ctx, "DELETE", m.management.URI("hooks", id), nil, opts...)
}

// List hooks.
//
// For information on how to paginate using this function see https://pkg.go.dev/github.com/auth0/go-auth0/management#hdr-Page_Based_Pagination
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks
func (m *HookManager) List(ctx context.Context, opts ...RequestOption) (l *HookList, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("hooks"), &l, applyListDefaults(opts))
	return
}

// CreateSecrets adds one or more secrets to an existing hook. A hook can have a
// maximum of 20 secrets.
//
// See: https://auth0.com/docs/api/management/v2#!/Hooks/post_secrets
func (m *HookManager) CreateSecrets(ctx context.Context, hookID string, s HookSecrets, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "POST", m.management.URI("hooks", hookID, "secrets"), &s, opts...)
}

// UpdateSecrets updates one or more existing secrets for an existing hook.
//
// See: https://auth0.com/docs/api/management/v2#!/Hooks/patch_secrets
func (m *HookManager) UpdateSecrets(ctx context.Context, hookID string, s HookSecrets, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "PATCH", m.management.URI("hooks", hookID, "secrets"), &s, opts...)
}

// ReplaceSecrets replaces existing secrets with the provided ones.
//
// Note: ReplaceSecrets is a wrapper method and will internally call Secrets,
// CreateSecrets, UpdateSecrets or RemoveSecrets as needed in order to replicate
// PUT semantics.
func (m *HookManager) ReplaceSecrets(ctx context.Context, hookID string, s HookSecrets, opts ...RequestOption) (err error) {
	o, err := m.Secrets(ctx, hookID, opts...)
	if err != nil {
		return err
	}
	if add := s.difference(o); len(add) > 0 {
		err = m.CreateSecrets(ctx, hookID, add, opts...)
	}
	if update := s.intersection(o); len(update) > 0 {
		err = m.UpdateSecrets(ctx, hookID, update, opts...)
	}
	if rm := o.difference(s); len(rm) > 0 {
		err = m.RemoveSecrets(ctx, hookID, rm.Keys(), opts...)
	}
	return err
}

// Secrets retrieves a hook's secrets by the ID of the hook.
//
// Note: For security, hook secret values cannot be retrieved outside rule
// execution (they all appear as "_VALUE_NOT_SHOWN_").
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/get_secrets
func (m *HookManager) Secrets(ctx context.Context, hookID string, opts ...RequestOption) (s HookSecrets, err error) {
	err = m.management.Request(ctx, "GET", m.management.URI("hooks", hookID, "secrets"), &s, opts...)
	return
}

// RemoveSecrets deletes one or more existing secrets for a given hook. Accepts
// an array of secret names to delete.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/delete_secrets
func (m *HookManager) RemoveSecrets(ctx context.Context, hookID string, keys []string, opts ...RequestOption) (err error) {
	return m.management.Request(ctx, "DELETE", m.management.URI("hooks", hookID, "secrets"), keys, opts...)
}

// RemoveAllSecrets removes all secrets associated with a given hook.
func (m *HookManager) RemoveAllSecrets(ctx context.Context, hookID string, opts ...RequestOption) (err error) {
	s, err := m.Secrets(ctx, hookID, opts...)
	if err != nil {
		return err
	}
	keys := s.Keys()
	if len(keys) > 0 {
		err = m.RemoveSecrets(ctx, hookID, keys, opts...)
	}
	return err
}
