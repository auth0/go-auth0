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

// PatchSuspiciousIpThrottlingRequestStage Holds per-stage configuration options (max_attempts and rate).
type PatchSuspiciousIpThrottlingRequestStage struct {
	PreLogin            *PatchSuspiciousIpThrottlingRequestStagePreLogin            `json:"pre-login,omitempty"`
	PreUserRegistration *PatchSuspiciousIpThrottlingRequestStagePreUserRegistration `json:"pre-user-registration,omitempty"`
}

// GetPreLogin returns the PreLogin field value if set, zero value otherwise.
func (o *PatchSuspiciousIpThrottlingRequestStage) GetPreLogin() PatchSuspiciousIpThrottlingRequestStagePreLogin {
	if o == nil || IsNil(o.PreLogin) {
		var ret PatchSuspiciousIpThrottlingRequestStagePreLogin
		return ret
	}
	return *o.PreLogin
}

// GetPreLoginOk returns a tuple with the PreLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSuspiciousIpThrottlingRequestStage) GetPreLoginOk() (*PatchSuspiciousIpThrottlingRequestStagePreLogin, bool) {
	if o == nil || IsNil(o.PreLogin) {
		return nil, false
	}
	return o.PreLogin, true
}

// HasPreLogin returns a boolean if a field has been set.
func (o *PatchSuspiciousIpThrottlingRequestStage) HasPreLogin() bool {
	if o != nil && !IsNil(o.PreLogin) {
		return true
	}

	return false
}

// SetPreLogin gets a reference to the given PatchSuspiciousIpThrottlingRequestStagePreLogin and assigns it to the PreLogin field.
func (o *PatchSuspiciousIpThrottlingRequestStage) SetPreLogin(v PatchSuspiciousIpThrottlingRequestStagePreLogin) {
	o.PreLogin = &v
}

// GetPreUserRegistration returns the PreUserRegistration field value if set, zero value otherwise.
func (o *PatchSuspiciousIpThrottlingRequestStage) GetPreUserRegistration() PatchSuspiciousIpThrottlingRequestStagePreUserRegistration {
	if o == nil || IsNil(o.PreUserRegistration) {
		var ret PatchSuspiciousIpThrottlingRequestStagePreUserRegistration
		return ret
	}
	return *o.PreUserRegistration
}

// GetPreUserRegistrationOk returns a tuple with the PreUserRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSuspiciousIpThrottlingRequestStage) GetPreUserRegistrationOk() (*PatchSuspiciousIpThrottlingRequestStagePreUserRegistration, bool) {
	if o == nil || IsNil(o.PreUserRegistration) {
		return nil, false
	}
	return o.PreUserRegistration, true
}

// HasPreUserRegistration returns a boolean if a field has been set.
func (o *PatchSuspiciousIpThrottlingRequestStage) HasPreUserRegistration() bool {
	if o != nil && !IsNil(o.PreUserRegistration) {
		return true
	}

	return false
}

// SetPreUserRegistration gets a reference to the given PatchSuspiciousIpThrottlingRequestStagePreUserRegistration and assigns it to the PreUserRegistration field.
func (o *PatchSuspiciousIpThrottlingRequestStage) SetPreUserRegistration(v PatchSuspiciousIpThrottlingRequestStagePreUserRegistration) {
	o.PreUserRegistration = &v
}

func (o PatchSuspiciousIpThrottlingRequestStage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSuspiciousIpThrottlingRequestStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreLogin) {
		toSerialize["pre-login"] = o.PreLogin
	}
	if !IsNil(o.PreUserRegistration) {
		toSerialize["pre-user-registration"] = o.PreUserRegistration
	}
	return toSerialize, nil
}

type NullablePatchSuspiciousIpThrottlingRequestStage struct {
	value *PatchSuspiciousIpThrottlingRequestStage
	isSet bool
}

func (v NullablePatchSuspiciousIpThrottlingRequestStage) Get() *PatchSuspiciousIpThrottlingRequestStage {
	return v.value
}

func (v *NullablePatchSuspiciousIpThrottlingRequestStage) Set(val *PatchSuspiciousIpThrottlingRequestStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSuspiciousIpThrottlingRequestStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSuspiciousIpThrottlingRequestStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSuspiciousIpThrottlingRequestStage(val *PatchSuspiciousIpThrottlingRequestStage) *NullablePatchSuspiciousIpThrottlingRequestStage {
	return &NullablePatchSuspiciousIpThrottlingRequestStage{value: val, isSet: true}
}

func (v NullablePatchSuspiciousIpThrottlingRequestStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSuspiciousIpThrottlingRequestStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
