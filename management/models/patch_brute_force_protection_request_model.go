/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// PatchBruteForceProtectionRequest struct for PatchBruteForceProtectionRequest
type PatchBruteForceProtectionRequest struct {
	// Whether or not brute force attack protections are active.
	Enabled *bool `json:"enabled,omitempty"`
	// Action to take when a brute force protection threshold is violated.         Possible values: <code>block</code>, <code>user_notification</code>.
	Shields []GetBruteForceProtection200ResponseShieldsInner `json:"shields,omitempty"`
	// List of trusted IP addresses that will not have attack protection enforced against them.
	Allowlist []GetBruteForceProtection200ResponseAllowlistInner `json:"allowlist,omitempty"`
	Mode      *GetBruteForceProtection200ResponseMode            `json:"mode,omitempty"`
	// Maximum number of unsuccessful attempts.
	MaxAttempts *int32 `json:"max_attempts,omitempty"`
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchBruteForceProtectionRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBruteForceProtectionRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchBruteForceProtectionRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchBruteForceProtectionRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShields returns the Shields field value if set, zero value otherwise.
func (o *PatchBruteForceProtectionRequest) GetShields() []GetBruteForceProtection200ResponseShieldsInner {
	if o == nil || IsNil(o.Shields) {
		var ret []GetBruteForceProtection200ResponseShieldsInner
		return ret
	}
	return o.Shields
}

// GetShieldsOk returns a tuple with the Shields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBruteForceProtectionRequest) GetShieldsOk() ([]GetBruteForceProtection200ResponseShieldsInner, bool) {
	if o == nil || IsNil(o.Shields) {
		return nil, false
	}
	return o.Shields, true
}

// HasShields returns a boolean if a field has been set.
func (o *PatchBruteForceProtectionRequest) HasShields() bool {
	if o != nil && !IsNil(o.Shields) {
		return true
	}

	return false
}

// SetShields gets a reference to the given []GetBruteForceProtection200ResponseShieldsInner and assigns it to the Shields field.
func (o *PatchBruteForceProtectionRequest) SetShields(v []GetBruteForceProtection200ResponseShieldsInner) {
	o.Shields = v
}

// GetAllowlist returns the Allowlist field value if set, zero value otherwise.
func (o *PatchBruteForceProtectionRequest) GetAllowlist() []GetBruteForceProtection200ResponseAllowlistInner {
	if o == nil || IsNil(o.Allowlist) {
		var ret []GetBruteForceProtection200ResponseAllowlistInner
		return ret
	}
	return o.Allowlist
}

// GetAllowlistOk returns a tuple with the Allowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBruteForceProtectionRequest) GetAllowlistOk() ([]GetBruteForceProtection200ResponseAllowlistInner, bool) {
	if o == nil || IsNil(o.Allowlist) {
		return nil, false
	}
	return o.Allowlist, true
}

// HasAllowlist returns a boolean if a field has been set.
func (o *PatchBruteForceProtectionRequest) HasAllowlist() bool {
	if o != nil && !IsNil(o.Allowlist) {
		return true
	}

	return false
}

// SetAllowlist gets a reference to the given []GetBruteForceProtection200ResponseAllowlistInner and assigns it to the Allowlist field.
func (o *PatchBruteForceProtectionRequest) SetAllowlist(v []GetBruteForceProtection200ResponseAllowlistInner) {
	o.Allowlist = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchBruteForceProtectionRequest) GetMode() GetBruteForceProtection200ResponseMode {
	if o == nil || IsNil(o.Mode) {
		var ret GetBruteForceProtection200ResponseMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBruteForceProtectionRequest) GetModeOk() (*GetBruteForceProtection200ResponseMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchBruteForceProtectionRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given GetBruteForceProtection200ResponseMode and assigns it to the Mode field.
func (o *PatchBruteForceProtectionRequest) SetMode(v GetBruteForceProtection200ResponseMode) {
	o.Mode = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *PatchBruteForceProtectionRequest) GetMaxAttempts() int32 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret int32
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBruteForceProtectionRequest) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *PatchBruteForceProtectionRequest) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given int32 and assigns it to the MaxAttempts field.
func (o *PatchBruteForceProtectionRequest) SetMaxAttempts(v int32) {
	o.MaxAttempts = &v
}

func (o PatchBruteForceProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchBruteForceProtectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Shields) {
		toSerialize["shields"] = o.Shields
	}
	if !IsNil(o.Allowlist) {
		toSerialize["allowlist"] = o.Allowlist
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.MaxAttempts) {
		toSerialize["max_attempts"] = o.MaxAttempts
	}
	return toSerialize, nil
}

type NullablePatchBruteForceProtectionRequest struct {
	value *PatchBruteForceProtectionRequest
	isSet bool
}

func (v NullablePatchBruteForceProtectionRequest) Get() *PatchBruteForceProtectionRequest {
	return v.value
}

func (v *NullablePatchBruteForceProtectionRequest) Set(val *PatchBruteForceProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBruteForceProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBruteForceProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBruteForceProtectionRequest(val *PatchBruteForceProtectionRequest) *NullablePatchBruteForceProtectionRequest {
	return &NullablePatchBruteForceProtectionRequest{value: val, isSet: true}
}

func (v NullablePatchBruteForceProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBruteForceProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
