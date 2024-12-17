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

// TenantSettingsSessions Sessions related settings for tenant
type TenantSettingsSessions struct {
	// Whether to bypass prompting logic (false) when performing OIDC Logout
	OidcLogoutPromptEnabled *bool `json:"oidc_logout_prompt_enabled,omitempty"`
}

// GetOidcLogoutPromptEnabled returns the OidcLogoutPromptEnabled field value if set, zero value otherwise.
func (o *TenantSettingsSessions) GetOidcLogoutPromptEnabled() bool {
	if o == nil || IsNil(o.OidcLogoutPromptEnabled) {
		var ret bool
		return ret
	}
	return *o.OidcLogoutPromptEnabled
}

// GetOidcLogoutPromptEnabledOk returns a tuple with the OidcLogoutPromptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettingsSessions) GetOidcLogoutPromptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OidcLogoutPromptEnabled) {
		return nil, false
	}
	return o.OidcLogoutPromptEnabled, true
}

// HasOidcLogoutPromptEnabled returns a boolean if a field has been set.
func (o *TenantSettingsSessions) HasOidcLogoutPromptEnabled() bool {
	if o != nil && !IsNil(o.OidcLogoutPromptEnabled) {
		return true
	}

	return false
}

// SetOidcLogoutPromptEnabled gets a reference to the given bool and assigns it to the OidcLogoutPromptEnabled field.
func (o *TenantSettingsSessions) SetOidcLogoutPromptEnabled(v bool) {
	o.OidcLogoutPromptEnabled = &v
}

func (o TenantSettingsSessions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantSettingsSessions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OidcLogoutPromptEnabled) {
		toSerialize["oidc_logout_prompt_enabled"] = o.OidcLogoutPromptEnabled
	}
	return toSerialize, nil
}

type NullableTenantSettingsSessions struct {
	value *TenantSettingsSessions
	isSet bool
}

func (v NullableTenantSettingsSessions) Get() *TenantSettingsSessions {
	return v.value
}

func (v *NullableTenantSettingsSessions) Set(val *TenantSettingsSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantSettingsSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantSettingsSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantSettingsSessions(val *TenantSettingsSessions) *NullableTenantSettingsSessions {
	return &NullableTenantSettingsSessions{value: val, isSet: true}
}

func (v NullableTenantSettingsSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantSettingsSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
