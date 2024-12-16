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

// ConnectionCreateOptionsValidation Options for validation
type ConnectionCreateOptionsValidation struct {
	Username NullableConnectionCreateOptionsValidationUsername `json:"username,omitempty"`
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionCreateOptionsValidation) GetUsername() ConnectionCreateOptionsValidationUsername {
	if o == nil || IsNil(o.Username.Get()) {
		var ret ConnectionCreateOptionsValidationUsername
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionCreateOptionsValidation) GetUsernameOk() (*ConnectionCreateOptionsValidationUsername, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsValidation) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableConnectionCreateOptionsValidationUsername and assigns it to the Username field.
func (o *ConnectionCreateOptionsValidation) SetUsername(v ConnectionCreateOptionsValidationUsername) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ConnectionCreateOptionsValidation) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ConnectionCreateOptionsValidation) UnsetUsername() {
	o.Username.Unset()
}

func (o ConnectionCreateOptionsValidation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCreateOptionsValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return toSerialize, nil
}

type NullableConnectionCreateOptionsValidation struct {
	value *ConnectionCreateOptionsValidation
	isSet bool
}

func (v NullableConnectionCreateOptionsValidation) Get() *ConnectionCreateOptionsValidation {
	return v.value
}

func (v *NullableConnectionCreateOptionsValidation) Set(val *ConnectionCreateOptionsValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateOptionsValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateOptionsValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateOptionsValidation(val *ConnectionCreateOptionsValidation) *NullableConnectionCreateOptionsValidation {
	return &NullableConnectionCreateOptionsValidation{value: val, isSet: true}
}

func (v NullableConnectionCreateOptionsValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateOptionsValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
