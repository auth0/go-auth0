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

// GetBruteForceProtection200Response struct for GetBruteForceProtection200Response
type GetBruteForceProtection200Response struct {
	// Whether or not brute force attack protections are active.
	Enabled *bool `json:"enabled,omitempty"`
	// Action to take when a brute force protection threshold is violated.         Possible values: <code>block</code>, <code>user_notification</code>.
	Shields []GetBruteForceProtection200ResponseShieldsInner `json:"shields,omitempty"`
	// List of trusted IP addresses that will not have attack protection enforced against them.
	Allowlist []GetBruteForceProtection200ResponseAllowlistInner `json:"allowlist,omitempty"`
	Mode      *GetBruteForceProtection200ResponseMode            `json:"mode,omitempty"`
	// Maximum number of unsuccessful attempts.
	MaxAttempts          *int32 `json:"max_attempts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBruteForceProtection200Response GetBruteForceProtection200Response

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetBruteForceProtection200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBruteForceProtection200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetBruteForceProtection200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetBruteForceProtection200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShields returns the Shields field value if set, zero value otherwise.
func (o *GetBruteForceProtection200Response) GetShields() []GetBruteForceProtection200ResponseShieldsInner {
	if o == nil || IsNil(o.Shields) {
		var ret []GetBruteForceProtection200ResponseShieldsInner
		return ret
	}
	return o.Shields
}

// GetShieldsOk returns a tuple with the Shields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBruteForceProtection200Response) GetShieldsOk() ([]GetBruteForceProtection200ResponseShieldsInner, bool) {
	if o == nil || IsNil(o.Shields) {
		return nil, false
	}
	return o.Shields, true
}

// HasShields returns a boolean if a field has been set.
func (o *GetBruteForceProtection200Response) HasShields() bool {
	if o != nil && !IsNil(o.Shields) {
		return true
	}

	return false
}

// SetShields gets a reference to the given []GetBruteForceProtection200ResponseShieldsInner and assigns it to the Shields field.
func (o *GetBruteForceProtection200Response) SetShields(v []GetBruteForceProtection200ResponseShieldsInner) {
	o.Shields = v
}

// GetAllowlist returns the Allowlist field value if set, zero value otherwise.
func (o *GetBruteForceProtection200Response) GetAllowlist() []GetBruteForceProtection200ResponseAllowlistInner {
	if o == nil || IsNil(o.Allowlist) {
		var ret []GetBruteForceProtection200ResponseAllowlistInner
		return ret
	}
	return o.Allowlist
}

// GetAllowlistOk returns a tuple with the Allowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBruteForceProtection200Response) GetAllowlistOk() ([]GetBruteForceProtection200ResponseAllowlistInner, bool) {
	if o == nil || IsNil(o.Allowlist) {
		return nil, false
	}
	return o.Allowlist, true
}

// HasAllowlist returns a boolean if a field has been set.
func (o *GetBruteForceProtection200Response) HasAllowlist() bool {
	if o != nil && !IsNil(o.Allowlist) {
		return true
	}

	return false
}

// SetAllowlist gets a reference to the given []GetBruteForceProtection200ResponseAllowlistInner and assigns it to the Allowlist field.
func (o *GetBruteForceProtection200Response) SetAllowlist(v []GetBruteForceProtection200ResponseAllowlistInner) {
	o.Allowlist = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetBruteForceProtection200Response) GetMode() GetBruteForceProtection200ResponseMode {
	if o == nil || IsNil(o.Mode) {
		var ret GetBruteForceProtection200ResponseMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBruteForceProtection200Response) GetModeOk() (*GetBruteForceProtection200ResponseMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetBruteForceProtection200Response) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given GetBruteForceProtection200ResponseMode and assigns it to the Mode field.
func (o *GetBruteForceProtection200Response) SetMode(v GetBruteForceProtection200ResponseMode) {
	o.Mode = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *GetBruteForceProtection200Response) GetMaxAttempts() int32 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret int32
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBruteForceProtection200Response) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *GetBruteForceProtection200Response) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given int32 and assigns it to the MaxAttempts field.
func (o *GetBruteForceProtection200Response) SetMaxAttempts(v int32) {
	o.MaxAttempts = &v
}

func (o GetBruteForceProtection200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBruteForceProtection200Response) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBruteForceProtection200Response) UnmarshalJSON(data []byte) (err error) {
	varGetBruteForceProtection200Response := _GetBruteForceProtection200Response{}

	err = json.Unmarshal(data, &varGetBruteForceProtection200Response)

	if err != nil {
		return err
	}

	*o = GetBruteForceProtection200Response(varGetBruteForceProtection200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "shields")
		delete(additionalProperties, "allowlist")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "max_attempts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBruteForceProtection200Response struct {
	value *GetBruteForceProtection200Response
	isSet bool
}

func (v NullableGetBruteForceProtection200Response) Get() *GetBruteForceProtection200Response {
	return v.value
}

func (v *NullableGetBruteForceProtection200Response) Set(val *GetBruteForceProtection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBruteForceProtection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBruteForceProtection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBruteForceProtection200Response(val *GetBruteForceProtection200Response) *NullableGetBruteForceProtection200Response {
	return &NullableGetBruteForceProtection200Response{value: val, isSet: true}
}

func (v NullableGetBruteForceProtection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBruteForceProtection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
