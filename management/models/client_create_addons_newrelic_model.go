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

// ClientCreateAddonsNewrelic New Relic SSO configuration.
type ClientCreateAddonsNewrelic struct {
	// Your New Relic Account ID found in your New Relic URL after the `/accounts/` path. e.g. `https://rpm.newrelic.com/accounts/123456/query` would be `123456`.
	Account              *string `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientCreateAddonsNewrelic ClientCreateAddonsNewrelic

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ClientCreateAddonsNewrelic) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsNewrelic) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ClientCreateAddonsNewrelic) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ClientCreateAddonsNewrelic) SetAccount(v string) {
	o.Account = &v
}

func (o ClientCreateAddonsNewrelic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateAddonsNewrelic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientCreateAddonsNewrelic) UnmarshalJSON(data []byte) (err error) {
	varClientCreateAddonsNewrelic := _ClientCreateAddonsNewrelic{}

	err = json.Unmarshal(data, &varClientCreateAddonsNewrelic)

	if err != nil {
		return err
	}

	*o = ClientCreateAddonsNewrelic(varClientCreateAddonsNewrelic)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientCreateAddonsNewrelic struct {
	value *ClientCreateAddonsNewrelic
	isSet bool
}

func (v NullableClientCreateAddonsNewrelic) Get() *ClientCreateAddonsNewrelic {
	return v.value
}

func (v *NullableClientCreateAddonsNewrelic) Set(val *ClientCreateAddonsNewrelic) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateAddonsNewrelic) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateAddonsNewrelic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateAddonsNewrelic(val *ClientCreateAddonsNewrelic) *NullableClientCreateAddonsNewrelic {
	return &NullableClientCreateAddonsNewrelic{value: val, isSet: true}
}

func (v NullableClientCreateAddonsNewrelic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateAddonsNewrelic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
