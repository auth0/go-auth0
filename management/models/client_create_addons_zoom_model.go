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

// ClientCreateAddonsZoom Zoom SSO configuration.
type ClientCreateAddonsZoom struct {
	// Zoom account name usually first segment of your Zoom URL, e.g. `https://acme-org.zoom.us` would be `acme-org`.
	Account              *string `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientCreateAddonsZoom ClientCreateAddonsZoom

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ClientCreateAddonsZoom) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsZoom) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ClientCreateAddonsZoom) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *ClientCreateAddonsZoom) SetAccount(v string) {
	o.Account = &v
}

func (o ClientCreateAddonsZoom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateAddonsZoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientCreateAddonsZoom) UnmarshalJSON(data []byte) (err error) {
	varClientCreateAddonsZoom := _ClientCreateAddonsZoom{}

	err = json.Unmarshal(data, &varClientCreateAddonsZoom)

	if err != nil {
		return err
	}

	*o = ClientCreateAddonsZoom(varClientCreateAddonsZoom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientCreateAddonsZoom struct {
	value *ClientCreateAddonsZoom
	isSet bool
}

func (v NullableClientCreateAddonsZoom) Get() *ClientCreateAddonsZoom {
	return v.value
}

func (v *NullableClientCreateAddonsZoom) Set(val *ClientCreateAddonsZoom) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateAddonsZoom) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateAddonsZoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateAddonsZoom(val *ClientCreateAddonsZoom) *NullableClientCreateAddonsZoom {
	return &NullableClientCreateAddonsZoom{value: val, isSet: true}
}

func (v NullableClientCreateAddonsZoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateAddonsZoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
