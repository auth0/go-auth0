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

// ConnectionCreateOptionsAttributesEmailSignup struct for ConnectionCreateOptionsAttributesEmailSignup
type ConnectionCreateOptionsAttributesEmailSignup struct {
	Status       *ConnectionCreateOptionsAttributesEmailSignupStatus       `json:"status,omitempty"`
	Verification *ConnectionCreateOptionsAttributesEmailSignupVerification `json:"verification,omitempty"`
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsAttributesEmailSignup) GetStatus() ConnectionCreateOptionsAttributesEmailSignupStatus {
	if o == nil || IsNil(o.Status) {
		var ret ConnectionCreateOptionsAttributesEmailSignupStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsAttributesEmailSignup) GetStatusOk() (*ConnectionCreateOptionsAttributesEmailSignupStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsAttributesEmailSignup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ConnectionCreateOptionsAttributesEmailSignupStatus and assigns it to the Status field.
func (o *ConnectionCreateOptionsAttributesEmailSignup) SetStatus(v ConnectionCreateOptionsAttributesEmailSignupStatus) {
	o.Status = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsAttributesEmailSignup) GetVerification() ConnectionCreateOptionsAttributesEmailSignupVerification {
	if o == nil || IsNil(o.Verification) {
		var ret ConnectionCreateOptionsAttributesEmailSignupVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsAttributesEmailSignup) GetVerificationOk() (*ConnectionCreateOptionsAttributesEmailSignupVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsAttributesEmailSignup) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given ConnectionCreateOptionsAttributesEmailSignupVerification and assigns it to the Verification field.
func (o *ConnectionCreateOptionsAttributesEmailSignup) SetVerification(v ConnectionCreateOptionsAttributesEmailSignupVerification) {
	o.Verification = &v
}

func (o ConnectionCreateOptionsAttributesEmailSignup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCreateOptionsAttributesEmailSignup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	return toSerialize, nil
}

type NullableConnectionCreateOptionsAttributesEmailSignup struct {
	value *ConnectionCreateOptionsAttributesEmailSignup
	isSet bool
}

func (v NullableConnectionCreateOptionsAttributesEmailSignup) Get() *ConnectionCreateOptionsAttributesEmailSignup {
	return v.value
}

func (v *NullableConnectionCreateOptionsAttributesEmailSignup) Set(val *ConnectionCreateOptionsAttributesEmailSignup) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateOptionsAttributesEmailSignup) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateOptionsAttributesEmailSignup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateOptionsAttributesEmailSignup(val *ConnectionCreateOptionsAttributesEmailSignup) *NullableConnectionCreateOptionsAttributesEmailSignup {
	return &NullableConnectionCreateOptionsAttributesEmailSignup{value: val, isSet: true}
}

func (v NullableConnectionCreateOptionsAttributesEmailSignup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateOptionsAttributesEmailSignup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
